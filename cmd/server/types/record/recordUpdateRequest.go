package Record

type RecordUpdateRequest struct {
	RecordType  string `json:"RecordType,omitempty"`
	Name        string `json:"name,omitempty"`
	Date        int64  `json:"date,omitempty"`
	Lot         string `json:"lot,omitempty"`
	Result      string `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
	NextDate    int64  `json:"nextDate,omitempty"`
}
