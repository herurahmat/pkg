package strings

import (
	"fmt"
	"reflect"
)

type MaskingType string

const (
	Any           MaskingType = "any"
	Left          MaskingType = "left"
	Middle        MaskingType = "middle"
	Right         MaskingType = "right"
	DefaultMasked             = "******"
)

func MaskAny(s string) string {
	if s == "" {
		return ""
	}
	return DefaultMasked
}

func MaskLeft(s string) string {
	if s == "" {
		return ""
	}
	if len(s) < 12 {
		return fmt.Sprintf("%s%s", DefaultMasked, s[len(s)/2+1:])
	}
	return fmt.Sprintf("%s%s", DefaultMasked, s[len(s)-6:])
}

func MaskMiddle(s string) string {
	if s == "" {
		return ""
	}
	if len(s) < 12 {
		return fmt.Sprintf("***%s***", s[len(s)/2-len(s)/3+1:len(s)/2+len(s)/3])
	}
	return fmt.Sprintf("***%s***", s[len(s)/2-3:len(s)/2+3])
}

func MaskRight(s string) string {
	if s == "" {
		return ""
	}
	if len(s) < 12 {
		return fmt.Sprintf("%s%s", s[:len(s)/2], DefaultMasked)
	}
	return fmt.Sprintf("%s%s", s[:6], DefaultMasked)
}

func Mask(t MaskingType, s string) string {
	switch t {
	case Any:
		return MaskAny(s)
	case Left:
		return MaskLeft(s)
	case Middle:
		return MaskMiddle(s)
	case Right:
		return MaskRight(s)
	}
	return s
}

func MaskStruct(s any) {
	e := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < e.NumField(); i++ {
		value, ok := e.Field(i).Tag.Lookup("mask")
		if ok && v.Field(i).CanSet() {
			v.Field(i).SetString(Mask(MaskingType(value), v.Field(i).Interface().(string)))
		}
	}
}
