package main

import (
	"fmt"
	"github.com/netwar1994/goroutines/pkg/card"
	"math/rand"
	"time"
)

func main() {
	start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	finish := time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)

	months := make([]*card.Part, 0)

	next := start
	for next.Before(finish) {
		months = append(months, &card.Part{
			MonthTimestamp: next.Unix(),
		})
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, &card.Part{
		MonthTimestamp: finish.Unix(),
	})


	transactions := make([]int64, 10_00)
	for transaction := range transactions{
		transactions[transaction] = rand.Int63n(100)
	}


	for _, transaction := range transactions {
		month := months[rand.Intn(len(months))]
		month.Transactions = append(month.Transactions, &card.Transaction{Sum: transaction})
	}

	for m := range months{
		d := time.Unix(months[m].MonthTimestamp, 0)
		fmt.Println(d.Year(), d.Month(), card.Sum(months[m].Transactions))
	}
}
