package main

import (
	"flag"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/polozon/GolangLearnApp/interfacetest"
	"github.com/polozon/GolangLearnApp/morestrings"
)

func measure(g interfacetest.Geometry) {
	fmt.Println(g)
	fmt.Println("Area  =", g.Area())
	fmt.Println("Perim =", g.Perim())
}

func doInterfaceTest() {
	r := interfacetest.GetIt(1)
	c := interfacetest.GetIt(2)

	fmt.Println("Measure geomery 1")
	measure(r)
	fmt.Println("Measure geomery 2")
	measure(c)
}

func main() {
	var ip = flag.Int("ext", 1234, "help message for flagname")
	flag.Parse()

	fmt.Printf("IP = %v\n", *ip)
	s := morestrings.GetStr()
	fmt.Println(morestrings.ReverseRunes(s))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	doInterfaceTest()
}
