package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func main() {
	router := gin.Default()
	router.GET("/api", showAll)
	router.GET("/api/:id", getapi)

	router.Run("localhost:8080") // listen and serve on
}

func showAll(c *gin.Context) {
	if len(assets) == 0 {
		url = "https://api.coincap.io/v2/assets"
		cryptoList := []string{"bitcoin", "ethereum", ""}
	}
	c.IndentedJSON(http.StatusOK, assets)
}

func getapi(c *gin.Context) {
	id := c.Param("id")

	response, err := http.Get("https://api.coincap.io/v2/assets/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintln(c.Writer, "Error:", response.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var asset Asset
	if err := json.Unmarshal(body, &asset); err != nil {
		log.Fatal(err)
	}
	assets = append(assets, asset)
	fmt.Fprintln(c.Writer, string(body))
	c.JSON(http.StatusOK, asset)

}
