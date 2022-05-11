# golang_search_in_rotated_sorted_array

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **possibly rotated** at an unknown pivot index `k` (`1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (**0-indexed**). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the possible rotation and an integer `target`, return *the index of* `target` *if it is in* `nums`*, or* `-1` *if it is not in* `nums`.

You must write an algorithm with `O(log n)` runtime complexity.

## Examples

**Example 1:**

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

```

**Example 3:**

```
Input: nums = [1], target = 0
Output: -1

```

**Constraints:**

- `1 <= nums.length <= 5000`
- $`-10^4$<= nums[i] <=$10^4$`
- All values of `nums` are **unique**.
- `nums` is an ascending array that is possibly rotated.
- $`-10^4$ <= target <= $10^4$`

## 解析

nums 是一個平移 過 k 位置的整數陣列，也就是假設原本陣列 nums[0] < nums[1] < ... <nums[n-1]

經過平移 k 位置會是  [nums[k], nums[k+1], num[k+2], ... nums[n-1], nums[0], nums[1], ... num[k-1]]

題目給定一個整數 target，要求實做出一個演算法在時間複雜度 O(logn) 判斷 target 是否在這陣列內

如果直接去做逐步察看時間複雜度會是 O(n)

要達到 O(logn) 必須使用 binary search 。 然而，平移過的陣列並不像原本陣列單純左右界平移

需要思考一下怎麼透過平移的特行來做有效的逼近

首先看一下下圖

![](https://i.imgur.com/SSHObxM.png)

假設 shifted k 位置後， 假設給 $L = 0, R = len(nums)$, $M = (L+R)/ 2$ 

可以透過 $nums[L] < nums[M]$ 是否成立來判斷 $nums[M]$ 是在左邊還是右邊

考慮假設以上圖 $nums[M]$ 在 shift 過後的左邊的狀況 要把 $L$  更新為 $M+1$ 的狀況有以下兩種

1. $target < nums[L]$

![](https://i.imgur.com/yBdwe9d.png)

1. $target > nums[M]$

![](https://i.imgur.com/WFuitu6.png)

其他當 $nums[M]$ 在左邊的狀況則需要 將 $R$  更新為 $M-1$

考慮假設以上圖 $nums[M]$ 在 shift 過後的右邊的狀況 要把 $R$  更新為 $M-1$ 的狀況有以下兩種

1. $target < nums[M]$

![](https://i.imgur.com/QasL81k.png)

1. $target < nums[R]$

![](https://i.imgur.com/xHSfNUW.png)

其他當 $nums[M]$ 在右邊的狀況則需要 將 $L$ 更新為 $M+1$

## 程式碼

```go
func search(nums []int, target int) int {
    if len(nums) == 1 {
        if nums[0] != target {
            return -1
        }
        return 0
    }
    left := 0
    right := len(nums) -1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[left] { // left portion
            if target > nums[mid] || target < nums[left] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else { // right portion
            if target < nums[mid] || target > nums[right] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        
    }
    return -1
}
```

## 困難點

1. 經過平移 k 位置雖然還是 保持著一個順序，但需要去理解如何做到有效平移來逼近目標值
2. 需要透過圖形才能比較好理解為何要做適當的平移，不夠直覺

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity