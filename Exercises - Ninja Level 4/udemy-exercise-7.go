package main

import "fmt"

func main() {

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{x1, x2}
	fmt.Println(x)
	for index, value := range x {
		fmt.Printf("Records %v\n", index)
		for j, value := range value {
			fmt.Printf("Details of record %v, index %v, %v \n", index, j, value)
		}
	}

}
