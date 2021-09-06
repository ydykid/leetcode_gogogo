package week20210904

import "fmt"

func numberOfGoodSubsets(nums []int) int {
	p := 1000000007
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	ss := make([]int, 1024)
	for _, x := range nums {
		u, e := 0, 1
		for _, pi := range ps {
			if x%pi == 0 {
				u |= e
				x /= pi
			}
			if x%pi == 0 {
				u = 0
				break
			}
			e <<= 1
		}
		if u != 0 {
			ss[u]++
		}
	}
	// dp := make([][]int, 1024)
	// for i := 0; i < 1024; i++ {
	// 	dp[i] = make([]int, 1024)
	// }
	// dp[0][0] = 1
	// n := 0
	dp := make([]int, 1024)
	dp[0] = 1
	for i := 0; i < 1024; i++ {
		if ss[i] > 0 {
			// n++
			// for j := 0; j < 1024; j++ {
			// 	dp[n][j] = dp[n-1][j]
			// }
			for j := 0; j < 1024; j++ {
				if i&j == 0 {
					// dp[n][i|j] = (dp[n][i|j] + ss[i]*dp[n-1][j]%p) % p
					dp[i|j] = (dp[i|j] + ss[i]*dp[j]%p) % p
				}
			}
		}
	}
	c := 1
	for _, pi := range nums {
		if pi == 1 {
			c = (c << 1) % p
		}
	}
	ans := 0
	for i := 1; i < 1024; i++ {
		// ans = (ans + dp[n][i]) % p
		ans = (ans + dp[i]) % p
	}
	return (ans * c) % p
	// p := 1000000007
	// ps := []int{2,3,5,7,11,13,17,19,23,29}
	// pn := 10
	// n := len(nums)
	// ss := make([]int, n)
	// for i,x := range nums{
	//     u, e := 0, 1
	//     for _,p := range ps{
	//         if x%p==0{
	//             u |= e
	//             x /= p
	//         }
	//         if x%p==0{
	//             u = 0
	//             break
	//         }
	//         // for x%p==0{
	//         //     if u&e!=0{
	//         //         u = 0
	//         //         break
	//         //     }else{
	//         //         u |= e
	//         //     }
	//         //     x/=p
	//         // }
	//         // if x%p==0{
	//         //     if u==0{
	//         //         break
	//         //     }
	//         // }
	//         e <<= 1
	//     }
	//     ss[i] = u
	//     // if nums[i]==1{ss[i]=1<<10}
	// }
	// // fmt.Println(ss)
	// // dp := make([][]int, n)
	// // for i:=0; i<n; i++{dp[i]=make([]int, 1<<pn)}
	// dp := make([]int, 1<<pn)
	// if ss[0]>0{
	//     // dp[0][ss[0]] = 1
	//     dp[ss[0]] = 1
	// }
	// for i:=1; i<n; i++{
	//     if ss[i]>0{
	//         // dp[i][ss[i]] = 1
	//         dp[ss[i]] = (dp[ss[i]]+1)%p
	//     }
	//     for x:=0; x<1<<pn; x++{
	//         // if ss[i]!=0&&x&ss[i]==0&&dp[i-1][x]>0{
	//         if ss[i]!=0&&x&ss[i]==0&&dp[x]>0{
	//             // dp[i][x|ss[i]] = (dp[i][x|ss[i]]+dp[i-1][x])%p
	//             y := x|ss[i]
	//             dp[y] = (dp[y]+dp[x])%p
	//         }
	//         // dp[i][x] = (dp[i][x]+dp[i-1][x])%p
	//     }
	// }
	// // for i:=0; i<n; i++{
	// //     fmt.Println(dp[i][:10])
	// // }
	// ans := 0
	// for i:=0; i<1<<pn; i++{
	//     // ans = (ans+dp[n-1][i])%p
	//     ans = (ans+dp[i])%p
	// }
	// c := 1
	// for _,x := range nums{
	//     if x==1{
	//         c = (c<<1)%p
	//     }
	// }
	// // c = (1+c)*c/2+1
	// return (c*ans)%p
}

func Test() {
	qs := [][]int{
		{1, 2, 3, 4},
		{4, 2, 3, 15},
	}
	for qi, q := range qs {
		fmt.Printf("#%d:", qi)
		fmt.Println(q)
		ans := numberOfGoodSubsets(q)
		fmt.Printf("ans: %d\n", ans)
	}

}
