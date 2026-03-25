package main

import (
	"log"

	"github.com/victor-octavio/telehealth-audit-api/internal/handlers"
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories"
	"github.com/victor-octavio/telehealth-audit-api/internal/usecases"
	"github.com/victor-octavio/telehealth-audit-api/pkg/fabric"
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
