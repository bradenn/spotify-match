package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spotify-match/schemas"
	"spotify-match/service"
)

type Compare struct{}

func (e Compare) CompareController(c *gin.Context) {
	var json schemas.CompareInput
	if c.BindJSON(&json) == nil {
		obj, err := service.Compare(json)
		if err != nil {
			c.JSON(http.StatusBadRequest, http.ErrBodyNotAllowed)
		} else {
			c.JSON(200, obj)
		}
	} else {
		c.JSON(http.StatusBadRequest, http.ErrBodyNotAllowed)
	}
}
