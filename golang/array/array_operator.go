package main

import "fmt"

const max = 10

var arr [max]int
var len int

func insert(index int, enum int) {
	if index > max-1 {
		fmt.Println("index can't right array's maxgth")
		return
	}
	arr[index] = enum
	len++
}

func deleteEnum(enum int) {
	for i := 0; i < len; i++ {
		if arr[i] == enum {
			arr[i] = 0
			for j := i; j < max-1; j++ {
				arr[j] = arr[j+1]
			}
			len--
			break
		}
	}
}

func deleteIndexEnum(index int) {
	if index > len-1 {
		return
	}
	arr[index] = 0
	for i := index; i < max-1; i++ {
		arr[i] = arr[i+1]
	}
	len--
}

func update(index, enum int) {
	if index > len-1 {
		return
	}
	arr[index] = enum
}

func print() {
	for i := 0; i < len; i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	fmt.Println("array basic operator !")
	for i := 0; i < max; i++ {
		insert(i, i+1)
	}
	fmt.Println(arr)
	deleteEnum(10)
	fmt.Println(arr)
	deleteIndexEnum(8)
	deleteIndexEnum(0)
	fmt.Println(arr)
}
