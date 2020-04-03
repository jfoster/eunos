package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/jfoster/eunos/roadster"
	_ "github.com/jfoster/eunos/roadster"
)

type Data struct {
	Vin   *roadster.VIN
	Error error
}

func main() {
	var cfg struct {
		Port  string `env:"PORT" envDefault:"8080"`
		Table string `env:"TABLE" envDefault:"vins.yml"`
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/roadster", func(w http.ResponseWriter, r *http.Request) {
		var data = Data{
			Error: nil,
		}

		if v := r.FormValue("vin"); v != "" {
			vin, err := roadster.ParseVIN(v, cfg.Table)
			if err != nil {
				data.Error = err
			}
			if vin != nil {
				data.Vin = vin
			}
		}

		tmpl.Execute(w, data)
	})

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
