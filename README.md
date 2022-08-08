# golang_counting_bits

Given an integer `n`, return *an array* `ans` *of length* `n + 1` *such that for each* `i` **(`0 <= i <= n`)*,* `ans[i]` *is the **number of*** `1`***'s** in the binary representation of* `i`.

## Examples

**Example 1:**

```
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

```

**Example 2:**

```
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

```

**Constraints:**

- `0 <= n <= 105`

**Follow up:**

- It is very easy to come up with a solution with a runtime of `O(n log n)`. Can you do it in linear time `O(n)` and possibly in a single pass?
- Can you do it without using any built-in function (i.e., like `__builtin_popcount` in C++)?

## 解析

給定一個非負數字 n, 

回傳一個長度為 n+1 陣列 ans

陣列 ans 每個數值 ans[i] 代表數值 i 轉換成 2元表示式中非零的 bit 個數

對於 每個數值 可以發現  只要數值 是 2某個次方項 代表該 數值所使用非零的bit 數值是 1 

而對於每個 ans[i] 假設存在某一個 k 使得 $2^k$ > ans[i] 

則有以下關係式 ans[i] = 1 + ans[i - $2^k$]

而 $2^k$ 可以透過把 0b1 不斷向左 bit shift 來得到

![](https://i.imgur.com/hiXpItX.png)

透過這個方式每次 只要 loop 0 到 n

就可以算完所有結果所以時間複雜度是 O(n)

## 程式碼
```go
package sol

func countBits(n int) []int {
	k := 0
	ans := []int{0}
	for pos := 1; pos <= n; pos++ {
		if pos == (0b1 << k) {
			ans = append(ans, 1)
			k++
		} else {
			ans = append(ans, 1+ans[pos-(0b1<<(k-1))])
		}
	}
	return ans
}

```
## 困難點

1. 需要去理解每個 bit 個數的關係

## Solve Point

- [x]  建立一個 k:= 0 用來紀錄目前最大 2 的次方數, result := []int{0}
- [x]  loop i := 0..n 每次做以下檢查
- [x]  if  i == (0b1 << k) 則 把 1 加入 result ,並且更新 k = k+1
- [x]  否則把 1 + result[i- (0b1 << (k-1))] 放入 result
- [x]  回傳 result