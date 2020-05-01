package main

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New(":22123")
	setErr := mc.Set(&memcache.Item{Key: "test-key", Value: []byte("test-value")})
	if setErr != nil {
		log.Println(setErr)
		return
	}

	it, getErr := mc.Get("test-key")
	if getErr != nil {
		log.Println(getErr)
		return
	}

	log.Println(it.Key, string(it.Value))

	deleteErr := mc.Delete("test-key")
	if deleteErr != nil {
		log.Println(deleteErr)
		return
	}

}
