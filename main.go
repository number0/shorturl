package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/namsral/flag"
)

var (
	cfg Config
	db  *bolt.DB
)

func main() {
	var (
		config  string
		dbpath  string
		baseurl string
		bind    string
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&dbpath, "dbpath", "urls.db", "Database path")
	flag.StringVar(&baseurl, "baseurl", "", "Base URL for display purposes")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.Parse()

	var err error
	db, err = bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: Abstract the Config and Handlers better
	cfg.baseURL = baseurl

	NewServer(bind, cfg).ListenAndServe()
}
