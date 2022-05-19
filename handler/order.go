package handler

import (
	"eta/model"
	"eta/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) OrdersListByTransSerialAndConverted(c echo.Context) error {
	req := new(model.ListOrdersRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	fmt.Println("req")
	fmt.Println(req)
	result, err := h.orderRepo.ListByTransSerialAndConverted(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) OrdersConvertToEta(c echo.Context) error {
	serial, _ := strconv.ParseInt(c.Param("serial"), 0, 64)
	result, err := h.orderRepo.ConvertToEta(&serial)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
