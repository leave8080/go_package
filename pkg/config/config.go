package config

import (
	"errors"
	"flag"
	"os"

	"github.com/spf13/viper"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "conf", "", "config path, example: -conf /config.yaml")
}

// 解析配置文件
func Parse(env string, c interface{}) error {
	flag.Parse()
	if confPath == "" {
		return errors.New("load config file path failed")
	}
	envKey := os.Getenv(env)
	viper.SetConfigFile(confPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.UnmarshalKey(envKey, c)
	if err != nil {
		return err
	}
	return nil
}
