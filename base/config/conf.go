package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	//DEFAULT_SECTION 默认段名称
	DEFAULT_SECTION = "DEFAULT"
	//_DEPTH_VALUES Maximum allowed depth when recursively substituing variable names.
	_DEPTH_VALUES = 200
)

// Parse error types.
const (
	ErrSectionNotFound = iota + 1
	ErrKeyNotFound
	ErrBlankSectionName
	ErrCouldNotParse
)

var LineBreak = "\n"

// Variable regexp pattern: %(variable)s
var varPattern = regexp.MustCompile(`%\(([^\)]+)\)s`)

func init() {
	if runtime.GOOS == "windows" {
		LineBreak = "\r\n"
	}
}

//ConfigFile 对应ini配置文件的内容（除注释）
type ConfigFile struct {
	lock      sync.RWMutex                 // Go map is not safe.
	fileNames []string                     // Support mutil-files.
	data      map[string]map[string]string // Section -> key : value

	// Lists can keep sections and keys in order.
	sectionList []string            // Section name list.
	keyList     map[string][]string // Section -> Key name list
}

//LoadConfigFile 配置文件并读取内容
func LoadConfigFile(fileName string, moreFiles ...string) (c *ConfigFile, err error) {
	fileNames := make([]string, 1, len(moreFiles)+1)
	fileNames[0] = fileName
	if len(moreFiles) > 0 {
		fileNames = append(fileNames, moreFiles...)
	}

	c = new(ConfigFile)
	c.fileNames = fileNames
	c.data = make(map[string]map[string]string)
	c.keyList = make(map[string][]string)

	for _, name := range fileNames {
		if err = c.loadFile(name); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *ConfigFile) loadFile(fileName string) (err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.read(f)
}

// 读取一个配置文件内容
func (c *ConfigFile) read(reader io.Reader) (err error) {
	buf := bufio.NewReader(reader)

	count := 1
	section := DEFAULT_SECTION
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineLengh := len(line) //[SWH|+]
		if err != nil {
			if err != io.EOF {
				return err
			}

			if lineLengh == 0 {
				break
			}
		}

		// switch written for readability (not performance)
		switch {
		case lineLengh == 0: // Empty line
			continue
		case line[0] == '#' || line[0] == ';': // Comment
			continue
		case line[0] == '[' && line[lineLengh-1] == ']': // New sction.
			// Get section name.
			section = strings.TrimSpace(line[1 : lineLengh-1])
			// Make section exist even though it does not have any key.
			c.SetValue(section, " ", " ")
			// Reset counter.
			count = 1
			continue
		case section == "": // No section defined so far
			return readError{ErrBlankSectionName, line}
		default: // Other alternatives
			var (
				i        int
				keyQuote string
				key      string
				valQuote string
				value    string
			)
			//[SWH|+]:支持引号包围起来的字串
			if line[0] == '"' {
				if lineLengh >= 6 && line[0:3] == `"""` {
					keyQuote = `"""`
				} else {
					keyQuote = `"`
				}
			} else if line[0] == '`' {
				keyQuote = "`"
			}
			if keyQuote != "" {
				qLen := len(keyQuote)
				pos := strings.Index(line[qLen:], keyQuote)
				if pos == -1 {
					return readError{ErrCouldNotParse, line}
				}
				pos = pos + qLen
				i = strings.IndexAny(line[pos:], "=:")
				if i <= 0 {
					return readError{ErrCouldNotParse, line}
				}
				i = i + pos
				key = line[qLen:pos] //保留引号内的两端的空格
			} else {
				i = strings.IndexAny(line, "=:")
				if i <= 0 {
					return readError{ErrCouldNotParse, line}
				}
				key = strings.TrimSpace(line[0:i])
			}
			//[SWH|+];

			// Check if it needs auto increment.
			if key == "-" {
				key = "#" + fmt.Sprint(count)
				count++
			}

			//[SWH|+]:支持引号包围起来的字串
			lineRight := strings.TrimSpace(line[i+1:])
			lineRightLength := len(lineRight)
			firstChar := ""
			if lineRightLength >= 2 {
				firstChar = lineRight[0:1]
			}
			if firstChar == "`" {
				valQuote = "`"
			} else if lineRightLength >= 6 && lineRight[0:3] == `"""` {
				valQuote = `"""`
			}
			if valQuote != "" {
				qLen := len(valQuote)
				pos := strings.LastIndex(lineRight[qLen:], valQuote)
				if pos == -1 {
					return readError{ErrCouldNotParse, line}
				}
				pos = pos + qLen
				value = lineRight[qLen:pos]
			} else {
				value = strings.TrimSpace(lineRight[0:])
			}
			//[SWH|+];

			c.SetValue(section, key, value)
		}

		// Reached end of file.
		if err == io.EOF {
			break
		}
	}
	return nil
}

//SetValue 设置配置值
//@param            section         段名称
//@param            key             键
//@param            value           值
func (c *ConfigFile) SetValue(section, key, value string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Blank section name represents DEFAULT section.
	if len(section) == 0 {
		section = DEFAULT_SECTION
	}

	// Check if section exists.
	if _, ok := c.data[section]; !ok {
		// Execute add operation.
		c.data[section] = make(map[string]string)
		// Append section to list.
		c.sectionList = append(c.sectionList, section)
	}

	// Check if key exists.
	_, ok := c.data[section][key]
	c.data[section][key] = value
	if !ok {
		// If not exists, append to key list.
		c.keyList[section] = append(c.keyList[section], key)
	}
	return !ok
}

//GetValue 获取string配置值
//@param            section         段名称
//@param            key             键
func (c *ConfigFile) GetValue(section, key string) (string, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// Blank section name represents DEFAULT section.
	if len(section) == 0 {
		section = DEFAULT_SECTION
	}

	// Check if section exists
	if _, ok := c.data[section]; !ok {
		// Section does not exist.
		return "", getError{ErrSectionNotFound, section}
	}

	// Section exists.
	// Check if key exists or empty value.
	value, ok := c.data[section][key]
	if !ok {
		// Check if it is a sub-section.
		if i := strings.LastIndex(section, "."); i > -1 {
			return c.GetValue(section[:i], key)
		}

		// Return empty value.
		return "", getError{ErrKeyNotFound, key}
	}

	// Key exists.
	var i int
	for i = 0; i < _DEPTH_VALUES; i++ {
		vr := varPattern.FindString(value)
		if len(vr) == 0 {
			break
		}

		// Take off leading '%(' and trailing ')s'.
		noption := strings.TrimLeft(vr, "%(")
		noption = strings.TrimRight(noption, ")s")

		// Search variable in default section.
		nvalue, err := c.GetValue(DEFAULT_SECTION, noption)
		if err != nil && section != DEFAULT_SECTION {
			// Search in the same section.
			if _, ok := c.data[section][noption]; ok {
				nvalue = c.data[section][noption]
			}
		}

		// Substitute by new value and take off leading '%(' and trailing ')s'.
		value = strings.Replace(value, vr, nvalue, -1)
	}
	return value, nil
}

//Bool 获取bool配置值
//@param            section         段名称
//@param            key             键
func (c *ConfigFile) Bool(section, key string) (bool, error) {
	value, err := c.GetValue(section, key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(value)
}

//Float64 获取float64配置值
//@param            section         段名称
//@param            key             键
func (c *ConfigFile) Float64(section, key string) (float64, error) {
	value, err := c.GetValue(section, key)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(value, 64)
}

//Int 获取int配置值
//@param            section         段名称
//@param            key             键
func (c *ConfigFile) Int(section, key string) (int, error) {
	value, err := c.GetValue(section, key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

//Int64 获取int64配置值
//@param            section         段名称
//@param            key             键
func (c *ConfigFile) Int64(section, key string) (int64, error) {
	value, err := c.GetValue(section, key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

// 取值出错
type getError struct {
	Reason int
	Name   string
}

func (err getError) Error() string {
	switch err.Reason {
	case ErrSectionNotFound:
		return fmt.Sprintf("section '%s' not found", err.Name)
	case ErrKeyNotFound:
		return fmt.Sprintf("key '%s' not found", err.Name)
	}
	return "invalid get error"
}

// 读配置文件出错
type readError struct {
	Reason  int
	Content string
}

func (err readError) Error() string {
	switch err.Reason {
	case ErrBlankSectionName:
		return "empty section name not allowed"
	case ErrCouldNotParse:
		return fmt.Sprintf("could not parse line: %s", string(err.Content))
	}
	return "invalid read error"
}
