package common

func min(x, y int)int {
	if x<y {return x}
	return y
}

func max(x, y int)int {
	if x>y {return x}
	return y
}

// Union Set
type MySet struct{data []int}
func (s *MySet) find(x int)int{
	if x>=len(s.data){return -1}
	if s.data[x]==0||s.data[x]==x{return x}
	s.data[x] = s.find(s.data[x])
	return s.data[x]
}
func (s *MySet) merge(x, y int){
	if x>=len(s.data){return}
	xx, yy := s.find(x), s.find(y)
	s.data[xx] = yy
}

// Heap int

type MyHeapInt struct {
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