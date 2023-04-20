package repo

import (
	"github.com/jmoiron/sqlx"
	"log"
	"second/db"
	"second/models"
)

type UserPgRepo struct {
	db *sqlx.DB
}

func NewUserPgRepo() *UserPgRepo {
	dbConn := db.InitPg()
	return &UserPgRepo{db: dbConn}
}

func (repo *UserPgRepo) GetUser(id int64) *models.User {
	user := models.User{}
	err := repo.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil
		//log.Fatalf("Error while getting user: %v", err)
	}
	return &user
}

func (repo *UserPgRepo) GetUsers() []*models.User {
	users := []*models.User{}
	err := repo.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatalf("Error while getting users: %v", err)
	}
	return users
}

func (repo *UserPgRepo) SaveUser(user models.User) bool {
	_, err := repo.db.NamedExec("INSERT INTO users (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", user)
	if err != nil {
		//log.Fatalf("Error while saving user: %v", err)
		return false
	}
	return true
}

func (repo *UserPgRepo) UpdateUser(user models.User) bool {
	_, err := repo.db.NamedExec("UPDATE users SET first_name=:first_name, last_name=:last_name, email=:email WHERE id=:id", user)
	if err != nil {
		//log.Fatalf("Error while updating user: %v", err)
		return false
	}
	return true
}

// DeleteUser is a method for deleting user.
func (repo *UserPgRepo) DeleteUser(id int64) bool {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatalf("Error while deleting user: %v", err)
		//return false
	}
	return true
}
