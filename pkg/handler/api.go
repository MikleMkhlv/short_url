package handler

import (
	"api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	var input api.Short
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.service.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"short_Code": id,
	})

}

type GetLongURL struct {
	ShortCode string `json:"short_url" binding:"required" db:"short_url" uri:"short_url"'`
}

func (h *Handler) geturl(c *gin.Context) {
	var input GetLongURL

	err := c.BindUri(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	fmt.Println(input)
	it, err := h.service.GetUrl(input.ShortCode)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Long URL": it,
	})

}
