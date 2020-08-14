package model

import "errors"

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
	Rank       int     `json:"rank"`
}

type TopCoins struct {
	Coins []PriceRanking
}

func (t *TopCoins) GetTopCoins(limit int64) ([]PriceRanking, error) {

	if len(t.Coins) < 1 {
		return nil, errors.New("Server is down. Please check later")
	}

	if limit > int64(len(t.Coins)) {
		limit = int64(len(t.Coins))
	}
	topCoins := t.Coins[:limit]
	return topCoins, nil
}

func (t *TopCoins) AssignRank() {
	for i := 0; i < len(t.Coins); i++ {
		t.Coins[i].Rank = i + 1
	}
}
