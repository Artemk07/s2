package main

import (
	"fmt"
)

// func plus(massiv []int) {
// 	fmt.Println("q", massiv)

// 	fmt.Println("q", massiv)
// 	massiv = append(massiv, 0) //добовляем элемент к срезу
// 	fmt.Println("q", massiv)
// 	qw := len(massiv) //переменная равная длине среза
// 	fmt.Println("q", massiv)
// 	for i := 0; i < qw; i++ {
// 		if massiv[i] >= 9 {
// 			massiv[i] = 0
// 			massiv[i+1]++
// 			fmt.Println("massiv", massiv)
// 		}

// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(massiv)))
// 	fmt.Println("massiv", massiv)
// }

func plus(massiv []int) {
	massiv = append(massiv, 0)
	copy(massiv[1:], massiv)
	massiv[0] = 0
	qw := len(massiv) //переменная равная длине среза
	for i := qw - 1; i >= 0; i-- {
		if massiv[i] >= 9 {
			massiv[i] = 0
			massiv[i-1]++
		}

	}
	if massiv[0] == 0 {
		massiv = append(massiv[:0], massiv[1:]...)
	}
	fmt.Println("massiv", massiv)
}
func main() {
	ww := []int{1, 2, 5, 8, 0, 9}
	plus(ww)
	// var s = []int{2, 3, 4}
	// fmt.Println("12", s)
	// s = append(s, 0)
	// copy(s[1:], s)
	// s[0] = 1
	// fmt.Println("12", s)
}
