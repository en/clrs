package clrs

func iterativeBinarySearch(a []int, v, low, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if v > a[mid] {
			low = mid + 1
		} else if v < a[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func recursiveBinarySearch(a []int, v, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if v > a[mid] {
		return recursiveBinarySearch(a, v, mid+1, high)
	} else if v < a[mid] {
		return recursiveBinarySearch(a, v, low, mid-1)
	} else {
		return mid
	}
}
