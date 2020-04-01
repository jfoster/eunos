package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jfoster/eunos/roadster"
	"gopkg.in/yaml.v2"
)

func main() {
	flag.Parse()

	vin, err := roadster.ParseVIN(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("data.yml")
	if err != nil {
		log.Fatal(err)
	}

	var dates roadster.VinDates

	err = yaml.NewDecoder(file).Decode(&dates)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(dates)-2; i++ {
		x := dates[i]
		y := dates[i+1]

		if x.Sequence <= vin.Sequence && vin.Sequence < y.Sequence && x.Model == vin.Model && x.Engine == vin.Engine {
			fmt.Printf("%s: Eunos Roadster %s 1.%di Manufactured %s\n", vin, x.Model, x.Engine, x.Date)
			break
		}
	}
}
