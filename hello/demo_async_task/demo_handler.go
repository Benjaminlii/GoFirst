package main

import (
	"context"
	"time"

	"github.com/Benjaminlii/go_some_learning/common/logger"
)

type demoHandler struct {}

func (h *demoHandler) HandleMessage(ctx context.Context, taskRequest string) (taskResponse interface{}, err error){
	for i := 0; i < 1000; i++ {
		logger.Infof(ctx, "[HandleMessage] >>>>>>>>>>>>>>>>>>>>>>>>>>msg:%v", taskRequest)
		time.Sleep(time.Millisecond * 10)
	}
	return nil, nil
}
