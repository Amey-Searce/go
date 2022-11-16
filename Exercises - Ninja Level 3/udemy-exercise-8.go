package main

import (
	"fmt"
)

func main() {

	switch {
	case true:
		{
			fmt.Printf("True case")
		}
	case false:
		{
			fmt.Printf("False case")
		}

	default:
		{
			fmt.Printf("Default")
		}

	}

}
