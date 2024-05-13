package server

import (
	cognito "github.com/akhettar/gin-jwt-cognito"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/config"
	"github.com/ag7if/wendover/internal/database"
	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/internal/repositories/postgresrepo"
	"github.com/ag7if/wendover/internal/server/handlers"
)

type Handler interface {
	SetupRoutes(router *gin.RouterGroup, auth gin.HandlerFunc)
}

type Server struct {
	engine *gin.Engine
	repo   repositories.Repository
	auth   *cognito.AuthMiddleware
}

func setupRepository() (repositories.Repository, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	repo := postgresrepo.NewPostgresRepository(db)

	return repo, nil
}

func NewServer(version string) (*Server, error) {
	repo, err := setupRepository()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	/*
		iss := viper.GetString(config.AWSCognitoIss)
		userPoolID := viper.GetString(config.AWSCognitoUserpoolID)
		region := viper.GetString(config.AWSRegion)

		auth, err := cognito.AuthJWTMiddleware(iss, userPoolID, region)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	*/

	engine := gin.Default()

	server := &Server{
		repo:   repo,
		engine: engine,
		//auth:   auth,
	}
	server.setupRoutes(version)

	return server, nil
}

func (s *Server) setupRoutes(version string) {
	rootPath := viper.GetString(config.ServerRootPath)
	v1 := s.engine.Group(rootPath)

	healthcheckHandler := handlers.NewHealthcheckHandler(version)
	healthcheckHandler.SetupRoutes(v1, s.auth.MiddlewareFunc())

	activityHandler := handlers.NewActivityHandler(s.repo)
	activityHandler.SetupRoutes(v1, s.auth.MiddlewareFunc())

	userHandler := handlers.NewUserHandler(s.repo)
	userHandler.SetupRoutes(v1, s.auth.MiddlewareFunc())
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
