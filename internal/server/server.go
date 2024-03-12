package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/database"
	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/internal/repositories/postgresrepo"
	"github.com/ag7if/wendover/internal/server/handlers"
)

type Server struct {
	engine *gin.Engine
	repo   repositories.Repository
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

	engine := gin.Default()

	server := &Server{
		repo:   repo,
		engine: engine,
	}
	server.setupRoutes(engine, version)

	return server, nil
}

func (s *Server) setupRoutes(engine *gin.Engine, version string) {
	v1 := engine.Group("/api/v1")

	healthcheckHandler := handlers.NewHealthcheckHandler(version)
	healthcheckHandler.SetupRoutes(v1)

	activityHandler := handlers.NewActivityHandler(s.repo)
	activityHandler.SetupRoutes(v1)

	userHandler := handlers.NewUserHandler(s.repo)
	userHandler.SetupRoutes(v1)
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
