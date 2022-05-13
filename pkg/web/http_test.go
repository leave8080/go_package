package web

import (
	"github/leave8080/go_package/pkg/limit"
	"github/leave8080/go_package/pkg/log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInitServer(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skip TestInitServer")
	}
	c := &Config{
		Port: ":8080",
		Limit: &limit.Config{
			Rate:       1, // 0 速率不限流
			BucketSize: 10,
		},
	}

	g := InitServer(c)
	initRoute(g.Gin)
	g.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Warningf("get a signal %s, stop the process", si.String())
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func initRoute(g *gin.Engine) {
	g.GET("/a", func(c *gin.Context) {
		c.JSON(200, "a")
	})
	g.GET("/b", func(c *gin.Context) {
		c.JSON(200, "b")
	})
	g.GET("/c", func(c *gin.Context) {
		c.JSON(200, "c")
	})
	g.GET("/d", func(c *gin.Context) {
		c.JSON(200, "d")
	})
}
