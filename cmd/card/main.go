package main

import (
	"fmt"
	"github.com/netwar1994/goroutines/pkg/card"
	"math/rand"
	"strconv"
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

	monthsMap := make(map[string][]*card.Transaction)

	for m := range months{
		t :=  time.Unix(months[m].MonthTimestamp, 0)
		year := t.Year()
		month := t.Month().String()
		d := strconv.Itoa(year) + "-" + month

		monthsMap[d] = months[m].Transactions
	}

	fmt.Println("Transactions sum by month")
	for m, t := range monthsMap {
		fmt.Print(m, ":")
		card.SumConcurrently(t, 1)
	}

}
