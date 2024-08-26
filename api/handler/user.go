package handler

import (
	"bug_busters/pkg/models"
	"context"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/gin-gonic/gin"
)

const (
	MinioEndpoint  = "3.125.33.48:9000"
	MinioAccessKey = "minioadmin"
	MinioSecretKey = "minioadmin"
	BucketName     = "car" // Убедитесь, что бакет с таким именем существует
)

// GetProfile godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.UserProfile
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	id := c.MustGet("user_id").(string)
	profile, err := h.user.GetProfile(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get user profile", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetProfileAdmin godoc
// @Summary Get User Profile
// @Description Retrieve the profile of a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.UserProfile
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/profile/{id} [get]
func (h *Handler) GetProfileAdmin(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.user.GetProfile(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get user profile", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// AddImage godoc
// @Summary Upload an image and update car information
// @Description Uploads an image file to MinIO and updates car image information with the file URL.
// @Tags Images
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "Image file to upload"
// @Param id path string true "User ID to associate with the uploaded image"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error "Bad request, e.g., missing file or user ID"
// @Failure 500 {object} models.Error "Internal server error, e.g., failure in MinIO or external service"
// @Router /upload [post]
func (h *Handler) AddImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		h.log.Error("Failed to get file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectName := "cars/" + file.Filename

	// Инициализация MinIO клиента
	minioClient, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false, // Установите true, если используете HTTPS
	})
	if err != nil {
		h.log.Error("Failed to create MinIO client", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Открытие файла
	OpenFile, err := file.Open()
	if err != nil {
		h.log.Error("Failed to open file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer OpenFile.Close() // Закрытие файла после завершения работы

	// Определение типа содержимого файла
	contentType := mime.TypeByExtension(filepath.Ext(file.Filename))

	// Загрузка файла в MinIO
	info, err := minioClient.PutObject(context.Background(), BucketName, objectName, OpenFile, file.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		h.log.Error("Failed to upload file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Создание URL для доступа к файлу
	fileURL := fmt.Sprintf("http://%s/%s/%s", MinioEndpoint, BucketName, objectName)

	req := &models.UpdateCarImage{
		Url:    fileURL,
		UserId: c.PostForm("id"),
	}

	res, err := h.user.AddImage(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.log.Error("Failed to add image", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_url": fileURL, "etag": info.ETag, "message": res.Message})
}

// GetImage godoc
// @Summary Get Car Image
// @Description Retrieve the image of a user's car
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Url
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /upload [get]
func (h *Handler) GetImage(c *gin.Context) {
	id := c.Param("id")
	image, err := h.user.GetImage(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get image", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, image)
}

// GetPaidFinesU godoc
// @Summary Get Paid Fines
// @Description Retrieve all paid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/paid_fines [get]
func (h *Handler) GetPaidFinesU(c *gin.Context) {
	id := c.MustGet("user_id").(string)
	fines, err := h.user.GetPaidFinesU(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get paid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetPaidFinesAdmin godoc
// @Summary Get Paid Fines
// @Description Retrieve all paid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/paid_fines/{id} [get]
func (h *Handler) GetPaidFinesAdmin(c *gin.Context) {
	id := c.Param("id")
	fines, err := h.user.GetPaidFinesU(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get paid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetUnpaid godoc
// @Summary Get Unpaid Fines
// @Description Retrieve all unpaid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /user/unpaid_fines [get]
func (h *Handler) GetUnpaid(c *gin.Context) {
	id := c.MustGet("user_id").(string)
	fines, err := h.user.GetUnpaid(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get unpaid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// GetUnpaidAdmin godoc
// @Summary Get Unpaid Fines
// @Description Retrieve all unpaid fines for a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.UserFines
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/unpaid_fines/{id} [get]
func (h *Handler) GetUnpaidAdmin(c *gin.Context) {
	id := c.Param("id")
	fines, err := h.user.GetUnpaid(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to get unpaid fines", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, fines)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /admin/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	msg, err := h.user.DeleteUser(models.UserId{Id: id})
	if err != nil {
		h.log.Error("Failed to delete user", err)
		c.JSON(http.StatusInternalServerError, models.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}
