package main

import (
	"bflobox-api/internal/handler"
	"bflobox-api/internal/util"
)

// Main entrypoint to the app. Initializes & runs API struct
func main() {

	util.InitLogger()

	config, _ := util.LoadConfig()
	// if err != nil {
	// 	zap.L().Fatal("unable to load configuration", zap.Error(err))
	// }

	server := handler.API{Config: config}
	server.Initialize()
	server.Run()
}
