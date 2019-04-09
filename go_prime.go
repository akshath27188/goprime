package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello, playground")
	limit := 10
	var arr [][]int = make([][]int, limit)
	for i := range arr {
		arr[i] = make([]int, limit)
	}
	arr[0][0] = 2
	arr[0][1] = 3

	arr[1][0] = 3
	arr = LoadArrayWithPrimeNumbers(arr, limit)
	arr = LoadArrayWithMultiVal(arr, limit)
	PrintArr(arr, limit)
}

//Prime numbers can be represented as 6n+1 and 6n-1 except 2 and 3
//with possibility of 6n+1 and 6n-1 not being a prime number
//incase not a prime it is divisible by either 5 or 7
//Complexity - O(n)
func LoadArrayWithPrimeNumbers(arr [][]int, limit int) (arra [][]int) {
	str := "negative"
	//load with prime numbers
	counter := 1
	var num int
	count := 1
	for counter <= limit-2 {
		fmt.Println("index to formula-", count)
		if str == "negative" {
			num = 6*count - 1
			str = "positive"
			if math.Mod(float64(num), 5) != 0 && math.Mod(float64(num), 7) != 0 || num == 5 || num == 7 {
				counter++
				arr[0][counter] = num
				arr[counter][0] = num

				continue
			}
		}
		if str == "positive" {
			num = 6*count + 1
			count = count + 1
			str = "negative"
			if math.Mod(float64(num), 5) != 0 && math.Mod(float64(num), 7) != 0 || num == 5 || num == 7 {
				counter++

				arr[0][counter] = num
				arr[counter][0] = num
				continue
			}
		}
		fmt.Println("number-", num)
	}
	return arr
}

//load array with multiplication values
//Complexity:O(n^2) -- O(20)
func LoadArrayWithMultiVal(arr [][]int, limit int) (arra [][]int) {
	for i := 1; i < limit; i++ {
		for j := 1; j < limit; j++ {
			arr[i][j] = arr[0][j] * arr[i][0]
		}
	}
	return arr
}

//print in table format
//Complexity:O(n^2) -- O(20)
func PrintArr(arr [][]int, limit int) {
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Print(arr[i][j], " ")
			if j == limit-1 {
				fmt.Println("")
			}
		}
	}
}
