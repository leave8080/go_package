package config

import (
	"flag"
	"github/leave8080/go_package/pkg/log"
	"github/leave8080/go_package/pkg/orm"
	"github/leave8080/go_package/pkg/web"
	"os"
	"testing"
)

type Config struct {
	Name   string
	Number int
	Web    *web.Config
	MySQL  *orm.MySQLConfig
	Redis  *orm.RedisConfig
}

func TestParse(t *testing.T) {
	c := &Config{}
	flag.Set("conf", "config.yaml")
	os.Setenv("Leave8080_ENV", "dev")
	if err := Parse("Leave8080_ENV", c); err != nil {
		log.Error(err)
		return
	}
	log.Debug(">>")
	log.Debug(c.Name)
	log.Debug(c.Number)
	log.Debug(c.Web)
	log.Debug(c.MySQL)
	log.Debug(c.Redis)
}
