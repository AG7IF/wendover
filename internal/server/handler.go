package server

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Options(c *gin.Context)
	FetchAll(c *gin.Context)
	Fetch(c *gin.Context)
	SetupRoutes(router *gin.RouterGroup)
}

type CreateHandler interface {
	Create(c *gin.Context)
}

type UpdateHandler interface {
	Update(c *gin.Context)
}

type DeleteHandler interface {
	Delete(c *gin.Context)
}
