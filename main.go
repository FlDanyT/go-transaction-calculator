package main

import "fmt"

func main() {

	transactions := []float64{} // [кол-во значений в массиве] {значения через ,}

	for {

		transaction := scanTransaction()
		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
		
	}

	balance := calculateBalance(transactions)
	fmt.Println("Ваш баланс: ", balance)

}

func scanTransaction() float64 {

	var transaction float64

	fmt.Println("Введите транзакцию (n для выхода ): ")
	fmt.Scan(&transaction)

	return transaction

}

func calculateBalance(transactions []float64) float64{
	
	balance := 0.0 // Если 0 то это int

	for _, value := range transactions {

		balance += value // Добавляем значение value к balance

	}

	return balance
}
