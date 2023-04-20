package repo

import (
	"second/models"
)

type UserMockRepo struct {
	index int64
	users []*models.User
}

func (u *UserMockRepo) GetUser(id int64) *models.User {
	if id > 0 && u.index >= id {
		return u.users[id-1]
	}
	return nil
}

func (u *UserMockRepo) GetUsers() []*models.User {
	return u.users
}

func (u *UserMockRepo) SaveUser(user models.User) bool {
	u.index++
	user.ID = u.index
	u.users = append(u.users, &user)
	return true
}

func (u *UserMockRepo) UpdateUser(user models.User) bool {
	u.users[user.ID-1] = &user
	return true
}

func (u *UserMockRepo) DeleteUser(id int64) bool {
	for i, user := range u.users {
		if user.ID == id {
			u.users[i] = nil
		}
	}
	return true
}

func NewUserMockRepo() *UserMockRepo {
	return &UserMockRepo{0, []*models.User{}}
}
