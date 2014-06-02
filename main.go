package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/gzip"
)

func main() {
	var config, port, api string

	flag.StringVar(&config, "c", "config.json", "Config file")
	flag.StringVar(&port, "p", "8080", "Port")
	flag.StringVar(&api, "a", "DEFAULT", "API key")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [options]

Barycenter serves a JSON configuration file over HTTP
using basic authentication (so run it over SSL).

Run an endpoint as follows:

  %s -c config.json -a DEFAULT -p 8080

You can then make a request against the endpoint.

  curl -u DEFAULT: 127.0.0.1:8080

OPTIONS:
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	json, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalf("Could not read configuration: %s", err)
	}

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(auth.Basic(api, ""))

	m.Get("/", func(w http.ResponseWriter, req *http.Request) string {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return string(json)
	})

	http.ListenAndServe(":"+port, m)
}
