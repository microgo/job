package main

import (
	"job/chanel"
	"master/resource"
	"master/resource/helper"
	"master/utils"
)

func main() {
	config := resource.ResourceConfig{
		IsEnableRabbit: true,
	}

	r, err := resource.Init(config)
	if err != nil {
		utils.LogError("Connect resource fail, app will be shutdown...", err)
		return
	}
	defer r.Close()
	utils.LogInfo("Service dede is running...")
	h := helper.Helper{Resource: r}
	chanel := chanel.Channel{Helper: &h}
	chanel.InitChanel()
	chanel.Consume()
}
