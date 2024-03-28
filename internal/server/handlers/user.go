package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/internal/server/views"
	"github.com/ag7if/wendover/pkg/auth"
)

type UserHandler struct {
	repo repositories.Repository
}

func NewUserHandler(repo repositories.Repository) UserHandler {
	return UserHandler{repo: repo}
}

func (uh *UserHandler) Create(c *gin.Context) {
	av := &views.UserView{}

	err := c.BindJSON(av)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to bind user view")
		ResolveError(c, err)
		return
	}

	u := av.ToDomainObject()

	res, err := uh.repo.InsertUser(u)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to create new user")
		ResolveError(c, err)
		return
	}

	ruv := views.NewUserView(res)

	c.JSON(http.StatusCreated, ruv)
}

func (uh *UserHandler) FetchAll(c *gin.Context) {
	a, err := uh.repo.SelectUsers()
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to query users")
		ResolveError(c, err)
	}

	uvs := make([]views.UserView, 0)

	for _, v := range a {
		uv := views.NewUserView(v)
		uvs = append(uvs, uv)
	}

	c.JSON(http.StatusOK, uvs)
}

func (uh *UserHandler) Fetch(c *gin.Context) {
	username := c.Param("username")

	u, err := uh.repo.SelectUser(username)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to query user")
		ResolveError(c, err)
		return
	}

	uv := views.NewUserView(u)

	c.JSON(http.StatusOK, uv)
}

func (uh *UserHandler) Update(c *gin.Context) {
	username := c.Param("username")
	uv := views.UserView{}
	err := c.BindJSON(&uv)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to bind user view")
		ResolveError(c, err)
		return
	}

	u := uv.ToDomainObject()

	res, err := uh.repo.UpdateUser(username, u)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to update user")
		ResolveError(c, err)
		return
	}

	ruv := views.NewUserView(res)

	c.JSON(http.StatusOK, ruv)
}

func (uh *UserHandler) Delete(c *gin.Context) {
	username := c.Param("username")
	err := uh.repo.DeleteUser(username)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to delete user")
		ResolveError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("user %s deleted", username)})
}

func (uh *UserHandler) AddActivityRole(c *gin.Context) {
	username := c.Param("username")

	var role views.RoleActivityKeyView
	err := c.BindJSON(&role)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to update user roles")
		ResolveError(c, err)
		return
	}

	if role.Role == auth.NilRole {
		err := views.ErrInvalidValue{
			FieldName: "role",
			ValidValues: []string{
				"DIRECTOR",
				"ADMIN",
				"STAFF",
			},
		}
		log.Error().Stack().Err(err).Msg("invalid role passed in request body")
		ResolveError(c, err)
		return
	}

	user, err := uh.repo.AddUserRole(username, role.ActivityKey, role.Role)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to update user roles")
		ResolveError(c, err)
		return
	}

	updated := views.NewUserView(user)

	c.JSON(http.StatusOK, updated)
}

func (uh *UserHandler) RemoveActivityRole(c *gin.Context) {
	username := c.Param("username")
	activityKey := strings.ToUpper(c.Param("key"))

	user, err := uh.repo.RemoveUserRole(username, activityKey)
	if err != nil {
		log.Error().Stack().Err(err).Msg("failed to update user roles")
		ResolveError(c, err)
		return
	}

	updated := views.NewUserView(user)

	c.JSON(http.StatusOK, updated)
}

func (uh *UserHandler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/user", uh.Create)
	router.GET("/user", uh.FetchAll)
	router.GET("/user/:username", uh.Fetch)
	router.PUT("/user/:username", uh.Update)
	router.DELETE("/user/:username", uh.Delete)
	router.POST("/user/:username/activities", uh.AddActivityRole)
	router.DELETE("/user/:username/activities/:key", uh.RemoveActivityRole)
}
