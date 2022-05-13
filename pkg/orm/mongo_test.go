package orm

import (
	"context"
	"fmt"
	"github/leave8080/go_package/pkg/log"
	"testing"
	"time"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func Test_InitMongo(t *testing.T) {
	t.Skip()

	config := &MongoConfig{
		MongoUri:       "mongodb://time_machine:mxchip%402019@106.14.218.147:27017/time_machine",
		Database:       "time_machine",
		PoolSize:       100,
		ConnectTimeOut: 10,
	}
	client := InitMongo(config)
	res, _ := client.ListDatabases(context.Background(), nil)
	cnt, err := client.Database("time_machine").Collection("test").InsertOne(context.Background(), Trainer{Name: "hello2"})
	if err != nil {
		log.Error(err)
	}
	fmt.Println(cnt, res)
	MongoClose(client, time.Second*time.Duration(10))
}
