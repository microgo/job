package service

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"master/model/postgres"
	"master/resource/helper"
)

type Service struct {
	Helper *helper.Helper
}

func (s *Service) JobHandler(message *amqp.Delivery) {
	user := postgres.User{}
	err := json.Unmarshal(message.Body, &user)
	if err != nil {
		return
	}
	fmt.Println("Recived", user)
}
