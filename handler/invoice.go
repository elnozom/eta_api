package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InvoicesListByPosted(c echo.Context) error {
	postedFilter, _ := strconv.ParseBool(c.QueryParam("posted"))
	result, err := h.invoiceRepo.ListEInvoicesByPosted(&postedFilter)
	utils.CheckErr(&err)
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) InvoicePost(c echo.Context) error {
	serial, _ := strconv.ParseUint(c.Param("serial"), 0, 64)
	var invoice model.Invoice
	err := h.invoiceRepo.FindInvoiceData(&serial, &invoice)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	invoice.DocumentTypeVersion = "1.0"
	invoice.DocumentType = "I"

	return c.JSON(http.StatusOK, invoice)
}
