package treap

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	v, h, c int
	l, r    *TreeNode
}

func makeTreeNode(num int) *TreeNode {
	td := &TreeNode{num, rand.Intn(1000000), 1, nil, nil}
	return td
}

type Treap struct {
	root *TreeNode
}

func (t *Treap) Insert(num int) {
	t.root = t.root.insert(makeTreeNode(num))
}

func (t *Treap) Has(num int) bool {
	return t.root.has(num)
}

func (t *Treap) Get(index int) int {
	return t.root.get(index)
}

func (td *TreeNode) get(index int) int {
	if td == nil {
		return -1
	}
	lc := 0
	if td.l != nil {
		lc = td.l.c
	}
	if lc == index {
		return td.v
	} else if lc > index {
		return td.l.get(index)
	} else {
		return td.r.get(index - lc - 1)
	}
}

func (td *TreeNode) has(num int) bool {
	if td == nil {
		return false
	}
	if td.v == num {
		return true
	} else if num < td.v {
		return td.l.has(num)
	} else {
		return td.r.has(num)
	}
}
func (td *TreeNode) insert(x *TreeNode) *TreeNode {
	if td == nil {
		return x
	}
	ret := td
	ret.c++
	if x.v > ret.v {
		ret.r = ret.r.insert(x)
		if ret.h < ret.r.h {
			ret = ret.right()
		}
	} else {
		ret.l = ret.l.insert(x)
		if ret.h < ret.l.h {
			ret = ret.left()
		}
	}
	return ret
}

func (td *TreeNode) left() *TreeNode {
	if td == nil || td.l == nil {
		return td
	}
	l := td.l
	td.l = l.r
	l.r = td
	td.c = td.l.count() + td.r.count() + 1
	l.c = l.l.count() + l.r.count() + 1
	return l
}

func (td *TreeNode) right() *TreeNode {
	if td == nil || td.r == nil {
		return td
	}
	r := td.r
	td.r = r.l
	r.l = td
	td.c = td.l.count() + td.r.count() + 1
	r.c = r.l.count() + r.r.count() + 1
	return r
}

func (td *TreeNode) count() int {
	if td == nil {
		return 0
	}
	return td.c
}

func (t *Treap) Display() {
	var display func(*TreeNode, int)
	display = func(t *TreeNode, dep int) {
		if t == nil {
			return
		}
		display(t.r, dep+1)
		for i := 0; i < dep; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("%d(%d,%d)\n", t.v, t.c, t.h)
		display(t.l, dep+1)
	}
	fmt.Println("display:")
	display(t.root, 0)
	fmt.Println("display.end.")
}

func Test() {
	var treap Treap
	treap.Display()
	for i, x := range []int{3, 5, 2, 3, 6, 1} {
		fmt.Printf("#%d Treap insert: %d\n", i, x)
		treap.Insert(x)
		treap.Display()
	}
	fmt.Printf("Get%d: %d\n", 4, treap.Get(4))
	fmt.Printf("Has %d: %t\n", 3, treap.Has(3))
	fmt.Printf("Has %d: %t\n", 7, treap.Has(7))
}
