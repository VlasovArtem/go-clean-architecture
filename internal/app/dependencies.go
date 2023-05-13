package app

import (
	"clean-architecture/internal/controller/rest"
	"clean-architecture/internal/domain"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/service"
)

type dependenciesManager struct {
	authorRepository domain.AuthorRepository
	authorService    domain.AuthorService
	authorHandler    domain.AuthorHandler
}

func newDependenciesManager() *dependenciesManager {
	authorRepository := repository.NewAuthorRepository()
	authorService := service.NewAuthorService(authorRepository)
	return &dependenciesManager{
		authorRepository: authorRepository,
		authorService:    authorService,
		authorHandler:    rest.NewAuthorHandler(authorService),
	}
}
