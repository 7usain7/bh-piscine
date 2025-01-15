package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		smallest := i
		flag := false
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				smallest = j
				flag = true
			}
		}
		if flag {
			temp := table[i]
			table[i] = table[smallest]
			table[smallest] = temp
		}
	}
}
