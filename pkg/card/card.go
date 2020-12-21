package card

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Part struct {
	MonthTimestamp int64
	Transactions []*Transaction
}

type Transaction struct {
	//Id  int64
	Sum int64
}

func SortTransactions(transactions []Transaction) []Transaction {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Sum > transactions[j].Sum
	})
	return transactions
}

func Sum(transactions []*Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions{
		result += transaction.Sum
	}
	return result
}

func SumConcurrently(transactions []*Transaction, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		// ВАЖНО: просто с partSize не прокатит, вам нужно как-то заранее разделить слайс по месяцам
		// ВАЖНО: этот код - лишь шаблон, который показывает вам как запустить горутину
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			sum := Sum(part)
			fmt.Println("sum:",  sum)
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}

func SumByMonth(period *Part)  {
	fmt.Println(period.MonthTimestamp)
	t := time.Unix(period.MonthTimestamp, 0)
	fmt.Println(t.Month())
}
