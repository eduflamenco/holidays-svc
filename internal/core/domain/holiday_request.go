package domain

type HolidayPaymentRequest struct {
	HolidayType string `json:"type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
