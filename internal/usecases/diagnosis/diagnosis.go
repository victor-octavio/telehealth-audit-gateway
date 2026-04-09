package diagnosis

import (
	models "github.com/victor-octavio/telehealth-audit-api/internal/models/diagnosis"
	"github.com/victor-octavio/telehealth-audit-api/internal/repositories"
)

type IDiagnosis interface {
	GetById(ID string) (*models.DiagnosisRequest, error)
	GetHistory(ID string) ([]models.DiagnosisRequest, error)
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

func (d *DiagnosisImpl) GetHistory(ID string) ([]models.DiagnosisRequest, error) {
	return d.Repository.Diagnosis.GetHistory(ID)
}

func (d *DiagnosisImpl) GetById(ID string) (*models.DiagnosisRequest, error) {
	return d.Repository.Diagnosis.GetById(ID)
}
