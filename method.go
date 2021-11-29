package main

import (
	"bufio"
	"fmt"
	"log"
	"order-book-demo/data"
	"os"
	"sort"
	"strconv"
	"strings"
)

var order data.Order

func readMarketData() {
	// read market data
	file, err := os.Open("./data/orderData.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		messageType := strings.Split(scanner.Text(), ",")[0]
		fmt.Println(">> Message received:", scanner.Text())
		switch messageType {
		case "A":
			// add an order
			order, err := parseOrder(scanner.Text())
			if err != nil {
				break
			}
			if order.Id == 200968 {
				fmt.Println("debug")
			}
			addOrder(order)
			book.PrettyPrint()

		case "X":
			// cancel an order
			order, err := parseOrder(scanner.Text())
			if err != nil {
				break

			}

			cancelOrder(order)
			book.PrettyPrint()

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
}

func addOrder(receivedOrder data.Order) {
	// add order to orderbook
	switch receivedOrder.Side {
	case "B":
		// add to bids
		book.Bids[receivedOrder.Price] = append(book.Bids[receivedOrder.Price], receivedOrder)
		book.BidPricePoints = appendIfUnique(book.BidPricePoints, receivedOrder.Price)
		sort.Slice(book.BidPricePoints, func(i, j int) bool {
			return book.BidPricePoints[i] > book.BidPricePoints[j]
		})

	case "S":
		// add to asks

		book.Asks[receivedOrder.Price] = append(book.Asks[receivedOrder.Price], receivedOrder)
		book.AskPricePoints = appendIfUnique(book.AskPricePoints, receivedOrder.Price)
		sort.Slice(book.AskPricePoints, func(i, j int) bool {
			return book.AskPricePoints[i] < book.AskPricePoints[j]
		})

	}
	trade(receivedOrder)
}

func cancelOrder(receivedOrder data.Order) {
	// cancel order
	switch receivedOrder.Side {
	case "B":
		book.Bids[receivedOrder.Price] = removeOrder(book.Bids[receivedOrder.Price], receivedOrder)
	case "S":
		book.Asks[receivedOrder.Price] = removeOrder(book.Asks[receivedOrder.Price], receivedOrder)
	}
}

func trade(receivedOrder data.Order) {
	switch receivedOrder.Side {
	case "B":
		for _, price := range book.AskPricePoints {
			if price <= receivedOrder.Price {
				// trade
				for _, _ = range book.Asks[price] {
					//fmt.Println(k)
					askOrder := book.Asks[price][0]
					if receivedOrder.Quantity <= 0 {
						break
					}
					if askOrder.Quantity == receivedOrder.Quantity {
						// trade
						fmt.Println(">>>>> TRADE: Order", receivedOrder.Id, "Buy", receivedOrder.Quantity, "at", price, "fulfilled by Order", askOrder.Id)
						askOrder.Quantity -= receivedOrder.Quantity
						receivedOrder.Quantity = 0
						book.Bids[receivedOrder.Price] = removeOrder(book.Bids[receivedOrder.Price], receivedOrder)
						book.Asks[askOrder.Price] = removeOrder(book.Asks[askOrder.Price], askOrder)
						// book.asks[price] = updateOrderInList(book.asks[price], askOrder)

					}
					if askOrder.Quantity > receivedOrder.Quantity {
						// partial trade
						fmt.Println(">>>>> PARTIAL TRADE: Order", receivedOrder.Id, "Buy", receivedOrder.Quantity, "at", price, "fulfilled by Order", askOrder.Id)
						askOrder.Quantity -= receivedOrder.Quantity
						//askOrder.quantity -= receivedOrder.quantity
						receivedOrder.Quantity = 0
						book.Bids[receivedOrder.Price] = removeOrder(book.Bids[receivedOrder.Price], receivedOrder)
						book.Asks[askOrder.Price] = updateOrderInList(book.Asks[askOrder.Price], askOrder)
					}
					if askOrder.Quantity < receivedOrder.Quantity {
						// partial trade
						fmt.Println(">>>>> PARTIAL TRADE: Order", receivedOrder.Id, "Buy", askOrder.Quantity, "at", price, "fulfilled by Order", askOrder.Id)
						receivedOrder.Quantity -= askOrder.Quantity
						askOrder.Quantity = 0
						book.Asks[askOrder.Price] = removeOrder(book.Asks[askOrder.Price], askOrder)
						book.Bids[receivedOrder.Price] = updateOrderInList(book.Bids[receivedOrder.Price], receivedOrder)
					}
				}
			}
		}

	case "S":
		for _, price := range book.BidPricePoints {
			if price >= receivedOrder.Price {

				// trade
				for _, _ = range book.Bids[price] {
					bidOrder := book.Bids[price][0]
					if receivedOrder.Quantity <= 0 {
						break
					}
					if bidOrder.Quantity == receivedOrder.Quantity {
						// trade
						fmt.Println(">>>>> TRADE: Order", receivedOrder.Id, "Sell", receivedOrder.Quantity, "at", price, "fulfilled by Order", bidOrder.Id)
						bidOrder.Quantity -= receivedOrder.Quantity
						receivedOrder.Quantity = 0
						book.Asks[receivedOrder.Price] = removeOrder(book.Asks[receivedOrder.Price], receivedOrder)
						book.Bids[bidOrder.Price] = removeOrder(book.Bids[bidOrder.Price], bidOrder)

					}
					if bidOrder.Quantity > receivedOrder.Quantity {
						// partial trade
						fmt.Println(">>>>> PARTIAL TRADE: Order", receivedOrder.Id, "Sell", receivedOrder.Quantity, "at", price, "fulfilled by Order", bidOrder.Id)
						bidOrder.Quantity -= receivedOrder.Quantity
						receivedOrder.Quantity = 0
						book.Asks[receivedOrder.Price] = removeOrder(book.Asks[receivedOrder.Price], receivedOrder)
						book.Bids[bidOrder.Price] = updateOrderInList(book.Bids[bidOrder.Price], bidOrder)
					}
					if bidOrder.Quantity < receivedOrder.Quantity {
						// partial trade
						fmt.Println(">>>>> PARTIAL TRADE: Order", receivedOrder.Id, "Sell", bidOrder.Quantity, "at", price, "fulfilled by Order", bidOrder.Id)
						receivedOrder.Quantity -= bidOrder.Quantity
						bidOrder.Quantity = 0
						book.Bids[bidOrder.Price] = removeOrder(book.Bids[bidOrder.Price], bidOrder)
						book.Asks[receivedOrder.Price] = updateOrderInList(book.Asks[receivedOrder.Price], receivedOrder)
					}
				}
			}
		}
	}
}

func parseOrder(receivedOrder string) (data.Order, error) {
	orderParts := strings.Split(receivedOrder, ",")
	o := data.Order{
		Id:       parseInt(orderParts[1]),
		Side:     orderParts[2],
		Quantity: parseInt(orderParts[3]),
		Price:    parseFloat(orderParts[4]),
	}
	if o.Quantity < 1 {
		return data.Order{}, fmt.Errorf("quantity must be greater than 0")
	}
	return o, nil
}

func removeOrder(orders []data.Order, receivedOrder data.Order) []data.Order {
	// remove order from list
	for i, order := range orders {
		if order.Id == receivedOrder.Id {
			orders = append(orders[:i], orders[i+1:]...)
			return orders
		}
	}
	return orders
}

func updateOrderInList(orders []data.Order, receivedOrder data.Order) []data.Order {
	for i, order := range orders {
		if order.Id == receivedOrder.Id {
			orders[i] = receivedOrder
			return orders
		}
	}
	return orders
}

func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func appendIfUnique(slice []float64, value float64) []float64 {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}
