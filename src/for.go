package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("---------")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("---------")
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	for n := 0; n < 5; n++ {
		if n%5 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
