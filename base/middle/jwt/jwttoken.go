package jwt

import (
	"encoding/base64"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"his6/base/config"
)

var (
	//  JWT数字签名密码
	key []byte
	//  jwt有效时间秒数
	expiresDuration int64
	//
	excludeUrlPattern []string
	//  JWT名称
	name string
)

func init() {
	expiresDuration = config.GetConfigInt64("jwt", "expDuration", 1800)

	sk := config.GetConfigString("jwt", "key", "9588028E20109149")
	key, _ = base64.StdEncoding.DecodeString(sk)

	name = config.GetConfigString("jwt", "name", "JWT_TOKEN")
}

//  记录Token信息
type Info struct {
	//  登录路径，区分不同类型的登录
	loginWay string
	// 分支机构
	branchId string
	// 人员
	empId string
	// IP
	loginIp string
	// 权限列表串
	authString string
}

//  产生JWT信息实体
func CreateToken(brc, emp, ip, lgw, aut string) Info {
	return Info{
		branchId: brc,
		empId:    emp,
		loginIp:  ip,
		loginWay: lgw,
		authString: aut,
	}
}

//  获取JWT名称
func GetName() string {
	return name
}

//  获取IP地址
func (info Info) GetIp() string {
	return info.loginIp
}

//  获取权限列表
func (info Info) GetAuth() string {
	return info.authString
}

//  获取权限列表
func (info Info) ToString() string {
	s := "[机构:" + info.branchId + " "
	s += "人员ID:" + info.empId + " "
	s += "IP地址:" + info.loginIp + " "
	s += "拥有权限:" + info.authString + "]"
	return s
}

//  获取客户端IP
func GetClientIp(r *http.Request) string {
	ip := strings.TrimSpace(r.Header.Get("HTTP_CLIENT_IP"))
	if ip != "" {
		return ip
	}

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// 产生json web token
func (info Info) GenToken() (string, error) {
	claims := jwt.MapClaims{
		"brc": info.branchId,
		"emp": info.empId,
		"lgw": info.loginWay,
		"ip":  info.loginIp,
		"exp": time.Now().Unix() + expiresDuration,
		"iat": time.Now().Unix(),
		"aut": info.authString,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// 校验token是否有效, 返回token值或错误
func CheckToken(token string) (Info, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	jt := Info{}
	if err != nil {
		return jt, err
	}

	mt := (t.Claims).(jwt.MapClaims)
	//if time.Now().Unix() > int64(mt["exp"].(float64) * math.Pow10(0)) {
	//	err = errors.New("token已经过期")
	//	return jt, err
	//}
	jt.authString = mt["aut"].(string)
	jt.branchId = mt["brc"].(string)
	jt.empId = mt["emp"].(string)
	jt.loginWay = mt["lgw"].(string)
	jt.loginIp = mt["ip"].(string)


	return jt, err
}
