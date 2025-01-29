package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	charcount := 0
	valid := false
	result := ""
	for i, val := range str {
		if val != ' ' {
			charcount++
		}
		if (charcount == 6 && i != len(str)-1) || (charcount == 5 && val == ' ') {
			result += " "
			charcount = 0
			valid = true
		} else if val != ' ' && charcount != 6 {
			result += string(val)
		}
	}
	if valid || charcount == 5 {
		return result + "\n"
	} else {
		return "Invalid Output\n"
	}
}
