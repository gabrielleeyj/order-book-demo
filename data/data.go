package data

import "fmt"

// based on the example input so the data.
// The presumption is that the example inputs are unformatted strings.
// I’ll parse the entire string and format it in memory to get the order format out first.

type Order struct {
	Id       int64
	Side     string
	Quantity int64
	Price    float64
}

type Orderbook struct {
	BidPricePoints []float64
	AskPricePoints []float64
	Bids           map[float64][]Order
	Asks           map[float64][]Order
}

func (ob *Orderbook) PrettyPrint() {

	fmt.Println("===============================")

	fmt.Println("ASKS")
	//var a []float64
	// (a, ob.askPricePoints))
	for _, price := range reverse(copySlice(ob.AskPricePoints)) {
		orders := ob.Asks[price]
		if len(orders) == 0 {
			continue
		}
		fmt.Printf("%f: ", price)
		for _, order := range orders {
			fmt.Printf("%d ", order.Quantity)
			//fmt.Printf("%d-%d ", order.id, order.quantity)
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------")

	for _, price := range ob.BidPricePoints {
		orders := ob.Bids[price]
		if len(orders) == 0 {
			continue
		}
		fmt.Printf("%f: ", price)
		for _, order := range orders {
			fmt.Printf("%d ", order.Quantity)
			//fmt.Printf("%d-%d ", order.id, order.quantity)
		}
		fmt.Println()
	}
	fmt.Println("BIDS")

	fmt.Println("===============================")

}

func reverse(s []float64) []float64 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func copySlice(s []float64) []float64 {
	c := make([]float64, len(s))
	copy(c, s)
	return c
}
