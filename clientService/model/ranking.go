package model

type PriceRanking struct {
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
	PriceInUSD float64 `json:"price_usd"`
	Rank       int     `json:"rank"`
}

type APIResponse struct {
	Status   bool           `json:"status"`
	Response []PriceRanking `json:"response"`
	Error    string         `json:"error"`
}
