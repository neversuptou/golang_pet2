package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		fmt.Println("Введите транзацкцию (любой символ для выхода)")
		transaction := scanTransactions()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

func scanTransactions() float64 {
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}
