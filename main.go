package main

import (
	"fmt"
	"log"
	"os"
	"shell-app/app"
)

func main() {
	fmt.Println("-- Application Starting--")

	application := app.Generate()

	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
