package repositories

import (
	"encoding/json"
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

func (d *DiagnosisRepository) GetHistory(ID string) ([]models.DiagnosisRequest, error) {
	history, err := d.contract.EvaluateTransaction("GetHistory", ID)

	if err != nil {
		return nil, fmt.Errorf("error during GetHistory request")
	}

	var result []models.DiagnosisRequest

	if err := json.Unmarshal(history, &result); err != nil {
		return nil, fmt.Errorf("error during json deconding")
	}

	return result, nil
}

func (d *DiagnosisRepository) Add(req models.DiagnosisRequest) error {
	_, err := d.contract.SubmitTransaction(
		"InsertDiagnostic",
		req.ID,
		req.PatientID,
		req.PhysicianID,
		req.Diagnosis,
		req.Observation,
		"",
		"",
	)

	if err != nil {
		return fmt.Errorf("error during transaction: %w", err)
	}

	return nil
}
