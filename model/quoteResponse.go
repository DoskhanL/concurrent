package model

// QuoteResponse model
type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeVTD        float32
	ChangePercentVTD float32
	High             float32
	Low              float32
	Open             float32
}
