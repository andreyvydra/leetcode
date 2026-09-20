package tasks

func reverseDegree(s string) int {
	su := 0
	for idx, c := range s {
		num := 26 - (int(c) % 97)
		su += (num) * (idx + 1)
	}
	return su
}
