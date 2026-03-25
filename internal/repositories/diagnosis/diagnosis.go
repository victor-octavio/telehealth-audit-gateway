package diagnosis

import "github.com/hyperledger/fabric-gateway/pkg/client"

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

func (d *DiagnosisRepository) Add() {

}
