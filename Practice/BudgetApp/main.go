package main

import "fmt"

// Create a simple budget calculator. Add bills, add income, subtract bills from income,
// recommend savings allocation based on percentage
// Work in progress. Everyday coding practice.
// Practice using interfaces.

//type bills to collect bills
type bills []float64

//Method for Bills to add them together and return total
func (b bills) sumTotal() float64 {
	var sum float64
	fmt.Println(len(b))
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum
}

//type income to collect income sources
type income []float64

//Method to add income sources
func (i income) sumTotal() float64 {
	var sum float64
	for j := 0; j < len(i); j++ {
		sum += i[j]
	}
	return sum
}

//Interface to implement addition for each type
type total interface {
	sumTotal() float64
}

//method for type total to implement interface
func added(t total) float64 {
	return t.sumTotal()
}

func main() {
	bill1 := bills{100.00, 200.00, 100.00}
	income1 := income{1100.00, 900.00}
	totalIncome := added(income1)
	totalBills := added(bill1)
	savingsRec := (totalIncome - totalBills) * 0.2
	fmt.Println(" You have ", totalIncome, "total income \n", totalBills, "in bills \n", "Therefore, you should send ", savingsRec, "to savings")
}
