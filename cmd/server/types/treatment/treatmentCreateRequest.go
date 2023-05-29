package treatment

type TreatmentCreateRequest struct {
	TreatmentType string `json:"treatmentType"`
	Name          string `json:"name"`
	Date          int64  `json:"date"`
	Lot           string `json:"lot,omitempty"`
	Result        string `json:"result,omitempty"`
	Description   string `json:"description,omitempty"`
	Notes         string `json:"notes,omitempty"`
	RecurringRule string `json:"recurringRule,omitempty"`
}
