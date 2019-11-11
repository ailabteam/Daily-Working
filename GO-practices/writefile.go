package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){
	// cach tao va ghi file
	var d1 = []byte("hello\ngo\n")
	var err = ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	// cach tao file rong
	f, err := os.Create("dat2")
	check(err)

	defer f.Close()

	var d2 = []byte {115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)

	fmt.Println("wrote %d byte\n", n2)


	n3, err := f.WriteString("writes\n")
	fmt.Println("wrote %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferd\n")
	fmt.Println("wrote %d bytes \n", n4)

	w.Flush()
}