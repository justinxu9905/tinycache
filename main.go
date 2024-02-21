package main

import (
	"fmt"
	"log"
	"net/http"
	"tinycache/tinycache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	tinycache.NewGroup("scores", 2<<10, tinycache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := tinycache.NewHTTPPool(addr)
	log.Println("tinycache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
