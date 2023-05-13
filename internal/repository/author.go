package repository

import (
	"github.com/google/uuid"

	"clean-architecture/internal/domain"
	"clean-architecture/internal/errors"
)

type authorRepository struct {
	data       map[uuid.UUID]domain.AuthorEntity
	dataByName map[string]domain.AuthorEntity
}

func (a *authorRepository) Get(id uuid.UUID) (domain.AuthorEntity, error) {
	if entity, ok := a.data[id]; !ok {
		return domain.AuthorEntity{}, errors.NewAuthorNotFound(id)
	} else {
		return entity, nil
	}
}

func (a *authorRepository) Create(author domain.AuthorEntity) (domain.AuthorEntity, error) {
	if _, ok := a.data[author.Id]; ok {
		return domain.AuthorEntity{}, errors.NewAuthorAlreadyExists(author.Id)
	}
	a.dataByName[author.Name] = author
	a.data[author.Id] = author
	return author, nil
}

func (a *authorRepository) Delete(id uuid.UUID) error {
	if _, ok := a.data[id]; !ok {
		return errors.NewAuthorNotFound(id)
	}

	delete(a.data, id)
	return nil
}

func (a *authorRepository) Exists(id uuid.UUID) (bool, error) {
	if _, ok := a.data[id]; ok {
		return true, nil
	} else {
		return false, nil
	}
}

func (a *authorRepository) ExistsByName(name string) (bool, error) {
	if _, ok := a.dataByName[name]; ok {
		return true, nil
	} else {
		return false, nil
	}
}

func (a *authorRepository) GetAll() ([]domain.AuthorEntity, error) {
	authors := make([]domain.AuthorEntity, 0, len(a.data))
	for _, author := range a.data {
		authors = append(authors, author)
	}
	return authors, nil
}

func NewAuthorRepository() domain.AuthorRepository {
	return &authorRepository{
		data:       map[uuid.UUID]domain.AuthorEntity{},
		dataByName: map[string]domain.AuthorEntity{},
	}
}
