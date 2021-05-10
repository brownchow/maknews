package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	ES     ES     `yaml:"es"`
	Kafka  Kafka  `yaml:"kafka"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Timeout  int `yaml:"timeout"`
}

type Redis struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	DB      string `yaml:"db"`
	Timeout int `yaml:"timeout"`
}

type ES struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout string `yaml:"timeout"`
	Index   string `yaml:"index"`
}

type Kafka struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout string `yaml:"timeout"`
	Topic   string `yaml:"topic"`
}
var globalConfig *Config

func LoadConfig() *Config{
	config := Config{}
	cfgData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("fail to load config file %v", err)
		return nil
	}
	err = yaml.Unmarshal(cfgData, &config)
	if err != nil {
		fmt.Printf("fail to unmarshal configData %v", err)
	}
	globalConfig = &config
	return globalConfig
}