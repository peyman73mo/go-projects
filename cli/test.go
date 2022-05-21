package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	var method string
	var url string

	cmd := cli.NewApp()
	cmd.Name = "HTTP request"
	cmd.Usage = "This command sends a http request to given url"

	cmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "method",
			Usage:       "http method",
			Destination: &method,
		},
		cli.StringFlag{
			Name:        "url",
			Usage:       "destination url",
			Destination: &url,
		},
	}

	cmd.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			fmt.Println("type: test --method=[get/GET/post/POST] --url=[url]")
			return nil
		}

		if method != "" && url != "" {

			if strings.Contains(url, "https://www.") != true && strings.Index(url, "https://www.") != 0 {
				fmt.Println("enter : --url=https://www.[host].[domain]")
				return nil
			}

			switch method {
			case "get", "GET":
				// fmt.Println("get method")
				respose, err := http.Get(url)
				if err != nil {
					fmt.Println("response error")
					log.Fatal(err)
				}
				defer respose.Body.Close()

				body, err := ioutil.ReadAll(respose.Body)
				if err != nil {
					fmt.Println("ioutil error")
					log.Fatal(err)
				}
				fmt.Println(string(body))

			case "post", "POST":
				fmt.Println("post method")
			}

		} else {
			fmt.Println("type: test --method=[get/GET/post/POST] --url=[url]")
		}

		return nil
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
