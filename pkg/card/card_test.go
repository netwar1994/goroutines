package card

import (
	"reflect"
	"testing"
)

func TestSortTransactions(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	transaction1 := Transaction{
		Id:  0,
		Sum: 3000,
	}
	transaction2 := Transaction{
		Id:  1,
		Sum: 2000,
	}
	transaction3 := Transaction{
		Id:  2,
		Sum: 22000,
	}

	trans := []Transaction{transaction1, transaction2, transaction3}
	transSort := []Transaction{transaction3, transaction1, transaction2}

	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{name: "a",
			args: args{
				trans,
			},
			want: transSort,
		},
	}
	for _, tt := range tests {
		if got := SortTransactions(tt.args.transactions); !reflect.DeepEqual(tt.args.transactions, tt.want) {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}
