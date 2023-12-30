package handler

import (
	"encoding/json"
	"eta/model"
	"eta/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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

// func (h *Handler) ReceiptPost(c echo.Context) error {
//  	var receipt model.Receipt
// 	err := h.receiptRepo.FindReceiptData(&serial, &receipt)
// 	if utils.CheckErr(&err) {
// 		return c.JSON(http.StatusOK, err.Error())
// 	}

// 	var item model.ItemData
// 	var items []model.ItemData
// 	items = append(items, item)
// 	receipt.ItemData = items

// 	// receipt.DocumentTypeVersion = "1.0"
// 	// receipt.DocumentType = "I"

//		return c.JSON(http.StatusOK, receipt)
//	}
func (h *Handler) GenerateUUID(c echo.Context) error {
	req := new(model.GenerateUUIDRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	strSerial := strconv.Itoa(req.Serial)
	logReq := model.Log{
		InternalID:   "",
		SubmissionID: "",
		StoreCode:    req.Store,
		Serials:      strSerial,
		LogText:      "",
		ErrText:      "",
		Posted:       false,
	}

	reciept, loginReq, serial, err := h.receiptRepo.FindReceiptData(req, h.companyInfo)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	logReq.InternalID = reciept.Header.ReceiptNumber

	resp, err := h.apiClient.GenerateUUID(*reciept)
	if utils.CheckErr(&err) {
		logReq.LogText = "failed to generate uuid"
		logReq.ErrText = err.Error()
		_, err = h.logRepo.ELogInsert(&logReq)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	body := model.ReceiptSubmitRequest{
		Receipts: []model.Receipt{resp.Receipt},
	}

	submitResp, err := h.apiClient.SubmitReceitps(*loginReq, body)
	if utils.CheckErr(&err) {
		logReq.LogText = "failed to submit receipts"
		logReq.ErrText = err.Error()
		_, err = h.logRepo.ELogInsert(&logReq)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	mJsong, _ := json.Marshal(resp.Receipt)
	updateReq := model.ReceiptUpdateRequest{
		Serial:      serial,
		RequestBody: string(mJsong),
		Posted:      false,
		Uuid:        resp.Uuid,
	}
	if len(submitResp.AcceptedDocuments) > 0 {
		updateReq.Posted = true
	}
	err = h.receiptRepo.ReceiptUpdate(updateReq)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	logReq.LogText = "success"
	logReq.SubmissionID = submitResp.SubmissionId
	_, err = h.logRepo.ELogInsert(&logReq)

	log.Debug().Interface("submit", submitResp).Msg("submitted invoices")
	return c.JSON(http.StatusOK, resp.Uuid)
}
func (h *Handler) ReceiptPost(c echo.Context) error {
	// _, body, err := h.receiptRepo.EInvoiceListUnposted()
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// resp, err := h.apiClient.SubmitReceitps(*body)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// postedUUIDs := []string{}
	// if len(resp.AcceptedDocuments) > 0 {
	// 	for _, v := range resp.AcceptedDocuments {
	// 		postedUUIDs = append(postedUUIDs, v.UUID)
	// 	}
	// }

	// req := new(model.PostInvoicessRequest)
	// if err := c.Bind(req); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	// }
	// reciept, serial, err := h.receiptRepo.FindReceiptData(req, h.companyInfo)
	// if utils.CheckErr(&err) {

	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// resp, err := h.apiClient.GenerateUUID(*reciept)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	// mJsong, _ := json.Marshal(resp.Receipt)
	// updateReq := model.ReceiptUpdateRequest{
	// 	Serial:      serial,
	// 	RequestBody: string(mJsong),
	// 	Uuid:        resp.Uuid,
	// }
	// err = h.receiptRepo.ReceiptUpdate(updateReq)
	// if utils.CheckErr(&err) {

	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	return c.JSON(http.StatusOK, "resp")

}
