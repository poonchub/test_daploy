package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poonchub/sa-67-example/config"
	"github.com/poonchub/sa-67-example/entity"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender

	db := config.DB()

	db.Find(&genders)

	c.JSON(http.StatusOK, &genders)
}
