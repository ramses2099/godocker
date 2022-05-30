package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramses2099/godocker/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	log.Println("Application running in enviroment: ",
		config.RuntimeSetup, " and on port: ", config.AppPort)

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")

	router.Run(config.ServerAddress + ":" + config.AppPort)
}
