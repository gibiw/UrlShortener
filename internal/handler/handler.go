package handler

import (
	"github.com/gibiw/UrlShortener/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	host     string
	port     string
}

func NewHandler(services *service.Service, host, port string) *Handler {
	return &Handler{services: services, host: host, port: port}
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
