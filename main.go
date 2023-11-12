package main

import (
	"fmt"
	"log"
	"time"

	"github.com/andy-ahmedov/inMemoryCache/cache/cache"
)

func main() {
	c := cache.New()

	durat := time.Second * 2

	c.Set("userId", 42, durat)
	userId, err := c.Get("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("REZULTAT", userId)

	time.Sleep(time.Second)

	c.Set("userId", 42, time.Second*4)

	time.Sleep(time.Second * 2)

	c.Set("sda", 134, durat)
	c.Set("sa", 13, durat)
	Id, err := c.Get("sda")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("REZULTAT", Id)
	c.Set("da", 34, durat)

	userId, err = c.Get("userId")
	fmt.Println(userId)
	if err != nil {
		log.Fatal(err)
	}
}
