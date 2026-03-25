package diagnosis

import (
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories"
)

type IDiagnosis interface {
	GetById()
	GetHistory()
	Add()
}

type DiagnosisImpl struct {
	Repository *repositories.Repositories
}

func New(repo *repositories.Repositories) IDiagnosis {
	return &DiagnosisImpl{
		Repository: repo,
	}
}

func (d *DiagnosisImpl) Add() {

}

func (d *DiagnosisImpl) GetHistory() {

}

func (d *DiagnosisImpl) GetById() {

}
