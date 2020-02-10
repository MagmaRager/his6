package convert

import (
"reflect"
"time"
"unsafe"
)

func LongString(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}

func DateTimeString(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func DateString(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Str2Byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return !vi.IsValid() || vi.IsNil()
}
