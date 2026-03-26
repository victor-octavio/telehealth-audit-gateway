package repositories

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	models "github.com/victor-octavio/telehealth-audit-api/internal/models/diagnosis"
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories/diagnosis"
)

type Repositories struct {
	Diagnosis interface {
		GetById()
		GetHistory()
		Add(req models.DiagnosisRequest) error
	}
}

type Options struct {
	Contract *client.Contract
}

func New(opts Options) *Repositories {
	return &Repositories{
		Diagnosis: repositories.New(opts.Contract),
	}
}
