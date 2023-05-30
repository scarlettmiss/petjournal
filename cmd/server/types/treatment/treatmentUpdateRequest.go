package treatment

type TreatmentUpdateRequest struct {
	TreatmentType string `json:"treatmentType,omitempty"`
	Name          string `json:"name,omitempty"`
	Date          int64  `json:"date,omitempty"`
	Lot           string `json:"lot,omitempty"`
	Result        string `json:"result,omitempty"`
	Description   string `json:"description,omitempty"`
	Notes         string `json:"notes,omitempty"`
	RecurringRule string `json:"recurringRule,omitempty"`
}
