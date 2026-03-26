package diagnosis

import (
	models "github.com/victor-octavio/telehealth-audit-api/internal/models/diagnosis"
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories"
)

type IDiagnosis interface {
	GetById()
	GetHistory()
	Add(req models.DiagnosisRequest) error
}

type DiagnosisImpl struct {
	Repository *repositories.Repositories
}

func New(repo *repositories.Repositories) IDiagnosis {
	return &DiagnosisImpl{
		Repository: repo,
	}
}

func (d *DiagnosisImpl) Add(req models.DiagnosisRequest) error {
	return d.Repository.Diagnosis.Add(req)
}

func (d *DiagnosisImpl) GetHistory() {

}

func (d *DiagnosisImpl) GetById() {

}
