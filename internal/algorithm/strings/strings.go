package strings

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println(strings.ContainsAny("abc", "ue"))
	fmt.Println(strings.ContainsAny("abc", "duc"))
	fmt.Println(strings.ContainsAny("a", "duca"))
}