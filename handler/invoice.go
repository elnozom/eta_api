package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InvoicesList(c echo.Context) error {
	req := new(model.ListInvoicessRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	result, err := h.invoiceRepo.ListEInvoices(req)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) InvoicePost(c echo.Context) error {
	req := new(model.PostInvoicessRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	var log model.Log
	var log2 model.Log
	invoices, err := h.invoiceRepo.FindInvoiceData(req, h.companyInfo)
	if utils.CheckErr(&err) {
		log = model.Log{
			InternalID:   "",
			SubmissionID: "",
			StoreCode:    req.Store,
			Serials:      req.Serilas,
			LogText:      "failed to create request on the golang application",
			ErrText:      err.Error(),
			Posted:       false,
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log = model.Log{
		InternalID:   "",
		SubmissionID: "",
		StoreCode:    req.Store,
		Serials:      req.Serilas,
		LogText:      "request generated on the golang application successfully",
		ErrText:      "",
		Posted:       true,
	}
	_, err = h.logRepo.ELogInsert(&log)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp, err := utils.SignInvoices(&invoices)
	if utils.CheckErr(&err) {
		log2 = model.Log{
			InternalID:   "",
			SubmissionID: "",
			StoreCode:    req.Store,
			Serials:      req.Serilas,
			LogText:      "failed to receive requests from dotnet to golang application",
			ErrText:      err.Error(),
			Posted:       false,
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log2 = model.Log{
		InternalID:   "",
		SubmissionID: "",
		StoreCode:    req.Store,
		Serials:      req.Serilas,
		LogText:      *resp,
		ErrText:      "",
		Posted:       false,
	}
	_, err = h.logRepo.ELogInsert(&log2)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// for i := 0; i < len(invoices); i++ {
	// 	var invoice = invoices[i]
	// 	_, err = h.invoiceRepo.EInvoicePost(&invoice.Serial, &invoice.Issuer.Address.BranchId)
	// }
	return c.JSON(http.StatusOK, *resp)
}
