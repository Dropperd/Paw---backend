package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"websiteapi/entity"
	"websiteapi/service"
)

func GetFeedbacksFromUser(c *gin.Context) {
	imageId, ierr := strconv.ParseUint(c.Param("image_id"), 10, 64)
	userId, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if ierr != nil {
		c.JSON(400, gin.H{
			"message": "image id error",
		})
		return
	}
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "select feedbacks from user",
		"feedbacks": service.GetFeedbacks(imageId, userId),
	})
}

func GetImageFeedback(c *gin.Context) {
	imageId, ierr := strconv.ParseUint(c.Param("id"), 10, 64)
	if ierr != nil {
		c.JSON(400, gin.H{
			"message": "image id error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "get image feedback",
		"feedbacks": service.GetFeedbacks(imageId, 0),
	})
}

func UpdateFeedback(c *gin.Context) {
	var feedback entity.Feedback
	err := c.BindJSON(&feedback)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message":  "update feedback",
		"feedback": service.UpdateFeedback(feedback),
	})
}

func CreateFeedback(c *gin.Context) {
	clinicalid, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}
	var feedback entity.Feedback
	err := c.BindJSON(&feedback)
	if err != nil {
		return
	}

	feedback.IdClinical = clinicalid

	c.JSON(200, gin.H{
		"message":  "create feedback",
		"feedback": service.CreateFeedback(feedback),
	})
}

func DeleteFeedback(c *gin.Context) {
	feedbackId, err := strconv.ParseUint(c.Param("feedback_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "feedback id error",
		})
		return
	}

	service.DeleteFeedback(feedbackId)

	c.JSON(200, gin.H{
		"message": "delete feedback",
	})
}

func GetFeedbacksCount(c *gin.Context) {
	imageId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "image id error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "get feedbacks count",
		"count":   service.GetFeedbacksCount(imageId),
	})
}
