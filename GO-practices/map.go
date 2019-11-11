package main
import (
	"fmt"
)

func main(){
	var m = make(map[string] int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	var v1 = m["k1"]
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m["k2"])
}