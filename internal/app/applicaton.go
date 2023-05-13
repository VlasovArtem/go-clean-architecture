package app

import (
	"github.com/gin-gonic/gin"

	"clean-architecture/internal/config"
)

type Application interface {
	Run() error
}

type application struct {
	cfg                 config.Config
	dependenciesManager *dependenciesManager
}

func NewApplication() (Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return &application{
		cfg:                 cfg,
		dependenciesManager: newDependenciesManager(),
	}, nil
}

func (a *application) Run() error {
	router := gin.Default()

	a.initRouter(router)

	err := router.Run(":8080")

	return err
}

func (a *application) initRouter(router *gin.Engine) {
	handler := a.dependenciesManager.authorHandler

	router.Group("/api/author").
		GET("/:id", handler.Get()).
		POST("/", handler.Create()).
		DELETE("/:id", handler.Delete())

	router.Group("/api/authors").
		GET("/", handler.GetAll())
}
