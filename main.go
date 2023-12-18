package main

import (
	"eta/client"
	"eta/config"
	"eta/db"
	"eta/handler"
	"eta/repo"
	"eta/router"
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	r := router.New()
	v1 := r.Group("")

	state, err := config.LoadState("./config")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the state config")
	}
	config, err := config.LoadConfig("./config", state.State)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the config")
	}
	apiClient := client.NewApiClient(&config)

	resp, err := apiClient.Login()
	log.Debug().Interface("test", resp).Msg("hola")

	// store, err := db.InitDB(ctx, config.DBSource)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot connect to db")
	// }

	db, err := db.New(config.DBSource)
	if err != nil {
		panic(err)
	}
	userRepo := repo.NewUserRepo(db)
	orderRepo := repo.NewOrderRepo(db)
	invoiceRepo := repo.NewInvoiceRepo(db)
	receiptRepo := repo.NewReceiptRepo(db)
	storeRepo := repo.NewStoreRepo(db)
	logRepo := repo.NewLogRepo(db)
	companyRepo := repo.NewCompanyRepo(db)
	companyInfo, err := companyRepo.Find()
	if err != nil {
		panic(err)
	}
	h := handler.NewHandler(
		userRepo,
		orderRepo,
		invoiceRepo,
		receiptRepo,
		storeRepo,
		logRepo,
		companyInfo,
	)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Port)

	r.Logger.Fatal(r.Start(port))
}
