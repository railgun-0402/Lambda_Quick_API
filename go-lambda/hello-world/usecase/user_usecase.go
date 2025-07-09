package usecase

import (
	"hello-world/repository"

	"github.com/labstack/echo/v4"
)

type UserUseCase interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type useUserCase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUseCase {
	return &useUserCase{repo: repo}
}

func (u *useUserCase) Create(c echo.Context) error {
	return u.repo.CreateUser(c)
}

func (u *useUserCase) Get(c echo.Context) error {
	return u.repo.GetUser(c)
}

func (u *useUserCase) Update(c echo.Context) error {
	return u.repo.UpdateUser(c)
}

func (u *useUserCase) Delete(c echo.Context) error {
	return u.repo.DeleteUser(c)
}
