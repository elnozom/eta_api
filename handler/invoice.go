package handler

import (
	"eta/model"
	"eta/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	// var log2 model.Log
	// var log3 model.Log
	// var log4 model.Log
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

	// resp, err := utils.SignInvoices(&invoices)
	// if utils.CheckErr(&err) {
	// 	log2 = model.Log{
	// 		InternalID:   "",
	// 		SubmissionID: "",
	// 		StoreCode:    req.Store,
	// 		Serials:      req.Serilas,
	// 		LogText:      "failed to receive requests from dotnet to golang application",
	// 		ErrText:      err.Error(),
	// 		Posted:       false,
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// log2 = model.Log{
	// 	InternalID:   "",
	// 	SubmissionID: "",
	// 	StoreCode:    req.Store,
	// 	Serials:      req.Serilas,
	// 	LogText:      *resp,
	// 	ErrText:      "",
	// 	Posted:       false,
	// }
	// _, err = h.logRepo.ELogInsert(&log2)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// var finalResponse model.InvoiceSubmitResp
	// err = json.Unmarshal([]byte(*resp), &finalResponse)
	// if utils.CheckErr(&err) {
	// 	log3 = model.Log{
	// 		InternalID:   "",
	// 		SubmissionID: "",
	// 		StoreCode:    req.Store,
	// 		Serials:      req.Serilas,
	// 		LogText:      "error while parsing the response coming from the eta api '" + *resp + "'",
	// 		ErrText:      err.Error(),
	// 		Posted:       false,
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// log3 = model.Log{
	// 	InternalID:   "",
	// 	SubmissionID: finalResponse.SubmissionId,
	// 	StoreCode:    req.Store,
	// 	Serials:      req.Serilas,
	// 	LogText:      "eta api response parsed successfully \n'" + *resp + "'",
	// 	ErrText:      "",
	// 	Posted:       true,
	// }
	// _, err = h.logRepo.ELogInsert(&log3)
	// for i := 0; i < len(finalResponse.AcceptedDocuments); i++ {
	// 	var invoice = finalResponse.AcceptedDocuments[i]
	// 	internalID := strings.Split(invoice.InternalId, "-")
	// 	storeString := internalID[0]
	// 	serialString := internalID[1]
	// 	store, _ := strconv.Atoi(storeString)
	// 	serial, _ := strconv.Atoi(serialString)
	// 	_, err = h.invoiceRepo.EInvoicePost(&serial, &store, &invoice.UUID)
	// 	if utils.CheckErr(&err) {
	// 		return c.JSON(http.StatusInternalServerError, err.Error())
	// 	}

	// }
	return c.JSON(http.StatusInternalServerError, invoices)

	// return c.JSON(http.StatusOK, *invoices)
}

func (h *Handler) InvoicePosted(c echo.Context) error {
	req := new(model.InvoiceSubmitResp)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// resp, err := utils.SignInvoices(&invoices)
	// if utils.CheckErr(&err) {
	// 	log2 = model.Log{
	// 		InternalID:   "",
	// 		SubmissionID: "",
	// 		StoreCode:    req.Store,
	// 		Serials:      req.Serilas,
	// 		LogText:      "failed to receive requests from dotnet to golang application",
	// 		ErrText:      err.Error(),
	// 		Posted:       false,
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// log2 = model.Log{
	// 	InternalID:   "",
	// 	SubmissionID: "",
	// 	StoreCode:    req.Store,
	// 	Serials:      req.Serilas,
	// 	LogText:      *resp,
	// 	ErrText:      "",
	// 	Posted:       false,
	// }
	// _, err = h.logRepo.ELogInsert(&log2)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// var finalResponse model.InvoiceSubmitResp
	// err = json.Unmarshal([]byte(*resp), &finalResponse)
	// if utils.CheckErr(&err) {
	// 	log3 = model.Log{
	// 		InternalID:   "",
	// 		SubmissionID: "",
	// 		StoreCode:    req.Store,
	// 		Serials:      req.Serilas,
	// 		LogText:      "error while parsing the response coming from the eta api '" + *resp + "'",
	// 		ErrText:      err.Error(),
	// 		Posted:       false,
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// log3 = model.Log{
	// 	InternalID:   "",
	// 	SubmissionID: finalResponse.SubmissionId,
	// 	StoreCode:    req.Store,
	// 	Serials:      req.Serilas,
	// 	LogText:      "eta api response parsed successfully \n'" + *resp + "'",
	// 	ErrText:      "",
	// 	Posted:       true,
	// }
	// _, err = h.logRepo.ELogInsert(&log3)

	fmt.Println(req)
	fmt.Println("req")
	for i := 0; i < len(req.AcceptedDocuments); i++ {
		var invoice = req.AcceptedDocuments[i]
		internalID := strings.Split(invoice.InternalId, "-")
		storeString := internalID[0]
		serialString := internalID[1]
		store, _ := strconv.Atoi(storeString)
		serial, _ := strconv.Atoi(serialString)
		_, err := h.invoiceRepo.EInvoicePost(&serial, &store, &invoice.UUID)
		if utils.CheckErr(&err) {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}
	return c.JSON(http.StatusOK, "invoices")

	// return c.JSON(http.StatusOK, *invoices)
}
