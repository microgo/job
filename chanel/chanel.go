package chanel

import (
	"github.com/streadway/amqp"
	"job/config"
	"job/service"
	"master/resource/helper"
	"master/utils"
	"strconv"
)

const (
	JobQueueName = "tasks"
)

type Channel struct {
	Helper    *helper.Helper
	JobChanel *amqp.Channel
}

func (c *Channel) InitChanel() {
	JobChanel, err := c.Helper.MakeChanel(JobQueueName)
	if err != nil {
		panic(err)
		return
	}
	c.JobChanel = JobChanel
}

func (c *Channel) Consume() {
	service := service.Service{
		Helper: c.Helper,
	}
	forever := make(chan bool)
	for i := 1; i <= config.NumberJobConcurrent; i++ {
		utils.LogInfo("Job worker", i, "started")
		go c.Helper.MakeConsumeWithTag(JobQueueName, strconv.Itoa(i), c.JobChanel, service.JobHandler)
	}
	<-forever
}
