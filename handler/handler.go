package handler

import (
	"eta/repo"
)

type Handler struct {
	userRepo    repo.UserRepo
	orderRepo   repo.OrderRepo
	invoiceRepo repo.InvoiceRepo
	receiptRepo repo.ReceiptRepo
	storeRepo   repo.StoreRepo
}

func NewHandler(userRepo repo.UserRepo, orderRepo repo.OrderRepo, invoiceRepo repo.InvoiceRepo, receiptRepo repo.ReceiptRepo, storeRepo repo.StoreRepo) *Handler {
	return &Handler{
		userRepo:    userRepo,
		orderRepo:   orderRepo,
		invoiceRepo: invoiceRepo,
		receiptRepo: receiptRepo,
		storeRepo:   storeRepo,
	}
}
