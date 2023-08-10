package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanln(&str)
	b := []byte(str)
	// fmt.Println(b[0])
	length := len(b)
	ins := 0
	if length%2 != 0 {
		ins = 1
	}
	var evencount, oddcount byte
	flag := 0
	var total byte
	if length < 13 && length > 16 {
		flag = 1
	}
	if b[0] == '4' || b[0] == '5' || b[0] == '6' || (b[0] == '3' && b[1] == '7') {
		for i := length - 1; i >= 0; i-- {
			if b[i] < '0' && b[i] > '9' {
				flag = 1
			}
		}
	} else {
		flag = 1
	}
	if flag == 0 {
		for i := length - 1; i >= 0; i-- {
			if (i+ins)%2 == 0 {
				temp := (b[i] - '0') * 2
				for temp != 0 {
					rem := temp % 10
					evencount += rem
					temp /= 10
				}

			} else {
				oddcount += (b[i] - '0')

			}
		}
		total = evencount + oddcount
		if total%10 == 0 {
			fmt.Println("valid")
		} else {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Println("invalid")
	}
}
