package handler

import (
	"eta/model"
	"eta/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ValidateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}
func (h *Handler) Me(c echo.Context) error {
	id := userIDFromToken(c)
	u, err := h.userRepo.GetByCode(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, u)
}
func (h *Handler) Login(c echo.Context) error {
	req := new(model.UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	fmt.Println("req")
	fmt.Println(req)
	r := new(model.UserResponse)
	u, err := h.userRepo.GetByCode(&req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, "incorrect_uname")
	}
	fmt.Println(u.EmpPassword)
	fmt.Println(req.Password)
	if u.EmpPassword != req.Password {
		return c.JSON(http.StatusForbidden, "wrong_password")
	}
	r.User = *u
	r.Token = utils.GenerateJWT(uint(u.EmpCode))
	return c.JSON(http.StatusOK, r)
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
