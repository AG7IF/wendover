package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/internal/server/views"
)

type ActivityHandler struct {
	repo repositories.Repository
}

func NewActivityHandler(repo repositories.Repository) ActivityHandler {
	return ActivityHandler{repo: repo}
}

func (ah *ActivityHandler) Create(c *gin.Context) {
	av := &views.ActivityView{}

	err := c.BindJSON(av)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to bind activity view")
		ResolveError(c, err)
		return
	}

	a := av.ToDomainObject()

	res, err := ah.repo.InsertActivity(a)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to create new activity")
		ResolveError(c, err)
		return
	}

	rav := views.NewActivityView(res)

	c.JSON(http.StatusCreated, rav)
}

func (ah *ActivityHandler) FetchAll(c *gin.Context) {
	// TODO: Query by authentication parameters
	a, err := ah.repo.SelectActivities()
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to query activities")
		ResolveError(c, err)
	}

	avs := make([]views.ActivityView, 0)

	for _, v := range a {
		av := views.NewActivityView(v)
		avs = append(avs, av)
	}

	c.JSON(http.StatusOK, avs)
}

func (ah *ActivityHandler) Fetch(c *gin.Context) {
	key := strings.ToUpper(c.Param("key"))

	a, err := ah.repo.SelectActivity(key)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to query activity")
		ResolveError(c, err)
		return
	}

	av := views.NewActivityView(a)

	c.JSON(http.StatusOK, av)
}

func (ah *ActivityHandler) Update(c *gin.Context) {
	key := strings.ToUpper(c.Param("key"))
	av := &views.ActivityView{}
	err := c.BindJSON(av)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to bind activity view")
		ResolveError(c, err)
		return
	}

	a := av.ToDomainObject()

	res, err := ah.repo.UpdateActivity(key, a)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to update activity")
		ResolveError(c, err)
		return
	}

	rav := views.NewActivityView(res)

	c.JSON(http.StatusOK, rav)
}

func (ah *ActivityHandler) Delete(c *gin.Context) {
	key := strings.ToUpper(c.Param("key"))
	err := ah.repo.DeleteActivity(key)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to delete activity")
		ResolveError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("activity %s deleted", key)})
}

func (ah *ActivityHandler) SetupRoutes(router *gin.RouterGroup, auth gin.HandlerFunc) {
	router.POST("/activity", auth, ah.Create)
	router.GET("/activity", auth, ah.FetchAll)
	router.GET("/activity/:key", auth, ah.Fetch)
	router.PUT("/activity/:key", auth, ah.Update)
	router.DELETE("/activity/:key", auth, ah.Delete)
}
