package app

import "github.com/urfave/cli"

// Return the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "First Shell Application"
	app.Usage = "Search the Server IP on Web"
	return app
}
