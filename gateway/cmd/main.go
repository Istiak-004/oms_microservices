package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	cnf "github.com/istiak-004/gateway/config"
	"github.com/istiak-004/gateway/internals/routers"
	//_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := cnf.LoadEnv()

	engine := routers.NewRouter(gin.Default())
	engine.SetupRoutes()

	fmt.Printf("Starting server at port %v", config.PortAddr)
	engine.Run(config.PortAddr) // Start the server on port 8080
}
