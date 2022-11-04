package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	_ "store.getAway/config"
	apiRouter "store.getAway/router"
)

func main() {
	var appPort = viper.GetInt("application.port")
	router := apiRouter.InitRouter()
	if err := router.Run(fmt.Sprintf(":%d", appPort)); err != nil {
		log.Fatal(err.Error())
	}
}
