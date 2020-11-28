package sort

func QuickSort(array []int) []int {
	return quick1(array, 0, len(array)-1)
}

func quick1(array []int, low, high int) []int {
	if low < high {
		pivot := partition(array, low, high)
		quick1(array, low, pivot-1)
		quick1(array, pivot+1, high)
	}
	return array
}

func partition(array []int, low int, high int) int {
	pivot := array[low]
	for low < high {
		for low < high && pivot < array[high] {
			high--
		}
		array[low] = array[high]
		for low < high && pivot > array[low] {
			low++
		}
		array[high] = array[low]
	}

	array[low] = pivot
	return low
}
