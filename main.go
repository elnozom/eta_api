package main

import (
	"eta/config"
	"eta/db"
	"eta/handler"
	"eta/repo"
	"eta/router"
	"fmt"
)

func main() {
	r := router.New()
	v1 := r.Group("")
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	userRepo := repo.NewUserRepo(db)
	orderRepo := repo.NewOrderRepo(db)
	invoiceRepo := repo.NewInvoiceRepo(db)
	receiptRepo := repo.NewReceiptRepo(db)
	storeRepo := repo.NewStoreRepo(db)
	h := handler.NewHandler(
		userRepo,
		orderRepo,
		invoiceRepo,
		receiptRepo,
		storeRepo,
	)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	r.Logger.Fatal(r.Start(port))
}
