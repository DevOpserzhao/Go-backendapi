package tools

import (
	"reflect"
	"strconv"
)

func StructToMap(sct interface{}) map[string]interface{} {
	sctT := reflect.TypeOf(sct)
	sctV := reflect.ValueOf(sct)
	ret := make(map[string]interface{})
	for i := 0; i < sctT.NumField(); i++ {
		ret[sctT.Field(i).Name] = sctV.Field(i).Interface()
	}
	return ret
}

func StringToInt(s string) (int, bool) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return v, true
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func UintToString(i uint) string {
	return strconv.Itoa(int(i))
}

func IntToBool(i int) bool {
	return i > 0
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func JoinStrings(ss ...string) (ret string) {
	for _, s := range ss {
		ret += s
	}
	return
}
