package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Benjaminlii/async_task"
	"github.com/Benjaminlii/async_task/common"
	asyncConfig "github.com/Benjaminlii/async_task/config"
	config "github.com/Benjaminlii/go_some_learning/common/config"
	"github.com/Benjaminlii/go_some_learning/common/logger"
)

const (
	DEMO_TASK_TYPE common.TaskType = "demo"
)

func main() {
	ctx := context.Background()
	config.InitConfig()
	appConfig := config.GetConfig()
	c, err := async_task.Init(
		asyncConfig.WithRedis(&asyncConfig.RedisConfig{
			Address:  appConfig.Redis.Address,
			Password: appConfig.Redis.Password,
		}),
		asyncConfig.WithRocketMQ(&asyncConfig.RocketMQConfig{
			NameServers: appConfig.RocketMQ.NameServers,
			Topic:       appConfig.RocketMQ.Topic,
		}),
		asyncConfig.WithHandler("demo", &demoHandler{}),
	)
	if err != nil {
		logger.Errorf(ctx, "[main] async_task init error, err:%v", err.Error())
		return
	}

	var x interface{} = ""
	taskID, err := c.Create(ctx, DEMO_TASK_TYPE, &x, nil)
	if err != nil {
		logger.Errorf(ctx, "[main] Create error, err:%v", err.Error())
		return
	}

	fmt.Println(taskID)
	time.Sleep(30*time.Second)
}
