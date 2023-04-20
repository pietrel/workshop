package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"second/models"
	"second/repo"
	"strconv"
)

type Handlers struct {
	userRepo repo.UserInterface
	//client http , interface
}

func NewHandlers(userRepo repo.UserInterface) *Handlers {
	return &Handlers{userRepo: userRepo}
}

func (handlers *Handlers) DeleteUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	handlers.userRepo.DeleteUser(id)
	return c.JSON(http.StatusOK, "User deleted")
}

func (handlers *Handlers) UpdateUser(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	user.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	handlers.userRepo.UpdateUser(*user)
	if err != nil {
		log.Fatalf("Error while updating user: %v", err)
	}
	return c.JSON(http.StatusOK, user)
}

func (handlers *Handlers) SaveUser(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		log.Fatalf("Error while binding user: %v", err)
	}
	handlers.userRepo.SaveUser(*user)
	return c.JSON(http.StatusCreated, user)
}

func (handlers *Handlers) GetUser(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := handlers.userRepo.GetUser(id)
	if user == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, user)
}

func (handlers *Handlers) GetUsers(c echo.Context) error {
	//var users []models.User
	users := handlers.userRepo.GetUsers()
	return c.JSON(http.StatusOK, users)
}
