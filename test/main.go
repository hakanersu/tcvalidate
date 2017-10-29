package main

import (
	"fmt"

	"github.com/hakanersu/tcvalidate"
)

func main() {

	fmt.Println(validatetc.Validate("17223038680"))
	fmt.Println(validatetc.Validate("10000000146"))
	fmt.Println(validatetc.Validate("100000001"))

}
