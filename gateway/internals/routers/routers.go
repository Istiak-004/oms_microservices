package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/istiak-004/gateway/internals/controllers"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		router: engine,
	}
}

func (r *Router) SetupRoutes() {
	router := r.router
	c := controllers.NewController()
	orders := router.Group("/orders")
	{
		orders.POST("/orders", c.CreateOrder)
	}
}

func (r *Router) Run(addr string) {
	r.router.Run(addr)
}
