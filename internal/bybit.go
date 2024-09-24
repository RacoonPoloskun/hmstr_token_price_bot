package internal

import (
	"context"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"os"
)

type BybitTokenPricer struct {
	client *bybit.Client
}

func NewTokenPricer() *BybitTokenPricer {
	return &BybitTokenPricer{
		client: bybit.NewBybitHttpClient(
			os.Getenv("BYBIT_API_KEY"),
			os.Getenv("BYBIT_API_SECRET"),
			bybit.WithBaseURL(bybit.MAINNET),
		),
	}
}

func (t *BybitTokenPricer) GetTokenLastPrice(tokenPair string) (string, error) {
	params := map[string]interface{}{"category": "linear", "symbol": tokenPair}
	serverResult, err := t.client.NewUtaBybitServiceWithParams(params).GetMarketTickers(context.Background())
	if err != nil {
		return "", err
	}

	resultMap := serverResult.Result.(map[string]interface{})
	lastPrice := resultMap["list"].([]interface{})[0].(map[string]interface{})["lastPrice"]

	return lastPrice.(string), nil
}
