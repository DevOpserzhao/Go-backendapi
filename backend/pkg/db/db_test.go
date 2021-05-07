package db

import (
	"testing"
)

func TestNew(t *testing.T) {
	config := &MySQLConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "liyao961223",
		DbName:   "test",
	}
	db := New(config)
	database, invalid := db.Storage.DB();
	if invalid != nil {
		t.Error(invalid.Error())
	}
	if database != nil {
		if err := database.Ping(); err != nil {
			t.Error(err.Error())
		}
	}
}

func TestSetUp(t *testing.T) {
	config := &MySQLConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "liyao961223",
		DbName:   "test",
	}
	db := New(config)
	type Model struct {
		Name string
		Age  int
	}
	SetUp(db.Storage, Model{})
}

func TestClose(t *testing.T) {
	config := &MySQLConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "liyao961223",
		DbName:   "test",
	}
	db := New(config)
	Close(db.Storage)
}
