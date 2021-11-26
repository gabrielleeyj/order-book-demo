package data

import (
	"fmt"
	"strings"
)

// based on the example input so the data.
// The presumption is that the example inputs are unformatted strings.
// I’ll parse the entire string and format it in memory to get the order format out first.

orders := "A,100000,S,1,1075 A,100001,B,9,1000 A,100002,B,30,975 A,100003,S,10,1050 A,100004,B,10,950 A,100005,S,2,1025 A,100006,B,1,1000 X,100004,B,10,950 A,100007,S,5,1025 A,100008,B,3,1050 X,100008,B,3,1050 X,100005,S,2,1025"

type Orders struct {
	action  string
	orderId int
	side int
	price int
}

func formatData(data) []string{
	// split the string by space
	order := strings.Split(data, " ")
	return order
}
