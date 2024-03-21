// a performance problem with a database.
// caching to be implemented with map collection type in Go.
// there is a global limit on the number of items in th shared cache.
// both items in the same shared cache use the same ID, so we need a way to separate them.
// find a way to get items from the cache.

package main

import (
	"fmt"
)

// create Global Limit size as const
const GlobalLimit = 100

// max cacahe size that is 10X the global limit
const maxCacheSize int = 10 * GlobalLimit

// cache prefixes
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

// declare cache as map with string and values as strings
var cache map[string]string

// function to get items from the cache
func cacheGet(key string) string {
	return cache[key]
}

// function to set items into the cache
func cacheInsert(key, val string) {
	// make sure the cache is not going over the max cache size
	if len(cache)+1 >= maxCacheSize {
		return
	}
	cache[key] = val
}

// function to get a book from the cache
func getBook(isbn string) string {
	// create a unique key for the cached book
	return cacheGet(CacheKeyBook + isbn)
}

// function to add book to cache
func setBook(isbn, name string) {
	// use the cache prefix to create a unique key
	cacheInsert(CacheKeyBook+isbn, name)
	return
}

// function to get CD from cache
func getCD(sku string) string {
	// create unique key for CD
	return cacheGet(CacheKeyCD + sku)
}

// function to add book to cache
func setCD(sku, title string) {
	// create a unique identity to be used in the shared cache
	cacheInsert(CacheKeyCD+sku, title)
	return
}

func main() {
	// initialize a cache by creating a map
	cache = make(map[string]string)

	// add a book to the cache
	setBook("1234-5678", "Dune Messiah")
	// add a CD to the cache
	setCD("1234-5678", "Die Another Day")

	// get the book from the cache
	fmt.Println("Book :", getBook("1234-5678"))
	// get the CD from the cache
	fmt.Println("CD :", getCD("1234-5678"))
}
