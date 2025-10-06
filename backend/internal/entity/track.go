package entity

import (
	"time"
)

type TrackFile struct {
	ID          uint64    `json:"id" db:"id"`
	Bucket      string    `json:"bucket" db:"bucket"`
	ObjectKey   string    `json:"object_key" db:"object_key"`
	FileName    string    `json:"file_name" db:"file_name"`
	ContentType string    `json:"content_type" db:"content_type"`
	Size        int64     `json:"size" db:"size"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
	UploaderId  uint64    `json:"uploader_id" db:"uploader_id"`
	Duration    float64   `json:"duration" db:"duration"`
	Bitrate     int       `json:"bitrate" db:"bitrate"`
	TrackName   string    `json:"track_name" db:"track_name"`
}

type UploadRequest struct {
	FileData     []byte
	FileName     string
	ContentType  string
	UserID       uint64
	OriginalName string
}

type DownloadResponse struct {
	FileData    []byte
	FileName    string
	ContentType string
	Size        int64
}
