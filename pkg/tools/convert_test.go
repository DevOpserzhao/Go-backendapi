package tools

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestStructToMap(t *testing.T) {
	tests := []struct {
		Code int
		Msg  string
	}{
		{Code: 1, Msg: "Success"},
		{Code: 0, Msg: "Failed"},
	}
	wants := []map[string]interface{}{
		{"Code": 1, "Msg": "Success"},
		{"Code": 0, "Msg": "Failed"},
	}
	t.Helper()
	for index, tt := range tests {
		t.Run("StructToMap", func(t *testing.T) {
			assert.Equal(t, StructToMap(tt), wants[index])
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		S    string
		Want int
		Name string
	}{
		{S: "1", Want: 1, Name: "first"},
		{S: "2", Want: 2, Name: "second"},
		{S: "3", Want: 3, Name: "third"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			i, ok := StringToInt(tt.S)
			if ok {
				assert.Equal(t, i, tt.Want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		I    int
		Want string
		Name string
	}{
		{I: 99, Want: "99", Name: "first"},
		{I: 0, Want: "0", Name: "second"},
		{I: -1, Want: "-1", Name: "third"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, IntToString(tt.I), tt.Want)
		})
	}
}

func TestIntToBool(t *testing.T) {
	tests := []struct {
		I    int
		Want bool
		Name string
	}{
		{I: 0, Want: false, Name: "first"},
		{I: 1, Want: true, Name: "second"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, IntToBool(tt.I), tt.Want)
		})
	}
}

func TestBoolToInt(t *testing.T) {
	tests := []struct {
		B    bool
		Want int
		Name string
	}{
		{B: true, Want: 1, Name: "first"},
		{B: false, Want: 0, Name: "second"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, BoolToInt(tt.B), tt.Want)
		})
	}
}

func TestJoinStrings(t *testing.T) {
	tests := []struct{
		S []string
		Want string
		Name string
	}{
		{S: []string{"user", ":", "password"}, Want: "user:password", Name: "first"},
		{S: []string{"Hello", " ", "World!"}, Want: "Hello World!", Name: "second"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, JoinStrings(tt.S...), tt.Want)
		})
	}
}

func BenchmarkJoinStrings(b *testing.B) {
	test := struct{
		S []string
		Want string
		Name string
	}{
		S: []string{"user", ":", "password"}, Want: "user:password", Name: "first",
	}
	b.Helper()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JoinStrings(test.S...)
		}
	})
}