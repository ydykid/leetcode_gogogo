package main

import (
	"fmt"
	"leetcode_gogogo/internal/algorithm/random"
	"leetcode_gogogo/internal/algorithm/treap"
	"math/rand"
)

func testTreap() {
	fmt.Println("test treap:")
	random.RandInit()
	tp := treap.Treap{}
	for i := 0; i < 100; i++ {
		tp.Insert(rand.Intn(10000))
	}
	tp.Display()
	fmt.Println("test treap end.")
}

func main() {
	fmt.Println("treap:")
	testTreap()
	treap.Test()
	fmt.Println("treap------")
	fmt.Println("random:")
	random.Test()
	fmt.Println("random------")
}
