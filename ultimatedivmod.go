package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a
	*a = *a / *b
	*b = *b % temp
}
