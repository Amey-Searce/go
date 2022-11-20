package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	// print(x[`bond_james`][0])
	for index, value := range x {
		fmt.Println(index)
		for j, value1 := range value {
			fmt.Println(j, value1)
		}
	}
}
