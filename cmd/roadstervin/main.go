package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jfoster/eunos/roadster"
)

func main() {
	flag.Parse()

	vin, err := roadster.ParseVIN(flag.Args()[0], "vins.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: Eunos Roadster %s 1.%di Manufactured %s\n", vin, vin.Model, vin.Engine, vin.Date)
}
