package domain

type HoliDayResponse struct {
	Status  string    `json:"status"`
	Holiday []HoliDay `json:"data"`
}

type HoliDay struct {
	Date        string `json:"date" xml:"date"`
	Title       string `json:"title" xml:"title"`
	Type        string `json:"type" xml:"type"`
	Inalienable bool   `json:"inalienable" xml:"inalienable"`
	Extra       string `json:"extra" xml:"extra"`
}
