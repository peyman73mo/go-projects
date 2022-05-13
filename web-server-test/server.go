package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("server is running ", "\n----------------")

	router.HandleFunc("/", home)
	router.HandleFunc("/{route}", home)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("-------------------")
	switch r.Method {

	case "GET":
		path := r.URL.Path
		fmt.Println(r.URL.Path, r.Method, r.URL.Query())

		if path == "/" {
			path = "static/frames.html"
		} else {
			path = "static/" + r.URL.Path
		}
		df, err := ioutil.ReadFile(path)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("<h1>page not found</h1>"))
		}

		rw.Write(df)

	default:
		fmt.Println(r.URL.Path, r.Method, r.URL.Query())
		fmt.Println("something is wrong")

		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("<h1>page not found</h1>"))
	}

}
