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

type PriceList struct {
	Status struct {
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
	Data map[string]Price
}
