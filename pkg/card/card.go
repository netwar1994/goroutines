package card

import (
	"sort"
)

type Transaction struct {
	Id  int64
	Sum int64
}

func SortTransactions(transactions []Transaction) []Transaction {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Sum > transactions[j].Sum
	})
	return transactions
}
