package handler

import (
	"eta/router/middleware"
	"eta/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	v1.GET("/health", h.CheckHealth)
	api := v1.Group("/api", jwtMiddleware)
	api.GET("/validate", h.ValidateUser)
	//auth routes
	v1.POST("api/login", h.Login)
	// cuurent user routes
	currentUser := api.Group("/me")
	currentUser.GET("", h.Me)
	// orders routes
	orders := api.Group("/orders")
	orders.POST("/convert/:serial", h.OrdersConvertToEta)
	orders.GET("", h.OrdersListByTransSerialStoreConvertedDate)
	// orders.GET("/post", h.PosOrdersListByStoreConvertedDate)
	// orders.GET("/converted", h.OrdersListConverted)

	// invoices routes
	invoices := api.Group("/invoices")
	invoices.GET("", h.InvoicesListByPosted)
	invoices.POST("/post/:serial", h.InvoicePost)

	// receipt routes
	receipt := api.Group("/receipts")
	receipt.GET("", h.ReceiptsListByPosted)
	receipt.POST("/post/:serial", h.ReceiptPost)

	//clobal routes
	api.GET("/stores", h.StoresListAll)
	// global
	api.POST("/upload", h.Upload)

}
