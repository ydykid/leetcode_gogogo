package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quickSort(data []int, l, r int) {
	if l >= r {
		return
	}
	i, j, h, x := l, r, (l+r)/2, data[(l+r)/2]
	for data[h] = data[i]; i < j; {
		for i < j && data[j] > x {
			j--
		}
		if i < j {
			data[i], i = data[j], i+1
		}
		for i < j && data[i] <= x {
			i++
		}
		if i < j {
			data[j], j = data[i], j-1
		}
	}
	data[i] = x
	quickSort(data, l, i-1)
	quickSort(data, i+1, r)
}

func heapDown(data []int, i, n int) {
	if n >= len(data) {
		return
	}
	for x, y := i, i*2; y <= n && data[y] > data[x] || y < n && data[y+1] > data[x]; {
		if y+1 <= n && data[y+1] > data[y] {
			y++
		}
		if data[y] > data[x] {
			data[x], data[y], x = data[y], data[x], y
		}
		y = x * 2
	}
}

func heapSort(data []int) {
	n := len(data)
	for i := n / 2; i >= 0; i-- {
		heapDown(data, i, n-1)
	}
	for i := n - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapDown(data, 0, i-1)
	}
}

func verifySort(data []int) int {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return i
		}
	}
	return -1
}

func main() {
	n := 1000000
	data := make([]int, n)
	data2 := make([]int, n)
	data3 := make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n * 10)
		data2[i] = data[i]
		data3[i] = data[i]
	}
	//fmt.Println(data)
	fmt.Println("1.sort.Ints:")
	t1 := time.Now()
	sort.Ints(data)
	t2 := time.Now()
	u := verifySort(data)
	if u != -1 {
		fmt.Printf("Error:%d:%d->%d:%d", u-1, data[u-1], u, data[u])
	}
	fmt.Println("time:", t2.Sub(t1))

	fmt.Println("2.quickSort:")
	t1 = time.Now()
	quickSort(data2, 0, len(data2)-1)
	t2 = time.Now()
	u = verifySort(data2)
	if u != -1 {
		fmt.Printf("Error:%d:%d->%d:%d", u-1, data2[u-1], u, data2[u])
	}
	fmt.Println("time:", t2.Sub(t1))

	fmt.Println("3.heapSort:")
	t1 = time.Now()
	heapSort(data3)
	t2 = time.Now()
	u = verifySort(data3)
	if u != -1 {
		fmt.Printf("Error:%d:%d->%d:%d", u-1, data3[u-1], u, data3[u])
	}
	fmt.Println("time:", t2.Sub(t1))

}
