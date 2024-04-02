package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	version string
}

func NewHealthcheckHandler(version string) HealthcheckHandler {
	return HealthcheckHandler{version: version}
}

func (hch *HealthcheckHandler) Healthcheck(c *gin.Context) {
	response := map[string]string{
		"version": hch.version,
	}

	c.JSON(http.StatusOK, response)
}

func (hch *HealthcheckHandler) SetupRoutes(router *gin.RouterGroup, auth gin.HandlerFunc) {
	router.GET("/healthcheck", hch.Healthcheck)
}
