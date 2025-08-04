package domain

type HoliDayResponse struct {
	Status  string    `json:"status"`
	Holiday []HoliDay `json:"data"`
}

type HoliDay struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Inalienable bool   `json:"inalienable"`
	Extra       string `json:"extra"`
}
