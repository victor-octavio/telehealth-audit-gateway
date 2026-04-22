package main

import (
	"github.com/victor-octavio/telehealth-audit-api/gateway/internal/handlers"
	"github.com/victor-octavio/telehealth-audit-api/gateway/internal/repositories"
	"github.com/victor-octavio/telehealth-audit-api/gateway/internal/usecases"
	"github.com/victor-octavio/telehealth-audit-api/gateway/pkg/fabric"
	"log"
)

func main() {

	cfg := fabric.Load()
	gateway, err := fabric.NewGateway(cfg)

	if err != nil {
		log.Fatalf("Error starting hyperledger fabric connection: %v", err)
	}

	defer gateway.Close()
	contract := gateway.Contract()

	repo := repositories.New(repositories.Options{
		Contract: contract,
	})

	usecases := usecases.New(usecases.Options{
		Repo: repo,
	})

	handlers := handlers.New(handlers.Options{
		UC: usecases,
	})

	if err := handlers.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
