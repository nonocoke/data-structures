# 78. Subsets

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note:

```text
The solution set must not contain duplicate subsets.
```

Example:

```text
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```

解题思路：

```text
1. 输入为一个数组（内部元素不重复），输出为数组所有（不重复）的子集
2. 遍历元素的同时，增加进当前已存在的子集中遍历元素的同时，增加进当前已存在的子集中
```

## 遇到问题记录：
问题代码版本：

```go
package p0078

import "fmt"

func Subsets(nums []int) [][]int {
    size := len(nums)
    res := [][]int{nil}
    if size == 0 {
        return res
    }

    for i := 0; i < size; i++ {
        for _, tmp := range res {
            res = append(res, append(tmp, nums[i]))
        }
        fmt.Println(i, res)  // 由初始的nil，增加数组元素，形成新的子集集合
    }
    return res
}
```

输入数组nums = [9, 0, 3, 5, 7]时， 得到的结果与期待不一致。
<br>期待结果：
<br>[[] [9] [0] [9 0] [3] [9 3] [0 3] [9 0 3] [5] [9 5] [0 5] [9 0 5] [3 5] [9 3 5] [0 3 5] [9 0 3 7] [7] [9 7] [0 7] [9 0 7] [3 7] [9 3 7] [0 3 7] [9 0 3 7] [5 7] [9 5 7] [0 5 7] [9 0 5 7] [3 5 7] [9 3 5 7] [0 3 5 7] [9 0 3 5 7]]

实际结果：
<br>[[] [9] [0] [9 0] [3] [9 3] [0 3] [9 0 3] [5] [9 5] [0 5] [9 0 5] [3 5] [9 3 5] [0 3 5] [9 0 3 7] [7] [9 7] [0 7] [9 0 7] [3 7] [9 3 7] [0 3 7] [9 0 3 7] [5 7] [9 5 7] [0 5 7] [9 0 5 7] [3 5 7] [9 3 5 7] [0 3 5 7] [9 0 3 ~~7~~ 7]]

打印的调试信息如下：
<br>0 [[] [9]]
<br>1 [[] [9] [0] [9 0]]
<br>2 [[] [9] [0] [9 0] [3] [9 3] [0 3] [9 0 3]]
<br>3 [[] [9] [0] [9 0] [3] [9 3] [0 3] [9 0 3] [5] [9 5] [0 5] [9 0 5] [3 5] [9 3 5] [0 3 5] [9 0 3 5]]
<br>4 [[] [9] [0] [9 0] [3] [9 3] [0 3] [9 0 3] [5] [9 5] [0 5] [9 0 5] [3 5] [9 3 5] [0 3 5] [9 0 3 ~~7~~] [7] [9 7] [0 7] [9 0 7] [3 7] [9 3 7] [0 3 7] [9 0 3 7] [5 7] [9 5 7] [0 5 7] [9 0 5 7] [3 5 7] [9 3 5 7] [0 3 5 7] [9 0 3 ~~7~~ 7]]


状态： 问题待解决...

日期： 2019-09-08