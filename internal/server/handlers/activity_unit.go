package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ag7if/wendover/internal/repositories"
)

type ActivityUnitHandler struct {
	repo repositories.Repository
}

func NewActivityUnitHandler(repo repositories.Repository) ActivityUnitHandler {
	return ActivityUnitHandler{repo: repo}
}

func (a *ActivityUnitHandler) Options(c *gin.Context) {
	setHeaders(c)
	c.Header("Allow", "GET,POST,PUT,DELETE,OPTIONS")
	c.AbortWithStatus(http.StatusNoContent)
}

func (a *ActivityUnitHandler) Create(c *gin.Context) {
	setHeaders(c)

}

func (a *ActivityUnitHandler) FetchAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ActivityUnitHandler) Fetch(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ActivityUnitHandler) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ActivityUnitHandler) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ActivityUnitHandler) SetupRoutes(router *gin.RouterGroup) {
	router.OPTIONS("/activity/:key/units", a.Options)
	router.POST("/activity/:key/units", a.Create)
	router.GET("/activity/:key/units", a.FetchAll)
	router.GET("/activity/:key/units/:id", a.Fetch)
	router.PUT("/activity/:key/units/:id", a.Update)
	router.DELETE("/activity/:key/units/:id", a.Delete)
}
