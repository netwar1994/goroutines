package card

type Part struct {
	MonthTimestamp int64
	Transactions []*Transaction
}

type Transaction struct {
	//Id  int64
	Sum int64
}

func Sum(transactions []*Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions{
		result += transaction.Sum
	}
	return result
}


