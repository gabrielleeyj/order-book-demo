package main

import (
	"fmt"
	"order-book-demo/data"
)

var book data.Orderbook

func main() {
	fmt.Println("Starting application...")
	// init book to memory
	book = data.Orderbook{}
	book.Bids = map[float64][]data.Order{}
	book.Asks = map[float64][]data.Order{}

	// run the dataset of strings created to book and prints to console
	readMarketData()
	fmt.Println("Ending application...")
}
