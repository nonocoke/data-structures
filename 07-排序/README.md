# 排序

依据时间复杂度分类：

| 排序算法 | 时间复杂度 | 是否基于比较 |
| ---------------- | ---------- | ------------ |
| 冒泡、插入、选择 | O(n^2)     | 是           |
| 快速、归并       | O(nlogn)   | 是           |
| 桶、计数、基数   | O(n)       | 否           |


1. [冒泡、插入、选择、希尔](#link1)
2. [快速、归并](#link2)
3. [桶、计数、基数](#link3)


<h2><span id="link1">1. 冒泡、插入、选择、希尔</span></h2>
#### 冒泡排序

冒泡排序只会操作相邻的两个数据。每次冒泡会对相邻的的两个元素进行比较，看是否满足大小关系要求，若不满足，就互换元素。一次冒泡至少让一个元素移动到它应该在的位置，移动 n 次就完成了排序。

* 冒泡排序是原地排序算法，空间复杂度为O(1)
* 稳定的排序算法
* 时间复杂度：最好O(n)、最差O(n^2)、平均O(n^2)

```go
// bubbleSort sort a to a increasing array.
func bubbleSort(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}
	for i := 0; i < size; i++ {
		flag := false  // 提前退出的标志位
		for j := 0; j < size - i - 1; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				flag = true  // 表示有数据交换
			}
		}
		// 若当前冒泡无数据交换，表示数组已排序完成，直接退出
		if !flag {
			break
		}
	}
}
```

#### 插入排序

插入排序通过将数组分为已排序区间和未排序区间的方式。初始的已排序区间中只有一个元素，就是数组的第一个元素。插入排序的核心思想就是取未排序区间的元素，在已排序区间找到合适的插入位置将其插入，并保证已排序区间元素一直有序。重复这个过程，知道未排序区间为空。

Tip: 常为基本有序、或数据规模较小情况下使用。

* 插入排序是原地排序算法，空间复杂度为O(1)
* 稳定的排序算法
* 时间复杂度：最好O(n)、最差O(n^2)、平均O(n^2)

```go
// InsertionSort sort a to a increasing array.
func InsertionSort(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}
	for i := 0; i < size; i++ {
	    // 当前要插入的元素为a[i]
		value := a[i]
		// 已排序区间大小为 i - 1
		j := i - 1
	    // 查找 a[i] 在已排序区间中适合的位置
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]  // 数据移动
			} else {
				break
			}
		}
		a[j+1] = value  // 插入数据
	}
}
```

#### 选择排序

选择排序类似插入排序。不同的是，每次选择未排序区间中的最小元素，插入到已排序区间的末尾。

* 插入排序是原地排序算法，空间复杂度为O(1)
* 不稳定的排序算法
* 时间复杂度：最好O(n^2)、最差O(n^2)、平均O(n^2)

```go
// SelectionSort sort a to a increasing array.
func SelectionSort(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}
	for i := 0; i < size; i++ {
		// 查找最小值
		minIdx := i
		for j := i + 1; j < size; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		// 交换
		a[i], a[minIdx] = a[minIdx],a[i]

	}
}
```

#### 希尔排序

希尔排序（shell sort）是插入排序的改进版本。希尔排序是把记录按下标的一定增量分组，对每组使用直接插入排序。核心思想：将数组按步长 step 进行分组，遍历 step<v<len(a)，对每一组进行排序，排序后得到的数据有序程度提高。然后缩短步长，重新分组，遍历 newstep<v<len(a)，对每一组进行排序。直至 newstep = 1。

* 希尔排序是原地排序算法，空间复杂度为O(1)
* 不稳定的排序算法
* 时间复杂度：最好-根据步长序列的不同而不同、最差O(n^2)、平均O(nlogn)

```go
// ShellSort sort a to a increasing array.
func ShellSort(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}
	// 步长 step ，选择 n/3
	for step := size/3; step > 0; step /= 3 {
		// // 遍历 步长 = step 后的所有数据，按 step 进行分组排序
		for i := step; i < size; i++ {
			// 按 step，数组从 j 到 0 进行交换，即 a[i] a[i-step] a[i-step-step] ... 进行比较
			for j:= i - step; j >= 0; j -= step {
				if a[i] < a[j] {
					a[i], a[j] = a[j], a[i]
				} else {
					break
				}
			}
		}
	}
}
```

<h2><span id="link12">2. 快速、归并</span></h2>
#### 快速排序

#### 归并排序

<h2><span id="link12">3. 桶、计数、基数</span></h2>