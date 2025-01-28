package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		smallest := i
		flag := false
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[smallest] {
				smallest = j
				flag = true
			}
		}
		if flag {
			temp := a[i]
			a[i] = a[smallest]
			a[smallest] = temp
		}
	}
}
