//The prime factors of 13195 are 5, 7, 13 and 29.

//What is the largest prime factor of the number 600851475143 ?

package main

var largest_prime_factor int = 0
var num int = 600851475143

// find factors of a number

func main() {
	find_factors(num)
}

func is_prime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func find_factors(num int) {
	for i := 2; i <= num; i++ {
		if num%i == 0 {

			if is_prime(i) {
				largest_prime_factor = i
				print("\n", largest_prime_factor)
			}
		}
	}
}
