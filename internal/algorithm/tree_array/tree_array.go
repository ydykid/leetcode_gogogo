
type TreeArray struct{
	data []int
	n int
}

func Constructor(n int)TreeArray{
	t := TreeArray{}
	t.n = n
	t.data = make([]int, n+1)
}

func lb(x int){return x&(-x)}

func (t *TreeArray)Add(x int){
	for x<=t.n{
		t.data[x]++
		// x += lb(x)
		x = x|(x+1)
	}
}

func (t *TreeArray)Q(x int)(ret int){
	for x<=t.n&&x>=0{
		ret += t.data[x]
		// x -= lb(x)
		x = (x&(x+1)-1)
	}
}