package av

type GlobalQuote struct {
	Quote struct {
		Symbol          string  `json:"1. symbol"`
		Open            float64 `json:"2. open,string"`
		High            float64 `json:"3. high,string"`
		Low             float64 `json:"4. low,string"`
		Price           float64 `json:"5. price,string"`
		Volume          int64   `json:"6. volume,string"`
		LatesTradingDay string  `json:"07. latest trading day"`
		PreviousClose   float64 `json:"08. previous close,string"`
		Change          float64 `json:"09. change,string"`
		ChangePercent   string  `json:"10. change percent"`
	} `json:"Global Quote"`
}
