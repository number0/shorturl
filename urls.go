package main

import (
	"log"

	"github.com/boltdb/bolt"
)

// URL ...
type URL struct {
	id  string
	url string
}

// ID ...
func (u URL) ID() string {
	return u.id
}

// URL ...
func (u URL) URL() string {
	return u.url
}

// Save ...
func (u URL) Save() error {
	log.Printf("u: %v", u)
	if u.id == "" || u.url == "" {
		log.Printf("u is nil :/")
		return nil
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			log.Printf("create bucket failed: %s", err)
			return err
		}

		err = b.Put([]byte(u.id), []byte(u.url))
		if err != nil {
			log.Printf("put key failed: %s", err)
			return err
		}

		return nil
	})

	return err
}

// LookupURL ...
func LookupURL(id string) (u URL, ok bool) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		if b == nil {
			return nil
		}

		v := b.Get([]byte(id))
		if v != nil {
			u = URL{id: id, url: string(v)}
			ok = true
		}

		return nil
	})

	if err != nil {
		log.Printf("error looking up url for %s: %s", id, err)
	}

	return
}

// NewURL ...
func NewURL(url string) (u URL, err error) {
	log.Printf("NewURL: %v", url)
	var id string

	for {
		// TODO: Make length (5) configurable
		id = RandomString(5)
		_, ok := LookupURL(id)
		if ok {
			continue
		} else {
			break
		}
	}

	u = URL{id: id, url: url}
	err = u.Save()

	return
}
