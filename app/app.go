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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Search for IPs on Internet",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "server",
			Usage:  "Search for Servers Name on Internet",
			Flags:  flags,
			Action: searchServers,
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

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name server

	if erro != nil {
		log.Fatal(erro)
	}

	for _, serversName := range servers {
		fmt.Println(serversName.Host)
	}
}
