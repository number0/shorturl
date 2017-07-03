package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/namsral/flag"
)

var (
	db  *bolt.DB
	cfg Config
)

func main() {
	var (
		config string
		dbpath string
		bind   string
		fqdn   string
	)

	flag.StringVar(&config, "config", "", "config file")
	flag.StringVar(&dbpath, "dbpth", "urls.db", "Database path")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.StringVar(&fqdn, "fqdn", "localhost", "FQDN for public access")
	flag.Parse()

	// TODO: Abstract the Config and Handlers better
	cfg.FQDN = fqdn

	var err error
	db, err = bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	NewServer(bind, cfg).ListenAndServe()
}
