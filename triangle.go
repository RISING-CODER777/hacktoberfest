/*1
1 2
1 2 3
1 2 3 4
1 2 3 4 5*/


package main

import "fmt"

func main() {
	var rows = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
