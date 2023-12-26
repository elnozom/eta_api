package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"eta/model"
	"eta/utils"
	"fmt"
	"net/http"

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
func (h *Handler) ToUUID(c echo.Context) error {
	req := new(model.ToUUDRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	hash := sha256.New()
	hash.Write([]byte(req.Invoice))
	hashValue := hash.Sum(nil)
	// reciept.Header.ReferenceUUID = strings.ReplaceAll(hashString, "n", "s")
	log.Debug().Interface("string", req.Invoice).Msg("vanon")

	// Convert the hash value from an array of 32 bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashValue)

	// reciept.Header.ReferenceUUID = "fb6caa0bf99bc1582de1df29b7db6d287ba5c7238cb6e980ca7c7c061035308d"
	return c.JSON(http.StatusInternalServerError, hashString)
}
func (h *Handler) ReceiptPost(c echo.Context) error {
	req := new(model.PostInvoicessRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	reciept, err := h.receiptRepo.FindReceiptData(req, h.companyInfo)
	if utils.CheckErr(&err) {

		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// var jsonData map[string]interface{}

	// receipt :=
	// request := model.ReceiptSubmitRequest{
	// 	Receipts: []model.Receipt{*reciept},
	// }

	// mJson, _ := json.Marshal(reciept)
	canonicalString := utils.Serialize(*reciept)
	// err = json.Unmarshal(mJson, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// canonicalString = strings.ReplaceAll(canonicalString, "\"\"\"", "\"\"")
	// canonicalString = strings.ReplaceAll(canonicalString, ",", "")
	// canonicalString = strings.ReplaceAll(canonicalString, ":", "")
	// canonicalString = strings.ReplaceAll(canonicalString, "{", "")
	// canonicalString = strings.ReplaceAll(canonicalString, "}", "")
	// log.Debug().Interface("holholholaaa", reciept).Msg("req")
	hash := sha256.New()
	hash.Write([]byte(canonicalString))
	hashValue := hash.Sum(nil)
	// reciept.Header.ReferenceUUID = strings.ReplaceAll(hashString, "n", "s")

	// Convert the hash value from an array of 32 bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashValue)
	log.Debug().Interface("hast", hashString).Interface("canonicalString", canonicalString).Msg("vanocanonicalString")

	// reciept.Header.Uuid = hashString
	// log.Debug().Interface("mapValue", mapValue).Msg("vanomapValue")
	// log.Debug().Interface("hashString", hashString).Msg("vanohashString")

	// reciept.Header.ReferenceUUID = "fb6caa0bf99bc1582de1df29b7db6d287ba5c7238cb6e980ca7c7c061035308d"
	// reciept["Header"].(map[string]interface{})["Uuid"] = hashString

	// canonicalJSON := toCanonicalJSON(jsonData, "")

	return c.JSON(http.StatusOK, map[string]interface{}{"receipts": []interface{}{*reciept}})
}
