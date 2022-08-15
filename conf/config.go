package conf

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigHTTP struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type ConfigDB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int32  `json:"port"`
}

type Config struct {
	Http      *ConfigHTTP `json:"http"`
	DB        *ConfigDB   `json:"db"`
	AuthToken string      `json:"auth_token"`
}

var (
	ConfigInst *Config
)

func InitConfig(path string) error {
	ConfigInst = new(Config)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, ConfigInst)
	if err != nil {
		return err
	}
	return nil
}
