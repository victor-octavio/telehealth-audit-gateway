package diagnosis

type DiagnosisRequest struct {
	ID          string `json:"id,omitempty"`
	PatientID   string `json:"patient_id,omitempty"`
	PhysicianID string `json:"physician_id,omitempty"`
	Diagnosis   string `json:"diagnosis,omitempty"`
	Observation string `json:"observation,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
