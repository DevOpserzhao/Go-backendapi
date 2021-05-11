package tools

import (
	"strings"
)

func EmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func EmptySlice(sli []interface{}) bool {
	return len(sli) == 0
}

func EmptyMap(m map[string]interface{}) bool {
	return len(m) == 0
}
