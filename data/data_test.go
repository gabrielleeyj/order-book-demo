package data

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var orders []string
var temp []string

func TestFormatString(t *testing.T) {
	data := "A,100000,S,1,1075 A,100001,B,9,1000 A,100002,B,30,975 A,100003,S,10,1050 A,100004,B,10,950 A,100005,S,2,1025 A,100006,B,1,1000 X,100004,B,10,950 A,100007,S,5,1025 A,100008,B,3,1050 X,100008,B,3,1050 X,100005,S,2,1025"
	orders = strings.Split(data, " ")
}

func TestFormatData(t *testing.T) {
	var orderData []Order
	for _, s := range orders {
		temp = strings.Split(s, ",")

		if temp[0] == "A" {
			oid, err := strconv.Atoi(temp[1])
			if err != nil {
				return
			}
			share, err := strconv.Atoi(temp[3])
			if err != nil {
				return
			}
			price, err := strconv.Atoi(temp[4])
			if err != nil {
				return
			}

			newOrder := Order{
				Id:       int64(oid),
				Quantity: int64(share),
				Price:    float64(price),
			}

			orderData = append(orderData, newOrder)
		}
		fmt.Println(orderData)
		if temp[0] == "X" {
			// find the index of the element in orderData
			// then remove it from the slice via index
			// shift all the slices from the right to the left.

			// get the id of the order to cancel
			oid, err := strconv.Atoi(temp[1])
			if err != nil {
				return
			}
			fmt.Println("X", oid)

			// loop through the orders to find the id
			// might cause race condition
			for i, v := range orderData {
				temp := v
				if temp.Id == int64(oid) {
					fmt.Println("found oid", oid)
					fmt.Println("found orderId", temp.Id)
					orderData[i] = Order{}
				}
			}

		}
	}
	fmt.Println("final order")
	fmt.Println(orderData)
}

func TestCopySlice(*testing.T) {
	slice := []float64{1, 2, 3, 4, 5}
	sliceCopy := copySlice(slice)
	if sliceCopy[0] != 1 {
		panic("copySlice failed")
	}
}

func TestReverse(*testing.T) {
	slice := []float64{1, 2, 3, 4, 5}
	reverse(slice)
	if slice[0] != 5 {
		panic("reverseSlice failed")
	}
}
