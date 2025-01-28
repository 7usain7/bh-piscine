package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		smallest := i
		flag := false
		for j := i + 1; j < len(a); j++ {
			if f(a[j], a[smallest]) < 0 {
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
