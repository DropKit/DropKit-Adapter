package main

import (
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

	service, err := services.DependencyServicesCheck()
	if err != nil {
		logger.FatalDependencyService(service, err)
		return
	}

	logger.InfoDependencyService()
}

func main() {
	router := routes.SetupRouter()
	router.Run(":" + viper.GetString(`DROPKIT.PORT`))
}
