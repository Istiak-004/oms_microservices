package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/istiak-004/gateway/internals/routers"
)

func main() {
	port := ":8080" // Replace with your desired port number
	engine := routers.NewRouter(gin.Default())
	engine.SetupRoutes()

	fmt.Printf("Starting server at port %v", port)
	engine.Run(port) // Start the server on port 8080
}
