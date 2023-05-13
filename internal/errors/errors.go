package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type authorNotFound struct {
	authorId uuid.UUID
}

func (a *authorNotFound) Error() string {
	return fmt.Sprintf("Author [%s] not found", a.authorId)
}

func NewAuthorNotFound(authorId uuid.UUID) error {
	return &authorNotFound{authorId: authorId}
}

func IsErrorAuthorNotFound(err error) bool {
	if _, ok := err.(*authorNotFound); ok {
		return true
	}
	return false
}

type authorAlreadyExists struct {
	authorId uuid.UUID
}

func (a *authorAlreadyExists) Error() string {
	return fmt.Sprintf("Author [%s] already exists", a.authorId)
}

func NewAuthorAlreadyExists(authorId uuid.UUID) error {
	return &authorAlreadyExists{authorId: authorId}
}

func IsErrorAuthorAlreadyExists(err error) bool {
	if _, ok := err.(*authorAlreadyExists); ok {
		return true
	}
	return false
}

type authorWithNameAlreadyExists struct {
	authorName string
}

func (a *authorWithNameAlreadyExists) Error() string {
	return fmt.Sprintf("Author [%s] already exists", a.authorName)
}

func NewAuthorWithNameAlreadyExists(authorName string) error {
	return &authorWithNameAlreadyExists{authorName: authorName}
}

func IsErrorAuthorWithNameAlreadyExists(err error) bool {
	if _, ok := err.(*authorWithNameAlreadyExists); ok {
		return true
	}
	return false
}
