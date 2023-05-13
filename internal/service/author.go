package service

import (
	"time"

	"github.com/google/uuid"

	"clean-architecture/internal/domain"
	"clean-architecture/internal/errors"
)

type authorService struct {
	authorRepository domain.AuthorRepository
}

func (a *authorService) Get(id uuid.UUID) (domain.AuthorEntity, error) {
	return a.authorRepository.Get(id)
}

func (a *authorService) Create(author domain.AuthorEntity) (domain.AuthorEntity, error) {
	exists, err := a.authorRepository.ExistsByName(author.Name)
	if err != nil {
		return domain.AuthorEntity{}, err
	} else if exists {
		return domain.AuthorEntity{}, errors.NewAuthorWithNameAlreadyExists(author.Name)
	}
	author.Id = uuid.New()
	author.CreatedDate = time.Now()

	return a.authorRepository.Create(author)
}

func (a *authorService) Delete(id uuid.UUID) error {
	exists, err := a.authorRepository.Exists(id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NewAuthorNotFound(id)
	}
	return a.authorRepository.Delete(id)
}

func (a *authorService) GetAll() ([]domain.AuthorEntity, error) {
	return a.authorRepository.GetAll()
}

func NewAuthorService(authorRepository domain.AuthorRepository) domain.AuthorService {
	return &authorService{
		authorRepository: authorRepository,
	}
}
