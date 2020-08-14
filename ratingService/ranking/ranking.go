package ranking

import (
	"errors"
	"sort"
	"topcoin/ratingservice/model"
)

var priceRank model.TopCoins

// Converts map data to array and sorts it according to Price
func GenerateRanking(priceList map[string]model.Price) {
	priceRank.Coins = nil
	for _, value := range priceList {
		priceRank.Coins = append(priceRank.Coins, model.PriceRanking{Name: value.Name, Symbol: value.Symbol, PriceInUSD: value.Quote.USD.Price})
	}

	sort.Slice(priceRank.Coins, func(i, j int) bool {
		return priceRank.Coins[i].PriceInUSD > priceRank.Coins[j].PriceInUSD
	})
	priceRank.AssignRank()
}

// Get top Coins based on limit
func GetTopCoins(limit int64) ([]model.PriceRanking, error) {
	if limit == 0 {
		err := errors.New("Invalid Limit")
		return nil, err
	}

	topcoins, err := priceRank.GetTopCoins(limit)
	return topcoins, err
}
