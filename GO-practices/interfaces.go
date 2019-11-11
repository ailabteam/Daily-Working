package main
import (
	"fmt"
	"math"
)

type geomery interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geomery){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	var r = rect{3, 4}
	var c = circle{5}

	measure(r)
	measure(c)

}