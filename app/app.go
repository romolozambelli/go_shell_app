package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Return the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "First Shell Application"
	app.Usage = "Search the Server IP on Web"

	app.Commands = []cli.Command{
		{
			Name:  "ips",
			Usage: "Search for IPs on Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIp,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
