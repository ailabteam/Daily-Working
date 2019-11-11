package main

import  (
	"fmt"
	"math"
)

const PI = 3.14
const s string = "constant"

func main(){	
	// true, false
	fmt.Println("go" + "lang")
	fmt.Println(true && false)
	fmt.Println(s)
	var ss = PI / 4
	fmt.Println(math.Sin(ss))

	fmt.Println("For")
	for i := 1; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		if i % 5 == 0 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("if else")
	var a = 43
	var b = 4
	if a % b == 0 {
		fmt.Println("OK")
	} else if b < 5 {
		fmt.Println("NO")
	} else {
		fmt.Println("YYYY")
	}
}