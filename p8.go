//The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.
// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product.
//What is the value of this product?

package main

import "strconv"

var digit_number int = 1289365

// 7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450
var largest_product int = 0

var adjacent_product int

// var adjacent_product_2 int

func product_of_4_digits(number int) int {

	var product int = 1
	//var digit int = len(strconv.Itoa(number))

	for i := 0; i < 4; i++ {

		product = product * int(number%10)
		number = number / 10
	}
	return product
}

// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product.
//What is the value of this product?
func main() {

	var digit_len int = len(strconv.Itoa(digit_number)) / 2

	for i := 0; i < digit_len; i++ {

		adjacent_product = product_of_4_digits(digit_number)

		//adjacent_product_2 = product_of_4_digits(digit_number)

		if adjacent_product > largest_product {
			largest_product = adjacent_product
		}

		digit_number = digit_number / 10

		// print("\n", adjacent_product)
		// print("\n", adjacent_product_2)

	}

	print("\n", largest_product)
}

/*

Let's do it for a 10 digit number.

1. First find the left most 4 digits
	Get the product of the 4 digits.

2. Get rid of the left most digit

3. First find the left most 4 digits
	Get the product of the 4 digits.

4. Compare the two products.
	If the first product is greater than the second product, then keep the first product.
	If the second product is greater than the first product, then keep the second product.

5. Get rid of the left most digit

6. First find the left most 4 digits
	Get the product of the 4 digits.

7. Compare the two products.
	If the first product is greater than the second product, then keep the first product.
	If the second product is greater than the first product, then keep the second product.

*/

/* var digits = []float64{}
var number float64 = 254
*/

/* func main() {

	for number > 1 {
		var i int = 0
		digits[i] = math.Mod(number, 10)
		print("\n", digits[i])

		number = number / 10
		i++
	}
}
*/

// Extract digits from a number
// func extract_digits(num int) []int {
// 	digits := []int{}
//	for num > 0 {
//		digits = append(digits, num%10)
//		num = num / 10
//	}
//	return digits
// }
