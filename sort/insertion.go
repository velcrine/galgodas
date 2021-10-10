package sort

func InsertionSort(arr *[]int) {
	l := len(*arr)
	for i := 1; i < l; i++ {
		pointer := (*arr)[i]
		j := i - 1
		for ; i >= 0 && pointer < (*arr)[j]; j-- {
			(*arr)[j+1] = (*arr)[j]
		}
		(*arr)[j+1] = pointer
	}
}
