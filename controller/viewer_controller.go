package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"websiteapi/entity"
	"websiteapi/service"

	"github.com/gin-gonic/gin"
)

func GetUserImagesByClinicalId(c *gin.Context) {
	clinicalId, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}

	value := c.Query("start_date")
	var startDate time.Time
	if value != "" {
		if value != "Invalid date" {
			startDate, err = time.Parse(time.DateOnly, c.Query("start_date"))
		}
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}
	}

	value = c.Query("end_date")
	var endDate time.Time
	if value != "" {
		if value != "Invalid date" {
			endDate, err = time.Parse(time.DateOnly, c.Query("end_date"))
		}
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}
	}

	value = c.Query("body_pos")
	var bodyPos string
	if value != "" {
		if value != "none" {
			bodyPos = c.Query("body_pos")
		}
	}

	value = c.Query("user_id")
	var userId uint64
	if value != "" {
		userId, err = strconv.ParseUint(c.Query("user_id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "id error",
			})
			return
		}
	}

	images, err := service.GetUserImagesByClinicalId(clinicalId, startDate, endDate, bodyPos, userId)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	// Decrypt and convert the images
	decryptedImages := make([]entity.Image, len(images))
	for i, image := range images {
		encryptedImageBytes, err := base64.StdEncoding.DecodeString(image.Image)
		if err != nil {
			// Handle the decryption error
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		key := []byte("your-encryption1") // 32 bytes

		decryptedImageBytes, err := decryptImage(encryptedImageBytes, key)
		if err != nil {
			// Handle the decryption error
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		decryptedImage := entity.Image{
			ID:           image.ID,
			Image:        base64.StdEncoding.EncodeToString(decryptedImageBytes),
			Description:  image.Description,
			BodyPosition: image.BodyPosition,
			UserID:       image.UserID,
			Added_At:     image.Added_At,
			Updated_At:   image.Updated_At,
		}

		decryptedImages[i] = decryptedImage
	}

	c.JSON(200, gin.H{
		"message": "get user images by clinical id success",
		"images":  decryptedImages,
	})
}

func GetMyImages(c *gin.Context) {
	userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	value := c.Query("start_date")
	var startDate time.Time
	if value != "" {
		if value != "Invalid date" {
			startDate, err = time.Parse(time.DateOnly, c.Query("start_date"))
		}

		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}
	}

	value = c.Query("end_date")
	var endDate time.Time
	if value != "" {
		if value != "Invalid date" {
			endDate, err = time.Parse(time.DateOnly, c.Query("end_date"))
		}

		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}
	}

	value = c.Query("body_pos")
	var bodyPos string
	if value != "" {
		if value != "none" {
			bodyPos = c.Query("body_pos")
		}
	}

	images, err := service.GetMyImages(userId, startDate, endDate, bodyPos)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	// Decrypt and convert the images
	decryptedImages := make([]entity.Image, len(images))
	for i, image := range images {
		encryptedImageBytes, err := base64.StdEncoding.DecodeString(image.Image)
		if err != nil {
			// Handle the decryption error
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		key := []byte("your-encryption1") // 32 bytes

		decryptedImageBytes, err := decryptImage(encryptedImageBytes, key)
		if err != nil {
			// Handle the decryption error
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		decryptedImage := entity.Image{
			ID:           image.ID,
			Image:        base64.StdEncoding.EncodeToString(decryptedImageBytes),
			Description:  image.Description,
			BodyPosition: image.BodyPosition,
			UserID:       image.UserID,
			Added_At:     image.Added_At,
			Updated_At:   image.Updated_At,
		}

		decryptedImages[i] = decryptedImage
	}

	c.JSON(200, gin.H{
		"message": "get my images",
		"images":  decryptedImages,
	})
}

func GetImageById(c *gin.Context) {
	imageId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}

	userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}

	image, err := service.GetImageById(imageId)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	if image.UserID == userId || service.IsClinical(userId) == true {

		encryptedImageBytes, err := base64.StdEncoding.DecodeString(image.Image)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		key := []byte("your-encryption1") // 32 bytes

		decryptedImageBytes, err := decryptImage(encryptedImageBytes, key)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to decrypt image",
			})
			return
		}

		decryptedImage := entity.Image{
			ID:           image.ID,
			Image:        base64.StdEncoding.EncodeToString(decryptedImageBytes),
			Description:  image.Description,
			BodyPosition: image.BodyPosition,
			UserID:       image.UserID,
			Added_At:     image.Added_At,
			Updated_At:   image.Updated_At,
		}

		c.JSON(200, gin.H{
			"message": "get image by id",
			"image":   decryptedImage,
		})
	} else {
		c.JSON(403, gin.H{
			"message": "unauthorized",
		})
		return
	}
}

func InsertImage(c *gin.Context) {
	userId, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if uerr != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}

	// Retrieve the uploaded image file
	file, err := c.FormFile("imageOut")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to get uploaded image",
		})
		return
	}

	// Save the uploaded image file to a temporary location
	err = c.SaveUploadedFile(file, "temp.jpg")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to save uploaded image",
		})
		return
	}

	var image entity.Image
	err = c.ShouldBind(&image)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	image.UserID = userId

	// Read the temporary image file
	imageBytes, err := ioutil.ReadFile("temp.jpg")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to read uploaded image",
		})
		return
	}

	// Delete the temporary image file
	err = os.Remove("temp.jpg")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to delete temporary image file",
		})
		return
	}

	key := []byte("your-encryption1")

	encryptedImageBytes, err := encryptImage(imageBytes, key)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error encrypting image",
		})
		return
	}

	encryptedImageString := base64.StdEncoding.EncodeToString(encryptedImageBytes)

	image.Image = encryptedImageString

	// Bind the remaining form data fields
	err = c.ShouldBind(&image)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	imageInsert, err := service.InsertImage(image)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "add image",
		"image":   imageInsert,
	})
}

func UpdateImage(c *gin.Context) {
	imageId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, uerr := strconv.ParseUint(c.GetString("user_id"), 10, 64)
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

	var image entity.Image
	err := c.ShouldBind(&image)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}

	image.ID = imageId
	imageUpdate, err := service.UpdateImageById(image, userId)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update image",
		"image":   imageUpdate,
	})
}

func DeleteImage(c *gin.Context) {
	userId, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "user id error",
		})
		return
	}

	imageId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}

	err = service.DeleteImage(imageId, userId)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete image",
	})
}

func GetUsersForFilter(c *gin.Context) {

	clinicalId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	users := service.GetUsersForFilter(clinicalId)
	c.JSON(200, gin.H{
		"message": "get users for filter",
		"users":   users,
	})
}

func GetAllBodyPosition(c *gin.Context) {
	bodyPos, err := service.GetAllBodyPosition()

	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "get all body position",
		"bodyPos": bodyPos,
	})
}

func encryptImage(plainImageBytes []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	encryptedImageBytes := aesgcm.Seal(nil, nonce, plainImageBytes, nil)
	encryptedImageBytes = append(nonce, encryptedImageBytes...)

	return encryptedImageBytes, nil
}

func decryptImage(encryptedImageBytes []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := encryptedImageBytes[:12]
	encryptedImageBytes = encryptedImageBytes[12:]

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainImageBytes, err := aesgcm.Open(nil, nonce, encryptedImageBytes, nil)
	if err != nil {
		return nil, err
	}

	return plainImageBytes, nil
}
