package main

import (
	"fmt"
	"math/rand"

	//"myProject/gtool/smap"
)

func main() {
	fmt.Println(rand.Int31n(100))
	type Option struct {
		num int
	}
	var op *Option
	fmt.Println(op.num)

	type kenshi struct {
		int
	}

	var a *kenshi
	fmt.Printf("%+v", a)
}