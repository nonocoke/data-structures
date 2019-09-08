package p0078

func Subsets(nums []int) [][]int {
	size := len(nums)
	res := [][]int{nil}
	if size == 0 {
		return res
	}

	for i := 0; i < size; i++ {
		for _, tmp := range res {
			res = append(res, append(tmp, nums[i]))  // append 失败
		}
		//fmt.Println(i, res)
	}

	return res
}
