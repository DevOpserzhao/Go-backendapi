package config

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Path  string
	Name  string
	Type  string
	Mode  string
	Watch bool
}

type Viper struct {
	Cfg *viper.Viper
}

func NewConfig() *Config {
	var path string
	var name string
	var t string
	var mode string
	var watch bool
	// DEBUG Used "../backend/config"
	flag.StringVar(&path, "p", "../backend/config", "配置文件目录")
	flag.StringVar(&name, "n", "config", "配置文件名字")
	flag.StringVar(&t, "t", "yaml", "配置文件类型")
	flag.StringVar(&mode, "m", "dev", "运行模式")
	flag.BoolVar(&watch, "w", false, "配置文件改动监听")
	flag.Parse()
	return &Config{
		Path:  path,
		Name:  name,
		Type:  t,
		Mode:  mode,
		Watch: watch,
	}
}

func New(config *Config) *Viper {
	return &Viper{Cfg: mustLoad(config)}
}

func mustLoad(c *Config) *viper.Viper {
	vp := viper.New()
	vp.AddConfigPath(c.Path)
	vp.SetConfigName(c.Name)
	vp.SetConfigType(c.Type)
	if err := vp.ReadInConfig(); err != nil {
		log.Println("配置文件加载错误", err.Error())
		os.Exit(-1)
	}
	if c.Watch {
		vp.OnConfigChange(func(in fsnotify.Event) {
			log.Println("配置文件修改")
		})
		vp.WatchConfig()
	}
	return vp
}
