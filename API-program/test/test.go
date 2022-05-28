package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TestData struct {
	ID           int    `json:"id"`
	RandomNumber int    `json:"random_number"`
	Name         string `json:"name"`
	Time         string `json:"time"`
}

var data []TestData

func main() {
	data = []TestData{
		{1, 1, "name1", "time1"},
		{2, 2, "name2", "time2"},
		{3, 3, "name3", "time3"},
	}

	http.HandleFunc("/api", showAll)
	http.HandleFunc("/set", setData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data)
}
func setData(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	num, _ := strconv.Atoi(r.URL.Query().Get("num"))

	data = append(data, TestData{
		len(data) + 1,
		num,
		name,
		"time" + string(len(data)+1),
	})
}
