package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Channel")

	c1 := make(chan map[string]int)
	c2 := make(chan map[string]int)

	for _, i := range []int{1, 2, 3} {
		go func(iNum int) {
			c1 <- map[string]int{"test" + strconv.FormatInt(int64(iNum), 10): iNum}
		}(i)
	}
	for _, i := range []int{5, 6, 7} {
		go func(iNum int) {
			c2 <- map[string]int{"test" + strconv.FormatInt(int64(iNum), 10): iNum}
		}(i)
	}

	for index := 0; index < 6; index++ {
		select {
		case msg := <-c1:
			fmt.Printf("%v\n", msg)

		case msg2 := <-c2:
			fmt.Printf("%v\n", msg2)
		}
	}

	// x, ok = <-c1
	// fmt.Printf("%v , %v", x, ok)
	// select {
	// case x, ok := <-c1:
	// 	if ok {
	// 		fmt.Printf("Value %d was read.\n", x)
	// 	} else {
	// 		fmt.Println("Channel closed!")
	// 	}
	// default:
	// 	fmt.Println("No value ready, moving on.")
	// }

}
