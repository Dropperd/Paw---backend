package controller

import (
	"github.com/gin-gonic/gin"
	"websiteapi/entity/dto"
	"websiteapi/service"
)

func Login(c *gin.Context) {
	loginDTO := dto.LoginDTO{}

	err := c.ShouldBind(&loginDTO)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error binding login",
			"error":   err.Error(),
		})
		return
	}

	token, err := service.Login(loginDTO)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error invalid login",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"token":   token,
	})

}
