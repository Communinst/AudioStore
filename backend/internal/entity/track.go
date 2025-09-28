package entity

import (
	"time"
)

type TrackFile struct {
	ID           uint64    `json:"id" db:"id"`
	FileName     string    `json:"file_name" db:"file_name"`
	OriginalName string    `json:"original_name" db:"original_name"`
	BucketName   string    `json:"bucket_name" db:"bucket_name"`
	ObjectName   string    `json:"object_name" db:"object_name"`
	Size         int64     `json:"size" db:"size"`
	ContentType  string    `json:"content_type" db:"content_type"`
	ETag         string    `json:"etag" db:"etag"`
	UploadedAt   time.Time `json:"uploaded_at" db:"uploaded_at"`
	UserID       uint64    `json:"user_id" db:"user_id"`
	Duration     float64   `json:"duration" db:"duration"`
	Bitrate      int       `json:"bitrate" db:"bitrate"`
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
