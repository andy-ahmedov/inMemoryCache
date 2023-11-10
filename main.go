package main

import (
	"fmt"

	"github.com/andy-ahmedov/inMemoryCache/cache/cache"
)

func main() {
	c := cache.New()

	c.Set("fuch", 326)
	fmt.Println(c.Get("fuch"))

	c.Delete("fuch")
	fmt.Println(c.Get("fuch"))
}
