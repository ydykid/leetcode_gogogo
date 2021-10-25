package trie

import (
	"fmt"
)

type LowCaseTrie struct{
	data[26] *LowCaseTrie
	end bool
}

func Constructor()LowCaseTrie{
	t := LowCaseTrie{}
	return t
}

func (t *LowCaseTrie)Add(w string){
	p := t
	for i:=0; i<len(w); i++{
		x := w[i]-'a'
		if p.data[x]==nil{p.data[x]=&LowCaseTrie{}}
		p = p.data[x]
	}
	p.end = true
}

func (t *LowCaseTrie)Has(w string) bool{
	p := t
	for i:=0; i<len(w); i++{
		x := w[i]-'a'
		if p.data[x]==nil{return false}
		p = p.data[x]
	}
	return true
}

func TestTrie(){
	fmt.Println("Trie test...")
	t := Constructor()
	t.Add("abc")
	t.Add("ydy")
	fmt.Println("abc", t.Has("abc"))
	fmt.Println("ydy", t.Has("ydy"))
	fmt.Println("ab", t.Has("ab"))
	fmt.Println("bc", t.Has("bc"))
	fmt.Println("a", t.Has("a"))
	fmt.Println("dy", t.Has("dy"))
	fmt.Println("y", t.Has("y"))
}

