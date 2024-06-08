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
	/*
		setHeaders(c)
		activityKey := strings.ToUpper(c.Param("key"))
		superiorUnitID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			err = views.ErrInvalidValue{
				FieldName: "superior_unit_id",
			}
			log.Error().Err(err).Str("superior_unit_id", c.Param("id")).Msg("failed to parse ID")
			ResolveError(c, err)
			return
		}

		auv := &views.ActivityUnitView{}

		err := c.BindJSON(auv)
		if err != nil {
			log.Error().Err(err).Msg("failed to bind activity unit view")
			ResolveError(c, err)
			return
		}

		au := auv.ToDomainObject()

		res, err := a.repo.InsertActivityUnit()
	*/
	//TODO implement me
	panic("implement me")
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
	router.POST("/activity/:key/units/:id", a.Create)
	router.GET("/activity/:key/units", a.FetchAll)
	router.GET("/activity/:key/units/:id", a.Fetch)
	router.PUT("/activity/:key/units/:id", a.Update)
	router.DELETE("/activity/:key/units/:id", a.Delete)
}
