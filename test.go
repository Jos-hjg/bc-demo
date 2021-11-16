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

	changeValue := "none"
	if err := db.Put([]byte(key), []byte(changeValue), nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("change key(\"%s\") to \"%s\"", key, changeValue)
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data, string(data))

}
