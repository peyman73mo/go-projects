package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	body, err := ioutil.ReadFile("static/home.html")
	if err != nil {
		log.Fatal(err)
	}
	c.Header("Content-Type", "text/html")
	c.Writer.Write(body)
}

func showAll(c *gin.Context) {
	if len(assets) == 0 {
		url := "https://api.coincap.io/v2/assets"
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		type AssetSlice struct {
			Data      []DataAsset `json:"data"`
			TimeStamp int         `json:"timestamp"`
		}
		asset := AssetSlice{}
		if err := json.Unmarshal(body, &asset); err != nil {
			log.Fatal(err)
		}
		for _, data := range asset.Data {
			assets = append(assets, Asset{Data: data, TimeStamp: asset.TimeStamp})
		}
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

	c.IndentedJSON(http.StatusOK, asset)

}

func showID(c *gin.Context) {
	id := c.Param("id")

	checkFlag := true
	for _, asset := range assets {
		if asset.Data.ID == id {
			c.IndentedJSON(http.StatusOK, asset)
			checkFlag = false
		}
	}
	if checkFlag {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
	}
}
