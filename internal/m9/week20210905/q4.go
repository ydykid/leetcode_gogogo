package week20210905

import (
	"fmt"
	"sort"
)

func gcdSort(nums []int) bool {
	n := len(nums)
	mset := MySet{make([]int, 100000)}
	for _, x := range nums{
		u := x
		for i:=2; i<=x/i; i++{
			if u%i==0{mset.merge(x, i)}
			for u%i==0{u/=i}
		}
		if u>1{mset.merge(x, u)}
	}
	fmt.Println(mset.data[:11])
	for i,x := range mset.data[:11]{
		fmt.Printf("%d:%d ", i, x)
	}
	fmt.Println()
	data := make([]int, n)
	for i,x := range nums{data[i]=x}
	sort.Ints(data)
	for i:=0; i<n; i++{
		if nums[i]==data[i]{continue}
		if mset.find(nums[i])!=mset.find(data[i]){return false}
	}
	return true
}

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



func Test(){
	qs := [][]int{
		{7, 21, 3},
		{5, 2, 6, 2},
		{10, 5, 9, 3, 15},
		{10, 3, 9, 6, 8},
	}
	for qi, qes := range qs{
		fmt.Printf("#%d", qi)
		fmt.Println(qes)
		ans := gcdSort(qes)
		fmt.Println("ans:", ans)
	}
}