package main

import (
	"github.com/j4real2208/banking/app"
	"github.com/j4real2208/banking/logger"
)

func main() {
	//log.Println("Started from the Main.go")
	logger.Info("Starting the application and serving on port 8000 .... ")
	app.Start()
}
