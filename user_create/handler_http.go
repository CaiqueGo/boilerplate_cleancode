package usercreate

import (
	"net/http"

	"github.com/boilerplate_cleancode/internal/user"
	"github.com/gin-gonic/gin"
)

type HandlerHttp struct {
	service Service
}

func NewHandlerHttp(service Service) *HandlerHttp {
	return &HandlerHttp{
		service: service,
	}
}

func (ref *HandlerHttp) Create(c *gin.Context) {
	var payload user.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	response, err := ref.service.Create(c, payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, response)
}
