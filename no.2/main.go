package main

import "fmt"

func main() {
	var n int

	fmt.Print("jumlah elemen dalam array: ")
	fmt.Scan(&n)

	var arr1 = make([]string, n)
	fmt.Println("array 1:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	var arr2 = make([]string, n)
	fmt.Println("array 2: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr2[i])
	}

	isSame := true
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			isSame = false
			break
		}
	}

	if isSame {
		fmt.Println("kedua array sama")
	} else {
		diff := []int{}
		for i, v1 := range arr1 {
			v2 := arr2[i]
			if v1 != v2 {
				diff = append(diff, i)
			}
		}
		for _, index := range diff {
			fmt.Printf("index ke %d berbeda\n", index)
		}
	}
}
