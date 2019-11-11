package main
import (
	"fmt"
)

type person struct {
	name string
	age int
}
// tham so la name, co kieu string
// ket qua tra ve cua ham la mot con tro person.
func newPerson(_name string) *person {
	var p = person {name: _name}
	p.age = 32

	return &p;
}

func main(){
	fmt.Println(person{"teo", 12})
    fmt.Println(person{name: "Fred"})
    fmt.Println(person{age: 34})

    fmt.Println(&person{"teo2", 34})

    fmt.Println(newPerson("ti"))

    var ss = person{name: "sue", age: 12}
    fmt.Println(ss)

    var sn = &ss
    sn.age = 111
    fmt.Println(sn.age)
}