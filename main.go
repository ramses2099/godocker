package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramses2099/godocker/logging"
	"github.com/ramses2099/godocker/util"
	"github.com/spf13/viper"
)

var config *util.Config

func init() {
	myconfig, err := util.LoadConfig(".")
	if err != nil {
		logging.AppLog.WriteLogsError("cannot load config:",
			map[string]interface{}{"source": config, "error": err})
	}
	config = &myconfig

	// load book collection
	// util.LoadCollection("./static/")

	util.InitializeRedis()
}

func main() {

	// logs
	logging.SetUpLogging()
	//
	// _, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config: ", err)
	// }

	logging.AppLog.WriteLogsInfo("Application running in enviroment: ", map[string]interface{}{"runtime_setup": viper.GetString("RUNTIME_SETUP"),
		"app_port": viper.GetInt("PORT")})

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

func getColletion(ctx *gin.Context) {
	val := util.GetBookList()
	ctx.HTML(http.StatusOK, "library.html", gin.H{
		"books": val.BookList,
	})
}
