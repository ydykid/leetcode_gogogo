package week20210904

func numberOfGoodSubsets(nums []int) int {
    p := 1000000007
    ps := []int{2,3,5,7,11,13,17,19,23,29}
    pn := 11
    n := len(nums)
    ss := make([]int, n)
    for i,x := range nums{
        u, e := 0, 1
        for _,p := range ps{
            for x%p==0{
                if u&e!=0{
                    u = 0
                    break
                }else{
                    u |= e
                }
                x/=p
            }
            if x%p==0{
                if u==0{
                    break
                }
            }
            e <<= 1
        }
        ss[i] = u
        if nums[i]==1{ss[i]=1<<10}
    }
    fmt.Println(ss)
    dp := make([][]int, n)
    for i:=0; i<=n; i++{dp[i]=make([]int, 1<<pn)}
    dp[0][ss[0]] = 1
    for i:=0; i<n; i++{
        for x:=0; x<1<<pn; x++{
            if ss[i]!=0&&x&ss[i]==0&&dp[i][x]>0{
                dp[i+1][x|ss[i]] = (dp[i][x]+1)%p
            }
            dp[i+1][x] = (dp[i+1][x]+dp[i][x])%p
        }
    }
    for i:=0; i<=n; i++{
        fmt.Println(dp[i][:10])
    }
    ans := 0
    for i:=0; i<1<<pn; i++{
        ans = (ans+dp[n][i])%p
    }
    return ans
}