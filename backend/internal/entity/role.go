package entity

type Role struct {
	Id    uint8  `json:"id" db:"id"`
	Order uint8  `json:"order" db:"order"`
	Alias string `json:"alias" db:"alias"`
}
