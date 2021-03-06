package conf

import (
	"github.com/google/wire"
	"github.com/jinzhu/configor"
)

var ConfigSet = wire.NewSet(GetConfig)

type Config struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug"`
	DB    db     `yaml:"db"`
	Log   log    `yaml:"log"`
}

type db struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DbName       string `yaml:"dbname"`
	Log          bool   `yaml:"log"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

type log struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

var config Config

func InitConfig(cfgPath string) error {
	err := configor.Load(&config, cfgPath)
	return err
}

func GetConfig() Config {
	return config
}
