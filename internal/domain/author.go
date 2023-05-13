package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:generate mockgen -source=author.go -destination=mocks/author_mocks.go -package=mocks

type AuthorEntity struct {
	Id          uuid.UUID
	Name        string
	CreatedDate time.Time
}

type AuthorRepository interface {
	Get(id uuid.UUID) (AuthorEntity, error)
	Create(author AuthorEntity) (AuthorEntity, error)
	Delete(id uuid.UUID) error
	Exists(id uuid.UUID) (bool, error)
	ExistsByName(name string) (bool, error)
	GetAll() ([]AuthorEntity, error)
}

type AuthorService interface {
	Get(id uuid.UUID) (AuthorEntity, error)
	Create(author AuthorEntity) (AuthorEntity, error)
	Delete(id uuid.UUID) error
	GetAll() ([]AuthorEntity, error)
}

type AuthorHandler interface {
	Create() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Get() gin.HandlerFunc
	GetAll() gin.HandlerFunc
}
