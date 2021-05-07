package config

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	config := &Config{
		Path:  "../../../config",
		Name:  "config",
		Type:  "yaml",
		Mode:  "dev",
		Watch: false,
	}
	vp := New(config)
	log.Println(vp.Cfg.Get("APP.mode"))
	log.Println(vp.Cfg.GetString("DB.dbname"))
}
