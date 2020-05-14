package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/polozon/GolangLearnApp/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Peter Pan"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
