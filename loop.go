package main

import "fmt"

func main3() {
	fmt.Println("Loop")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("****************************************************************")

	k := 1
	for ; k < 10; k++ {
		fmt.Println(k)
	}

	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {

	// }

	fmt.Println("****************************************************************")

	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("****************************************************************")

outer:
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
			break outer
		}
		fmt.Println()
	}

	// finger := 4
	// fmt.Printf("Finger %d is ", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 5:
	// 	fmt.Println("Pinky")
	// default: //default case
	// 	fmt.Println("incorrect finger number")
	// }

	finger := 7
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5, 6, 7:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	fmt.Println("****************************************************************")

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

	fmt.Println("****************************************************************")

	switch num1 := 75; {
	case num1 < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		if num1 < 80 {
			break
		}
		fallthrough
	case num1 < 20:
		fmt.Printf("%d is lesser than 200", num)
	}

	fmt.Println("****************************************************************")

}
