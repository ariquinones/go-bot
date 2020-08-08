package main

import (
	"fmt"

	quotev2 "rsc.io/quote/v2"
	"rsc.io/quote/v3"
)

func main() {
	x := quote.GlassV3()
	v2 := quotev2.OptV2()
	fmt.Println(x)
	fmt.Println(v2)
}
