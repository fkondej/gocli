package coingecko

import (
	"time"

	"github.com/shopspring/decimal"
)

type PriceData struct {
	AssetSymbol string
	PriceUSD    decimal.Decimal
	Time        time.Time
}

func (c *CoingeckoClient) GetAssetPrices(assetIds []string) ([]PriceData, error) {
	response, err := c.requestSimplePrice(assetIds)
	if err != nil {
		return nil, err
	}
	result := []PriceData{}
	for assetSymbol, data := range response {
		result = append(result, PriceData{
			AssetSymbol: assetSymbol,
			PriceUSD:    data.PriceUSD,
			Time:        time.Unix(data.LastUpdatedAt, 0),
		})
	}
	return result, nil
}
