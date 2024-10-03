package routers

import (
	"github.com/gin-gonic/gin"
	pb "github.com/istiak-004/common/api"
	"github.com/istiak-004/gateway/internals/controllers"
)

type Router struct {
	router     *gin.Engine
	grpcClient pb.OrderServiceClient
}

func NewRouter(engine *gin.Engine, grpcClient pb.OrderServiceClient) *Router {
	return &Router{
		router:     engine,
		grpcClient: grpcClient,
	}
}

func (r *Router) SetupRoutes() {
	router := r.router
	c := controllers.NewController(r.grpcClient)
	orders := router.Group("/orders")
	{
		orders.POST("/customer/:customerId", c.CreateOrder)
	}
}

func (r *Router) Run(addr string) {
	r.router.Run(addr)
}
