package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello, playground")

	var arr [][]int = make([][]int, 10)
	for i := range arr {
		arr[i] = make([]int, 10)
	}
	arr[0][0] = 2
	arr[0][1] = 3

	arr[1][0] = 3
	arr = LoadArrayWithPrimeNumbers(arr, 10)
	arr = LoadArrayWithMultiVal(arr)
	PrintArr(arr)
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
		//fmt.Println("index to formula-",count)
		if str == "negative" {
			num = 6*count - 1
			if math.Mod(float64(num), 5) != 0 && math.Mod(float64(num), 7) != 0 || num == 5 || num == 7 {
				counter++
				arr[0][counter] = num
				arr[counter][0] = num
				str = "positive"

				continue
			}
		}
		if str == "positive" {
			num = 6*count + 1
			count = count + 1
			if math.Mod(float64(num), 5) != 0 && math.Mod(float64(num), 7) != 0 || num == 5 || num == 7 {
				counter++
				str = "negative"
				arr[0][counter] = num
				arr[counter][0] = num
				continue
			}
		}

	}
	return arr
}

//load array with multiplication values
//Complexity:O(n^2) -- O(20)
func LoadArrayWithMultiVal(arr [][]int) (arra [][]int) {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			arr[i][j] = arr[0][j] * arr[i][0]
		}
	}
	return arr
}

//print in table format
//Complexity:O(n^2) -- O(20)
func PrintArr(arr [][]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(arr[i][j], " ")
			if j == 9 {
				fmt.Println("")
			}
		}
	}
}
