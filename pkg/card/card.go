package card

import (
	"sync"
)

type Part struct {
	MonthTimestamp int64
	Transactions   []*Transaction
}

type Transaction struct {
	Id  int64
	Sum int64
}

func Sum(transactions []*Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions {
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
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			sum := Sum(part)
			total += sum
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}
