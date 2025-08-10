package entity

import "time"

type User struct {
	Id         uint64    `json:"id" db:"id"`
	Login      string    `json:"login" db:"login"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Nickname   string    `json:"nickname" db:"nickname"`
	Registered time.Time `json:"registered" db:"registered"`
	RoleId     uint8     `json:"role_id" db:"role_id"`
}

func (u *User) Update(fresh *User) {
	if fresh.Email != "" {
		u.Email = fresh.Email
	}
	if fresh.Password != "" {
		u.Password = fresh.Password
	}
	if fresh.Nickname != "" {
		u.Nickname = fresh.Nickname
	}
	if fresh.RoleId > 0 {
		u.RoleId = fresh.RoleId
	}
}

func DefaultUser() *User {
	return &User{
		Id:         0,
		Login:      "",
		Email:      "",
		Password:   "",
		Nickname:   "",
		Registered: time.Now(),
		RoleId:     0,
	}
}
