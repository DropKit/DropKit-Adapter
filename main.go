package main

import (
	"net/http"

	routes "github.com/DropKit/DropKit-Adapter/router"
	"github.com/DropKit/DropKit-Adapter/services"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		print(err)
	}

	services.QuorumAlive()
	services.DBAlive()
}

func main() {

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}
