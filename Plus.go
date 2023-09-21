package main

import (
	"fmt"
	"strconv"
	"strings"
)
func numtoArr(num int) []string {
	num++
	slice := strings.Split(strconv.Itoa(num), "")
	fmt.Println("", slice)
	return slice
}

func main() {
}
