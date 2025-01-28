package piscine

func Rot14(s string) string {
	result := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if v+14 > 'z' {
				result += string((13 - ('z' - v)) + 'a')
			} else {
				result += string(v + 14)
			}
		} else if v >= 'A' && v <= 'Z' {
			if v+14 > 'Z' {
				result += string((13 - ('Z' - v)) + 'A')
			} else {
				result += string(v + 14)
			}
		} else {
			result += string(v)
		}
	}
	return result
}
