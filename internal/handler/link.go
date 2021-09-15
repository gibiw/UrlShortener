package handler

import (
	"fmt"
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
		"url": fmt.Sprintf("http://%s:%s/%s", h.host, h.port, mod),
	})
}

func (h *Handler) getLink(c *gin.Context) {
	guid := c.Param("hash")

	link, err := h.services.GetByHash(guid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, link)
}
