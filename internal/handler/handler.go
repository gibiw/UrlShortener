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

	redirect := router.Group("/")
	{
		redirect.GET("/:hash", h.getLink)
	}

	links := router.Group("/links")
	{
		links.POST("/", h.createLink)
	}

	return router
}
