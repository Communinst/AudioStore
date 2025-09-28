package entity

type Dump struct {
	ID       int    `db:"id" json:"id"`             
	Filename string `db:"filename" json:"filename"` 
	Size     int64  `db:"size" json:"size"`         
}
