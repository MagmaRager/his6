package database

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"newhis/base/convert"
	"strconv"
)

type NullableString struct {
	sql.NullString
}

type NullableInt64 struct {
	sql.NullInt64
}

func (ns NullableString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return convert.Str2Byte("null"), nil
}

func (ns *NullableString) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == 34 {
		var str string
		err := json.Unmarshal(b, &str)
		if err != nil {
			return err
		}

		ns.Valid = true
		ns.String = str
	} else {
		ns.Valid = false
	}
	return nil
}

func (ns NullableInt64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int64)
	}
	return convert.Str2Byte("null"), nil
}

func (ns *NullableInt64) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		var buf bytes.Buffer
		buf.WriteByte(34)
		buf.Write(b)
		buf.WriteByte(34)
		var str string
		err := json.Unmarshal(buf.Bytes(), &str)
		if err != nil {
			return err
		}

		ns.Valid = true
		i, errc := strconv.ParseInt(str, 10, 64)
		if errc != nil {
			return errc
		}
		ns.Int64 = i
	} else {
		ns.Valid = false
	}
	return nil
}
