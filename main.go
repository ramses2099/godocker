package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ramses2099/godocker/util"
	"github.com/spf13/viper"
)

var config *util.Config

func init() {
	myconfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	config = &myconfig

	// load book collection
	// util.LoadCollection("./static/")

	util.InitializeRedis()
}

func main() {

	// logs
	setUpLogging()
	//
	// _, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config: ", err)
	// }

	// log.Println("Application running in enviroment: ",
	// 	config.RuntimeSetup, " and on port: ", config.AppPort)
	log.Println("Application running in enviroment: ",
		viper.GetString("RUNTIME_SETUP"), " and on port: ", viper.GetInt("PORT"))

	var router *gin.Engine
	router = gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/collection", getColletion)

	//router.Run(config.ServerAddress + ":" + config.AppPort)
	router.Run(viper.GetString("SERVER_ADDRESS") + ":" + viper.GetString("PORT"))
}

func setUpLogging() {
	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func getColletion(ctx *gin.Context) {
	val := util.GetBookList()
	ctx.HTML(http.StatusOK, "library.html", gin.H{
		"books": val.BookList,
	})
}
