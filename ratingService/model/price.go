package model

type Price struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

type PriceRanking struct {
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
	PriceInUSD float64 `json:"price_usd"`
}
