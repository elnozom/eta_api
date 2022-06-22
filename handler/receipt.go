package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ReceiptsListByPosted(c echo.Context) error {
	// resp, err := utils.SubmitInvoice()
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusOK, err.Error())
	// }
	// // postedFilter, _ := strconv.ParseBool(c.QueryParam("posted"))
	// // result, err := h.receiptRepo.ListReceiptsByPosted(&postedFilter)
	// // utils.CheckErr(&err)
	return c.JSON(http.StatusOK, "resp")
}

func (h *Handler) ReceiptPost(c echo.Context) error {
	serial, _ := strconv.ParseUint(c.Param("serial"), 0, 64)
	var receipt model.Receipt
	err := h.receiptRepo.FindReceiptData(&serial, &receipt)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}

	var item model.ItemData
	var items []model.ItemData
	items = append(items, item)
	receipt.ItemData = items

	// receipt.DocumentTypeVersion = "1.0"
	// receipt.DocumentType = "I"

	return c.JSON(http.StatusOK, receipt)
}
