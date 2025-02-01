package main

import "fmt"

func main() {
	var x1, y1, x2, y2 uint

	fmt.Println("введи 4 натуральных числа")

	fmt.Scan(&x1, &y1, &x2, &y2)

	if x1 > 8 || x1 < 1 || y1 > 8 || y1 < 1 || x2 > 8 || x2 < 1 || y2 > 8 || y2 < 1 {
		fmt.Println("НЕТ")
	} else {
		if x1+1 == x2 {
			if y1+2 == y2 || y1-2 == y2 {
				fmt.Println("ДА")
				return
			}
		}

		if x1+2 == x2 {
			if y1-1 == y2 || y1+1 == y2 {
				fmt.Println("ДА")
				return
			}
		}

		if x1-1 == x2 {
			if y1-2 == y2 || y1+2 == y2 {
				fmt.Println("ДА")
				return
			}
		}

		if x1-2 == x2 {
			if y1-1 == y2 || y1+1 == y2 {
				fmt.Println("ДА")
				return
			}
		}

		fmt.Println("НЕТ")
	}

}
