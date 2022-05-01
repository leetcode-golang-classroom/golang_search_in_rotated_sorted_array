package sol

func search(nums []int, target int) int {
	nLen := len(nums)
	if nLen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := nLen - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] { // left portion
			if nums[mid] < target || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // right portion
			if nums[mid] > target || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
