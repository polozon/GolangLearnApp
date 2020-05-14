package main

import (
	"flag"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/polozon/GolangLearnApp/morestrings"
)

func main() {
	var ip = flag.Int("ext", 1234, "help message for flagname")
	flag.Parse()

	fmt.Printf("IP = %v\n", *ip)
	fmt.Println(morestrings.ReverseRunes("Peter Pan"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
