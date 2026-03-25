package repositories

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories/diagnosis"
)

type Repositories struct {
	Diagnosis interface {
		GetById()
		GetHistory()
		Add()
	}
}

type Options struct {
	Contract *client.Contract
}

func New(opts Options) *Repositories {
	return &Repositories{
		Diagnosis: diagnosis.New(opts.Contract),
	}
}

