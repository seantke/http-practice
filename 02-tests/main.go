package main

func main() {
	multStuff(1, 1)
}

func multStuff(x, y int) int {
	// xs := []interface{}{
	// 	1, 2, 3, 4, 5,
	// }
	xs := []int{1, 2, 3, 4, 5}
	total := 1
	for _, x := range xs {
		total *= x
	}
	return total
}

//check out  goconvey
