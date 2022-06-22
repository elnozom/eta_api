package handler

import (
	"eta/model"
	"eta/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InvoicesListByStorePosted(c echo.Context) error {
	postedFilter, _ := strconv.ParseBool(c.QueryParam("active"))
	storeFilter, _ := strconv.Atoi(c.QueryParam("store"))

	result, err := h.invoiceRepo.ListEInvoicesByPosted(&postedFilter, &storeFilter)
	utils.CheckErr(&err)
	return c.JSON(http.StatusOK, result)
}

func (h *Handler) InvoicePost(c echo.Context) error {

	// _, err := utils.EtaLogin()
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusOK, err.Error())
	// }
	// // return c.JSON(http.StatusOK, token)

	serial, _ := strconv.ParseUint(c.Param("serial"), 0, 64)
	var invoice model.Invoice
	err := h.invoiceRepo.FindInvoiceData(&serial, &invoice)
	if utils.CheckErr(&err) {
		return c.JSON(http.StatusOK, err.Error())
	}
	invoice.DocumentTypeVersion = "1.0"
	invoice.DocumentType = "I"
	return c.JSON(http.StatusOK, invoice)

	// err = utils.SignDocument(&invoice)

	// _, err = utils.SubmitInvoice(&token)
	// if utils.CheckErr(&err) {
	// 	return c.JSON(http.StatusOK, err.Error())
	// }
}

// // NewSnapshot initilizes a SignedSnapshot with a given top level root
// // and targets objects
// func _signDocument(document *model.Invoice) error {
// 	fmt.Println(*document)
// 	jsonDocument, err := json.Marshal(document)
// 	if err != nil {
// 		return err
// 	}
// 	canonicalJson := http.CanonicalHeaderKey(string(jsonDocument))
// 	encodedJson, err := json.Marshal(canonicalJson)
// 	if err != nil {
// 		return err
// 	}

// 	cert, err := x509.ParseCertificate(encodedJson)
// 	if err != nil {
// 		return err
// 	}
// 	key, err := x509.ParseECPrivateKey(encodedJson)
// 	if err != nil {
// 		return err
// 	}
// 	der, err := cms.Sign(encodedJson, []*x509.Certificate{cert}, key)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(der)
// 	return nil
// }
