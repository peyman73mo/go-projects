package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func Execute() {

	// variables for flags
	var url string
	var findip bool = false

	//create new command
	cmd := cli.NewApp()
	cmd.Name = "DNS lookup and Get request of given URL"
	cmd.Usage = "This command sends a http (GET method) request or/and checks IP address of given URL"

	// declare command flags
	cmd.Flags = []cli.Flag{

		cli.StringFlag{
			Name:        "url", // --url
			Usage:       "destination url",
			Destination: &url,
		},
		cli.BoolFlag{
			Name:        "findIP", //--findIP
			Usage:       "DNS lookup",
			Destination: &findip,
		},
	}

	cmd.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "get request",
			Flags: cmd.Flags,
			Action: func(c *cli.Context) error {
				if strings.Contains(url, "https://www.") != true {
					url = "https://www." + url
				}
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

				if findip == true {
					url = strings.ReplaceAll(url, "https://www.", "")

					ips, err := net.LookupIP((url))
					if err != nil {
						log.Fatal(err)
					}
					for _, ip := range ips {
						fmt.Println(ip)
					}
					return nil
				}
				return nil
			},
		},
		{
			Name:  "dnslookup",
			Usage: "DNS lookup for given URL",
			Flags: cmd.Flags,
			Action: func(c *cli.Context) error {

				ips, err := net.LookupIP(url)
				if err != nil {
					log.Fatal(err)
				}
				for _, ip := range ips {
					fmt.Println(ip)
				}
				return nil
			},
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
