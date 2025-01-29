package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	charcount := 0
	valid := false
	result := ""
	for _, val := range str {
		if val != ' ' {
			charcount++
		}
		if charcount == 6 || (charcount == 5 && val == ' ') {
			result += " "
			charcount = 0
			valid = true
		} else if val != ' ' {
			result += string(val)
		}
	}
	if valid {
		return result + "\n"
	} else {
		return "Invalid Output\n"
	}
}
