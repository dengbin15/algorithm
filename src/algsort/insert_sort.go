package main

func InsertSort(strings []int) {
	for i := 1; i < len(strings); i++ {
		// 待排元素小于有序序列最后一个元素时，向前插入
		if strings[i] < strings[i-1] {
			temp := strings[i]
			for k := i; k >= 0; k-- {
				if k > 0 && temp < strings[k-1] {
					strings[k] = strings[k-1]
				} else {
					strings[k] = temp
					break
				}
			}
		}
	}
}
