package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"clean-architecture/internal/domain"
	"clean-architecture/internal/errors"
)

type CreateAuthorRequest struct {
	Name string `json:"name,omitempty"`
}

type CreateAuthorResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AuthorResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type authorHandler struct {
	authorService domain.AuthorService
}

func NewAuthorHandler(authorService domain.AuthorService) domain.AuthorHandler {
	return &authorHandler{
		authorService: authorService,
	}
}

func (a *authorHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createRequest CreateAuthorRequest

		err := c.BindJSON(&createRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
				Message:        "Failed to read request body",
				MessageDetails: err.Error(),
			})
			return
		}

		newAuthor, err := a.authorService.Create(domain.AuthorEntity{
			Name: createRequest.Name,
		})
		if err != nil {
			c.AbortWithStatusJSON(a.findResponseCode(err), ErrorResponse{
				Message:        "Failed to create author",
				MessageDetails: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusCreated, CreateAuthorResponse{
				Id:   newAuthor.Id,
				Name: newAuthor.Name,
			})
		}
	}
}

func (a *authorHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawId := c.Param("id")
		if rawId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Message:        "Failed to delete author",
				MessageDetails: "Required path param 'id' not found",
			})
			return
		}
		id, err := uuid.Parse(rawId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Message:        "Failed to delete author",
				MessageDetails: "Invalid path param 'id'",
			})
			return
		}
		err = a.authorService.Delete(id)
		if err != nil {
			c.AbortWithStatusJSON(a.findResponseCode(err), ErrorResponse{
				Message:        "Failed to delete author",
				MessageDetails: err.Error(),
			})
			return
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}

func (a *authorHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawId := c.Param("id")
		if rawId == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Message:        "Failed to get author",
				MessageDetails: "Required path param 'id' not found",
			})
			return
		}
		id, err := uuid.Parse(rawId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Message:        "Failed to get author",
				MessageDetails: "Invalid path param 'id'",
			})
			return
		}
		author, err := a.authorService.Get(id)
		if err != nil {
			c.AbortWithStatusJSON(a.findResponseCode(err), ErrorResponse{
				Message:        "Failed to get author",
				MessageDetails: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, AuthorResponse{
				Id:   author.Id,
				Name: author.Name,
			})
		}
	}
}

func (a *authorHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		authors, err := a.authorService.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(a.findResponseCode(err), ErrorResponse{
				Message:        "Failed to get authors",
				MessageDetails: err.Error(),
			})
			return
		} else {
			var response []AuthorResponse
			for _, author := range authors {
				response = append(response, AuthorResponse{
					Id:   author.Id,
					Name: author.Name,
				})
			}
			c.JSON(http.StatusOK, response)
		}
	}
}

func (a *authorHandler) findResponseCode(err error) int {
	switch {
	case errors.IsErrorAuthorNotFound(err):
		return http.StatusNotFound
	case errors.IsErrorAuthorAlreadyExists(err) || errors.IsErrorAuthorWithNameAlreadyExists(err):
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
