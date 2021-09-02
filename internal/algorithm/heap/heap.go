package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func main() {
	qs := []int{
		3, 2, 6, 1, 3, 4,
	}
	que := hp{}
	que2 := hp{}
	fmt.Println(que)
	for i, qi := range qs {
		fmt.Println("#", i, ":")
		heap.Push(&que, qi)
		fmt.Println(que, que.IntSlice[0])
		heap.Push(&que2, -qi)
		fmt.Println(que2, que2.IntSlice[0])
	}
	fmt.Println("que:")
	for que.Len() > 0 {
		fmt.Print(heap.Pop(&que), " ")
	}
	fmt.Println("que2:")
	for que2.Len() > 0 {
		fmt.Print(heap.Pop(&que2), " ")
	}
}
