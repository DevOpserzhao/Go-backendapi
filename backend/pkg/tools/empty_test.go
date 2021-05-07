package tools

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestEmptyString(t *testing.T) {
	tests := []struct {
		S    string
		Want bool
		Name string
	}{
		{S: "Hello", Want: true, Name: "first"},
		{S: "", Want: false, Name: "second"},
		{S: " ", Want: false, Name: "third"},
		{S: "World!", Want: true, Name: "fourth"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, IsNotEmptyString(tt.S), tt.Want)
		})
	}
}

func TestIsNotEmptySlice(t *testing.T) {
	tests := []struct {
		Sli  []interface{}
		Want    bool
		Name string
	}{
		{Sli: []interface{}{1, 2}, Want: true, Name: "first"},
		{Sli: []interface{}{1, 2, 3}, Want: true, Name: "second"},
		{Sli: []interface{}{}, Want: false, Name: "third"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, IsNotEmptySlice(tt.Sli), tt.Want)
		})
	}
}

func TestIsNotEmptyMap(t *testing.T) {
	tests := []struct{
		M map[string]interface{}
		Want bool
		Name string
	}{
		{M: map[string]interface{}{"k1": "v1"}, Want: true, Name: "first"},
		{M: map[string]interface{}{"k1": "v1", "k2": "v2"}, Want: true, Name: "second"},
		{M: map[string]interface{}{}, Want: false, Name: "third"},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, IsNotEmptyMap(tt.M), tt.Want)
		})
	}
}