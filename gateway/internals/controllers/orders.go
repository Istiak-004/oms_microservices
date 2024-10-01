package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/istiak-004/common/api"
)

type Order struct {
	Items []*ItemsWithQuantity `json:"items"`
}

type ItemsWithQuantity struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

func (c *Controller) CreateOrder(ctx *gin.Context) {

	customerId := ctx.Param("customerId")
	// Bind the request body JSON to the Order struct
	var order Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Manually convert your Go struct to the protobuf struct
	var pbItems []*pb.ItemsWithQuantity
	for _, item := range order.Items {
		pbItems = append(pbItems, &pb.ItemsWithQuantity{
			Id:       item.ID,
			Quantity: int32(item.Quantity), // Ensure correct types
		})
	}

	// Construct the gRPC request
	orderReq := pb.CreateOrderRequest{
		CustomerId: customerId,
		Items:      pbItems, // Use the converted items
	}

	// Call the gRPC service
	grpcCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := c.grpcClient.CreateOrder(grpcCtx, &orderReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the gRPC response to the HTTP client
	ctx.JSON(http.StatusOK, gin.H{
		"order_id": resp.OrderId,
		"status":   resp.Status,
	})
}
