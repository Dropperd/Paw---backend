package main

import (
	"websiteapi/config"
	"websiteapi/controller"
	"websiteapi/cors"
	"websiteapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	defer config.CloseDB()

	router := gin.Default()
	router.Use(cors.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{

		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
		}

		user := v1.Group("/user")
		{
			user.GET("/", middleware.Authorized(), controller.GetAlluser)
			user.POST("/", controller.Register)

			user.GET("/:id", middleware.Authorized(), controller.GetUserProfile)
			user.GET("/token", middleware.Authorized(), controller.GetUserProfileFromToken)
			user.PUT("/:id", middleware.Authorized(), controller.UpdateProfile)
			user.DELETE("/:id", middleware.Authorized(), controller.DeleteAccount)

			user.GET("/type", middleware.Authorized(), controller.UserType)

			user.GET("/clinicals", middleware.Authorized(), controller.GetSharedClinicals)

			user.POST("/clinicals", middleware.Authorized(), controller.InsertImageClinical)
		}

		images := v1.Group("/image")
		{
			images.GET("/", middleware.Authorized(), controller.GetMyImages)
			images.POST("/", middleware.Authorized(), controller.InsertImage)
			images.GET("/:id", middleware.Authorized(), controller.GetImageById)
			images.PUT("/:id", middleware.Authorized(), controller.UpdateImage)
			images.DELETE("/:id", middleware.Authorized(), controller.DeleteImage)

			images.GET("/:id/count", middleware.Authorized(), controller.GetFeedbacksCount)
			images.GET("/:id/feedback", middleware.Authorized(), controller.GetImageFeedback)
			images.GET("/:id/filter", middleware.Authorized(), controller.GetUsersForFilter)
			//images.GET("/:id/clinicals", middleware.Authorized(), controller.GetSharedClinicals)

			//images.POST("/:id/clinicals", middleware.Authorized(), controller.InsertImageClinical)
		}

		bodypos := v1.Group("/bodyposition")
		{
			bodypos.GET("/", middleware.Authorized(), controller.GetAllBodyPosition)
		}
		clinical := v1.Group("/clinical")
		{
			clinical.GET("/", middleware.Authorized(), controller.GetAllClinical)
			clinical.GET("/user", middleware.Authorized(), controller.GetUserImagesByClinicalId)
		}

		feedback := v1.Group("/feedback")
		{
			feedback.GET("/:image_id/", middleware.Authorized(), controller.GetFeedbacksFromUser)
			feedback.POST("/", middleware.Authorized(), controller.CreateFeedback)
			feedback.PUT("/:id", middleware.Authorized(), controller.UpdateFeedback)
			feedback.DELETE("/:feedback_id", middleware.Authorized(), controller.DeleteFeedback)
		}

	}

	//user

	err := router.Run(":8080")
	if err != nil {
		return
	}
	// listen and serve on
}
