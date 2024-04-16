package usercreate

import (
	"net/http"

	"github.com/boilerplate_cleancode/internal/log"
	"github.com/boilerplate_cleancode/internal/user"
	"github.com/gin-gonic/gin"
)

type HandlerHttp struct {
	logger  log.Logger
	service Service
}

func NewHandlerHttp(logger log.Logger, service Service) *HandlerHttp {
	return &HandlerHttp{
		logger:  logger,
		service: service,
	}
}

func (ref *HandlerHttp) Create(c *gin.Context) {
	var payload user.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	response, err := ref.service.Create(c, payload)
	ref.logger.Info(c, "Response", "user", response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, response)
}
