package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	dbPath := "leveldb"
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := "mango"
	if err := db.Put([]byte(key), []byte("yellow"), nil); err != nil {
		log.Fatal(err)
	}
	log.Println("put success!")

	data, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data, string(data))

}
