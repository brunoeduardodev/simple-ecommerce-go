package controllers

import "github.com/gin-gonic/gin"

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func Health(ctx *gin.Context) {
	response := HealthResponse{
		Status:  "OK",
		Version: "0.0.1",
	}

	ctx.JSON(200, response)
}
