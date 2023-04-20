package repo

import (
	"second/models"
)

// UserInterface is an interface for user repository.
type UserInterface interface {
	GetUser(id int64) *models.User
	GetUsers() []*models.User
	SaveUser(user models.User) bool
	UpdateUser(user models.User) bool
	DeleteUser(id int64) bool
}
