package handler

import (
	"net/http"

	link "github.com/gibiw/UrlShortener"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createLink(c *gin.Context) {
	var input link.LinkItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	mod, err := h.services.LinkItem.Create(input.Original)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": mod,
	})
}

func (h *Handler) getLink(c *gin.Context) {
	guid := c.Param("guid")

	link, err := h.services.GetByUrl(guid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": link,
	})
}
