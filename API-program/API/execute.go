package API

import (
	"github.com/gin-gonic/gin"
)

type DataAsset struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}
type Asset struct {
	Data      DataAsset `json:"data"`
	TimeStamp int       `json:"timestamp"`
}

var assets []Asset

func Execute() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/api", showAll)
	router.GET("/api/:id", getapi)
	router.GET("/show/:id", showID)

	router.Run("localhost:8080") // listen and serve on
}
