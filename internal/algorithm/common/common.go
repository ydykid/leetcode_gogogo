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
