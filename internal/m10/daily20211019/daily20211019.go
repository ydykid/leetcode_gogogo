package daily20211019

import (
    "fmt"
)

type WordDictionary struct {
    data [26]*WordDictionary
    end bool
}


func Constructor() WordDictionary {
    wd := WordDictionary{}
    return wd
}


func (this *WordDictionary) AddWord(word string)  {
    for p,i:=this,0; i<len(word); i++{
        x := word[i]-'a'
        if p.data[x]==nil{
            p.data[x] = &WordDictionary{}
        }
        if i==len(word)-1{
            p.end = true
        }
    }
}


func (this *WordDictionary) Search(word string) bool {
    if len(word)<=0{return this.end}
    if word[0]=='.'{
        for i:=0; i<26; i++{
            if this.data[i]!=nil&&this.data[i].Search(word[1:]){
                return true
            }
        }
        return false
    }
    x := word[0]-'a'
    if this.data[x]==nil{return false}
    return this.data[x].Search(word[1:])
}

func (wd *WordDictionary) Display(dep int){
    for x:=0; x<26; x++{
        for i:=0; i<dep; i++{fmt.Print(" ")}
        if wd.data[x]!=nil{
            fmt.Print(string('a'+x))
            if wd.data[x].end{fmt.Print("$")}
            fmt.Println()
            wd.data[x].Display(dep+1)
        }
    }
    

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

 func Test(){
     fmt.Println("qs1:")
     wd := WordDictionary{}
     wd.AddWord("bad")
     wd.AddWord("dad")
     wd.AddWord("mad")
     wd.Display(0)
 }
