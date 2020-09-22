package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DropKit/DropKit-Adapter/logger"
	routes "github.com/DropKit/DropKit-Adapter/router"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/ethereum/go-ethereum/log"
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
	// router := routes.SetupRouter()
	_, srv := routes.SetupRouter()
	// router.Run(":" + viper.GetString(`DROPKIT.PORT`))

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Forced to shutdown server")
	}

}
