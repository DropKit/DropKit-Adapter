package main

import (
	"net/http"

	"github.com/DropKit/DropKit-Adapter/logger"
	routes "github.com/DropKit/DropKit-Adapter/router"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.InternalLogger.WithField("component", "main").Error(err.Error())
	}

	services.DependencyServicesCheck()
}

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":5000", router)

}
