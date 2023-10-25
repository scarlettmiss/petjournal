package Record

type RecordCreateRequest struct {
	RecordType  string `json:"RecordType"`
	Name        string `json:"name"`
	Date        int64  `json:"date"`
	Lot         string `json:"lot,omitempty"`
	Result      string `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
	NextDate    int64  `json:"nextDate,omitempty"`
}
