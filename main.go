package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = []int{2, 10, 1, -1, 1, -2}
	temp := findTreasure(a)
	fmt.Println("result found ", temp)
}

func findTreasure(arr []int) []int {
	index := 0
	length := len(arr)
	var indexArr = []int{}
	for { // you can omit the condition ~ while (true)
		if index > length {
			fmt.Println("treasure found")
			return indexArr
		}
		if index < length {
			if arr[index] == 0 {
				fmt.Println("Zombie found")
				return []int{0}
			}
			indexArr = append(indexArr, arr[index])
			if checkInfiniteLoop(indexArr) {
				fmt.Println("infinite path found")
				return []int{-1}
			}
			index = index + arr[index]

		} else if index == length {
			fmt.Println("treasure found")
			return indexArr
		} else {
			fmt.Println("Invalid array for treasure")
			break
		}

	}

	var a = []int{0}
	return a
}

func checkInfiniteLoop(arr []int) bool {
	length := len(arr)

	if length == 2 {
		if arr[0]+arr[1] == 0 {
			return true
		} else {
			return false
		}
	}
	if length%2 == 0 {
		sliceArray1 := arr[0 : length/2]
		sliceArray2 := arr[length/2 : length]
		fmt.Println(reflect.DeepEqual(sliceArray1, sliceArray2))
		return reflect.DeepEqual(sliceArray1, sliceArray2)
	}
	if length > 2 {
		if arr[length-1]+arr[length-2] == 0 {
			return true
		}
	}
	return false
}
