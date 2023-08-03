package controller

import (
	"strconv"
	"websiteapi/entity"
	"websiteapi/service"

	"github.com/gin-gonic/gin"
)

func GetAllClinical(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "get all clinical",
		"clinical": service.GetAllClinical(),
	})
}

func GetSharedClinicals(c *gin.Context) {
	userID, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": uerr.Error(),
		})
		return
	}
	user, err := service.GetSharedClinicals(userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "get shared clinicals",
		"clinical": user,
	})
}

func InsertImageClinical(c *gin.Context) {
	userID, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": uerr.Error(),
		})
		return
	}
	var clinical entity.ImageClinical
	err := c.ShouldBind(&clinical)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	clinical.UserId = userID
	service.InsertImageClinical(clinical)
}

func GetAlluser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select all user",
		"user":    service.GetAlluser(),
	})
}

func Register(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	userRegister, err := service.InsertUser(user)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "add user",
		"user":    userRegister,
	})
}

func GetUserProfile(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	user, err := service.GetUserProfile(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "select user",
		"user":    user,
	})
}

func GetUserProfileFromToken(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	user, err := service.GetUserProfile(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "select user",
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	userId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if iderr != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	user.ID = userId
	userUpdate, err := service.UpdateUserByID(user, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update user by id",
		"user":    userUpdate,
	})
}

func DeleteAccount(c *gin.Context) {
	userId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if iderr != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}
	err := service.DeleteUserByID(userId, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "delete user by id",
	})
}

func UserType(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user type",
		"user":    service.GetUserType(),
	})
}
