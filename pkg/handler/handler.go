package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yaroshevichM/software-school-case/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/subscribe", h.createSubscription)
	router.GET("/rate", h.getRate)

	return router
}
