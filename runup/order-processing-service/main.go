package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	orderId = iota
	userId
	orderAmount
	orderStatus
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sep := ","
	var orders [][]string

	for sc.Scan() {
		input := sc.Text()
		if input == "" {
			break
		}
		order := strings.Split(input, sep)
		orders = append(orders, order)
	}

	var rev float64
	revPerUser := make(map[string]float64)
	var paidOrders int

	for _, order := range orders {
		switch order[orderStatus] {
		case "paid":
			amount, _ := strconv.ParseFloat(order[orderAmount], 64)
			rev += amount
			revPerUser[order[userId]] += amount
			paidOrders++
		}
	}

	var rpu strings.Builder
	for k, v := range revPerUser {
		rpu.WriteString(fmt.Sprintf("%v: %v", k, v))
		rpu.WriteRune('\n')
	}

	fmt.Fprintf(os.Stdout, "Total revenue: %v\nPaid orders: %v\nRevenue per user: \n%v", rev, paidOrders, rpu.String())
}
