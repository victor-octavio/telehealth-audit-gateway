package usecases

import (
	"github.com/victor-octavio/telehealth-audit-api/gateway/internal/repositories"
	"github.com/victor-octavio/telehealth-audit-api/gateway/internal/usecases/diagnosis"
)

type Usecases struct {
	Diagnosis diagnosis.IDiagnosis
}

type Options struct {
	Repo *repositories.Repositories
}

func New(opts Options) *Usecases {
	return &Usecases{
		Diagnosis: diagnosis.New(opts.Repo),
	}
}
