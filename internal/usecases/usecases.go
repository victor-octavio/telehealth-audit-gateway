package usecases

import (
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories"
	"github.com/victor-octavio/telehealth-audit-api/internal/usecases/diagnosis"
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
