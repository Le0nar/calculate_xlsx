package handler

import (
	"github.com/gin-gonic/gin"
)

type service interface {
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()

	return r
}
