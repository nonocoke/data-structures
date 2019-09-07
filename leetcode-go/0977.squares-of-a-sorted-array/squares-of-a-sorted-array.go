package p0977

func SortedSquares(A []int) []int {
	// A: 非递减顺序排序的整数数组，且 1 <= A.length <= 10000
	length := len(A)
	squares := make([]int, length)
	head, tail, i := 0, length - 1, length - 1
	// 临时存储首、尾元素值
	headValue, tailValue := A[head]*A[head], A[tail]*A[tail]
	for head < tail {
		if headValue <= tailValue {
			squares[i] = tailValue
			i --; tail --
			// 替换临时尾部元素值
			tailValue = A[tail]*A[tail]
		} else {
			squares[i] = headValue
			i--; head++
			// 替换临时首部元素值
			headValue = A[head] * A[head]
		}
	}
	// 添加遗漏的最后一个元素
	if head == tail {
		squares[i] = headValue
	}
	return squares
}