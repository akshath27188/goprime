package main

import (
	"reflect"
	"testing"
)

func TestLoadArrayWithPrimeNumbers(t *testing.T) {
	var arr_exp = make([][]int, 10)

	arr_exp[0] = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 31}
	arr_exp[1] = []int{3, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[2] = []int{5, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[3] = []int{7, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[4] = []int{11, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[5] = []int{13, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[6] = []int{17, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[7] = []int{19, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[8] = []int{23, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arr_exp[9] = []int{31, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var arr [][]int = make([][]int, 10)
	for i := range arr {
		arr[i] = make([]int, 10)
	}
	arr[0][0] = 2
	arr[0][1] = 3

	arr[1][0] = 3
	arr_res := LoadArrayWithPrimeNumbers(arr, 10)
	for i := range arr_res {
		if !reflect.DeepEqual(arr_res[i], arr_exp[i]) {
			t.Error("Unexpected prime values in array")
			break
		}
	}

}
