package entity

type Dump struct {
	ID       int    `db:"id" json:"id"`
	Filename string `db:"filename" json:"filename"`
	Size     int64  `db:"size" json:"size"`
}

type MinioDump struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	FileCount   int    `db:"file_count" json:"file_count"`
	TotalSize   int64  `db:"total_size" json:"total_size"`
	Status      string `db:"status" json:"status"` // "completed", "failed", "in_progress"
}
