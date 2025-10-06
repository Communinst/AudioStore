package handler

import (
	"AudioShare/backend/internal/entity"
	"AudioShare/backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TrackHandler реализует обработку треков: загрузка, скачивание, получение метаинформации

// TrackHandler содержит зависимость на сервис треков
type TrackHandler struct {
	service service.TrackServiceInterface
}

func NewTrackHandler(srv service.TrackServiceInterface) *TrackHandler {
	return &TrackHandler{service: srv}
}

// UploadTrack — загрузка трека пользователем
func (h *TrackHandler) UploadTrack(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	userID, ok := userIDVal.(uint64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
		return
	}
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
		return
	}
	defer openedFile.Close()

	fileData := make([]byte, file.Size)
	_, err = openedFile.Read(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot read file"})
		return
	}

	uploadReq := &entity.UploadRequest{
		FileData:     fileData,
		FileName:     file.Filename,
		ContentType:  file.Header.Get("Content-Type"),
		UserID:       userID,
		OriginalName: file.Filename,
	}

	track, err := h.service.UploadTrack(c.Request.Context(), uploadReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, track)
}

// DownloadTrack — скачивание трека по ссылке
func (h *TrackHandler) DownloadTrack(c *gin.Context) {
	bucket := c.Query("bucket")
	objectKey := c.Query("objectKey")
	if bucket == "" || objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bucket or objectKey"})
		return
	}

	resp, err := h.service.DownloadTrack(c.Request.Context(), bucket, objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+resp.FileName)
	c.Data(http.StatusOK, resp.ContentType, resp.FileData)
}

// GetTrackInfo — получение метаинформации о треке
func (h *TrackHandler) GetTrackInfo(c *gin.Context) {
	bucket := c.Query("bucket")
	objectKey := c.Query("objectKey")
	if bucket == "" || objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing bucket or objectKey"})
		return
	}

	trackInfo, err := h.service.GetTrackInfo(c.Request.Context(), bucket, objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trackInfo)
}
