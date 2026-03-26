package repositories

import (
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	models "github.com/victor-octavio/telehealth-audit-api/internal/models/diagnosis"
)

type DiagnosisRepository struct {
	contract *client.Contract
}

func New(contract *client.Contract) *DiagnosisRepository {
	return &DiagnosisRepository{
		contract: contract,
	}
}

func (d *DiagnosisRepository) GetById() {

}

func (d *DiagnosisRepository) GetHistory() {

}

func (d *DiagnosisRepository) Add(req models.DiagnosisRequest) error {
	_, err := d.contract.SubmitTransaction(
		"InsertDiagnostic",
		req.ID,
		req.PatientID,
		req.PhysicianID,
		req.Diagnosis,
		req.Observation,
	)

	if err != nil {
		return fmt.Errorf("error during transaction: %w", err)
	}

	return nil
}
