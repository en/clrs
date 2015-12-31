package clrs

func bruteForceFindMaximumSubarray(a []int, low, high int) (int, int, int) {
	sum := 0
	valid := false // 标记sum为无效值
	left := low
	right := high
	for i := low; i <= high; i++ {
		s := 0
		for j := i; j <= high; j++ {
			s = s + a[j]
			if !valid || s > sum {
				valid = true
				left = i
				right = j
				sum = s
			}
		}
	}
	return left, right, sum
}

func findMaxCrossingSubarray(a []int, low, mid, high int) (int, int, int) {
	leftSum := 0
	valid := false
	sum := 0
	maxLeft := mid
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if !valid || sum > leftSum {
			valid = true
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := 0
	valid = false
	sum = 0
	maxRight := mid + 1
	for j := mid + 1; j <= high; j++ {
		sum = sum + a[j]
		if !valid || sum > rightSum {
			valid = true
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

func findMaximumSubarray(a []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, a[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaximumSubarray(a, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(a, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(a, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func linearTimeFindMaximumSubarray(a []int, low, high int) (int, int, int) {
	left := low
	right := low
	sum := 0
	valid := false
	tempStart := low
	tempSum := 0
	for i := low; i <= high; i++ {
		m := tempSum + a[i]
		n := a[i]
		var max int
		if n > m {
			max = n
			tempStart = i
		} else {
			max = m
		}
		tempSum = max
		if !valid || max > sum {
			valid = true
			left = tempStart
			right = i
			sum = max
		}
	}
	return left, right, sum
}
