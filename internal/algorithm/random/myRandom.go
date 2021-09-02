package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInit() {
	rand.Seed(time.Now().UnixNano())
}

func Test() {
	RandInit()
	for i := 0; i < 20; i++ {
		fmt.Println(i, rand.Intn(10))
	}
}
