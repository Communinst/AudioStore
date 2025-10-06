package repository

import "AudioShare/backend/internal/entity"

type Entity interface {
	entity.User |
		entity.Role |
		entity.UserCache |
		entity.TrackFile |
		entity.Dump
}
