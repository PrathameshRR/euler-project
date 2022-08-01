//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

var smallest_num = 2520

func main() {

	for smallest_num < 1000000000 {
		if smallest_num%1 == 0 && smallest_num%2 == 0 && smallest_num%3 == 0 && smallest_num%4 == 0 && smallest_num%5 == 0 && smallest_num%6 == 0 && smallest_num%7 == 0 && smallest_num%8 == 0 && smallest_num%9 == 0 && smallest_num%10 == 0 && smallest_num%11 == 0 && smallest_num%12 == 0 && smallest_num%13 == 0 && smallest_num%14 == 0 && smallest_num%15 == 0 && smallest_num%16 == 0 && smallest_num%17 == 0 && smallest_num%18 == 0 && smallest_num%19 == 0 && smallest_num%20 == 0 {
			print(smallest_num)
			break
		}
		smallest_num++
	}

	print("End of Execusion")
}
