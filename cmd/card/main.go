package main

import (
	"fmt"
	"github.com/netwar1994/goroutines/pkg/card"
)

func main() {
	transaction1 := card.Transaction{
		Id:     0,
		Sum:    3000,
	}
	transaction2 := card.Transaction{
		Id:     1,
		Sum:    2000,
	}
	transaction3 := card.Transaction{
		Id:     2,
		Sum:    22000,
	}

	transactions := []card.Transaction{transaction1, transaction2, transaction3}
	fmt.Println("До сортировки:", transactions)
	fmt.Println("После сортировки:", card.SortTransactions(transactions))

}
