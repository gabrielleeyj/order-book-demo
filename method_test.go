package main

import (
	d "order-book-demo/data"
	"testing"
)

func TestParseFloat(*testing.T) {
	if parseFloat("1") != 1 {
		panic("parseFloat failed")
	}
}

func TestParseInt(*testing.T) {
	if parseInt("1") != 1 {
		panic("parseInt failed")
	}
}

func TestUpdateOrderInList(*testing.T) {
	o1 := d.Order{
		Id:       1,
		Side:     "B",
		Quantity: 1,
		Price:    1,
	}

	orderList := []d.Order{o1}

	o2 := d.Order{
		Id:       1,
		Side:     "B",
		Quantity: 2,
		Price:    1,
	}

	updateOrderInList(orderList, o2)
	if orderList[0].Quantity != 2 {
		panic("updateOrderInList failed")
	}
}

func TestAddOrder(*testing.T) {
	o1 := d.Order{
		Id:       1,
		Side:     "B",
		Quantity: 1,
		Price:    1,
	}

	o2 := d.Order{
		Id:       1,
		Side:     "S",
		Quantity: 1,
		Price:    1,
	}

	book = d.Orderbook{}
	book.Bids = map[float64][]d.Order{}
	book.Asks = map[float64][]d.Order{}
	addOrder(o1)
	if len(book.Bids) != 1 {
		panic("addOrder failed")
	}

	addOrder(o2)
	if len(book.Bids) != 1 {
		panic("addOrder failed")
	}
}

func TestCancelOrder(*testing.T) {
	o1 := d.Order{
		Id:       1,
		Side:     "B",
		Quantity: 1,
		Price:    1,
	}

	o2 := d.Order{
		Id:       1,
		Side:     "S",
		Quantity: 1,
		Price:    1,
	}

	book = d.Orderbook{}
	book.Bids = map[float64][]d.Order{}
	book.Asks = map[float64][]d.Order{}
	addOrder(o1)
	addOrder(o2)
	cancelOrder(o1)
	if len(book.Bids) != 1 {
		panic("cancelOrder failed")
	}

	cancelOrder(o2)
	if len(book.Bids) != 1 {
		panic("cancelOrder failed")
	}
}

func TestRemoveOrder(*testing.T) {
	o1 := d.Order{
		Id:       999,
		Side:     "B",
		Quantity: 777,
		Price:    888,
	}

	o2 := d.Order{
		Id:       1999,
		Side:     "S",
		Quantity: 1777,
		Price:    1888,
	}

	book = d.Orderbook{}
	book.Bids = map[float64][]d.Order{}
	book.Asks = map[float64][]d.Order{}
	addOrder(o1)
	addOrder(o2)
	book.Bids[o1.Price] = removeOrder(book.Bids[o1.Price], o1)
	if len(book.Bids[o1.Price]) != 0 {
		panic("removeOrder failed")
	}

	book.Asks[o2.Price] = removeOrder(book.Asks[o2.Price], o2)
	if len(book.Asks[o2.Price]) != 0 {
		panic("removeOrder failed")
	}
}

func TestParseOrder(*testing.T) {

	s := "A,200014,S,18,1005"

	o, _ := parseOrder(s)
	if o.Id != 200014 {
		panic("parseOrder failed")
	}

	s2 := "A,200014,S,0,1005"

	_, err := parseOrder(s2)
	if err == nil {
		panic("parseOrder failed")
	}
}

func TestAppendIfUnique(*testing.T) {
	slice := []float64{1, 2, 3, 4, 5}
	slice = appendIfUnique(slice, 2)
	if len(slice) != 5 {
		panic("appendIfUnique failed")
	}

	slice = appendIfUnique(slice, 6)
	if len(slice) != 6 {
		panic("appendIfUnique failed")
	}
}

func TestReadMarketData(*testing.T) {

	book = d.Orderbook{}
	book.Bids = map[float64][]d.Order{}
	book.Asks = map[float64][]d.Order{}
	readMarketData()
	if len(book.Bids) < 1 {
		panic("readMarketData failed")
	}
}
