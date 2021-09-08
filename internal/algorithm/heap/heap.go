package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// My Heap Item
type MyHeapItem struct {
	v, p, i int
}

type MyHeap []*MyHeapItem

func (h MyHeap) Len()int{return len(h)}
func (h MyHeap) Less(i, j int)bool{
	return h[i].p > h[j].p
}
func (h MyHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
	h[i].i, h[j].i = i, j
}
func (h *MyHeap) Push(v interface{}){
	item := v.(*MyHeapItem)
	item.i = len(*h)
	*h = append(*h, item)
}
func (h *MyHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	item.i = -1
	*h = (*h)[0:len(*h)-1]
	return item
}
func (h MyHeap) Show() {
	for i:=0; i<len(h); i++{
		fmt.Println("i=", i, "v:", h[i].v, "p:", h[i].p, "i:", h[i].i)
	}
}

func TestMyHeap() {
	data := [][]int{
		{1, 3}, {2, 1}, {9, 2}, {5, 0},
	}
	hp := MyHeap{}
	// heap.Init(hp)
	for i,x := range data{
		heap.Push(&hp, &MyHeapItem{v:x[0], p:x[1], i:i})
	}
	fmt.Println(hp)
	hp.Show()
	for len(hp) > 0{
		fmt.Println(heap.Pop(&hp))
		hp.Show()
	}
}

// My Heap Int
type MyHeapInt struct {
	sort.IntSlice
}

func (h *MyHeapInt) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *MyHeapInt) Pop() interface{} {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func main() {
	qs := []int{
		3, 2, 6, 1, 3, 4,
	}
	que := MyHeapInt{}
	que2 := MyHeapInt{}
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
