package handler

import (
	"github.com/gibiw/UrlShortener/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	links := router.Group("/links")
	{
		links.GET("/:guid", h.getLink)
		links.POST("/", h.createLink)
	}

	return router
}
