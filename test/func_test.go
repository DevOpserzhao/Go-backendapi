package test

import (
	"reflect"
	"strings"
	"testing"
)

type M struct{}

func (m *M) Foo(in string) string {
	return strings.ToUpper(in)
}

func TestFunc(t *testing.T) {
	v := reflect.ValueOf(&M{})
	f := v.MethodByName("Foo")
	args := []reflect.Value{reflect.ValueOf("abc")}
	t.Log(f.Call(args))
}
