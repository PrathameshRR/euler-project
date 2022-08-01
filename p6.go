//The sum of the squares of the first ten natural numbers is,

//The square of the sum of the first ten natural numbers is,

//Hence the difference between the sum of the squares of the first ten natural numbers 
//and the square of the sum is .

//Find the difference between the sum of the squares of the first one hundred natural numbers 
//and the square of the sum.

package main

var sum_of_squares int = 0
var sqaure_of_sum int = 0
var difference int = 0
func main() {

	for i := 1; i <= 100; i++ {

		sum_of_squares += i * i
	}

	for i := 1; i <= 100; i++ {

		sqaure_of_sum += i
	}
	sqaure_of_sum = sqaure_of_sum * sqaure_of_sum

	//print(sum_of_squares)
	//print(\nsqaure_of_sum)
	difference = sqaure_of_sum - sum_of_squares
	print(difference)
}