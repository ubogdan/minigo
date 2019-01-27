package main

import "fmt"

var gmap map[string]bool

//var debug [6]int // r10, r11,...


func f1() {
	var lmap map[string]bool
	fmt.Printf("%d\n", len(gmap) + 1) // 1
	fmt.Printf("%d\n", len(lmap) + 2) // 2
}

func f2() {
	var lmap map[int]int = map[int]int{
		4:7,
		5:8,
		6:9,
	}
	fmt.Printf("%d\n", len(lmap)) // 3
	for i := range lmap {
		fmt.Printf("%d\n", i) // 4,5,6
	}
}

func f3() {
	var lmap map[int]int = map[int]int{
		7:  8,
		9:  10,
		11: 12,
	}

	lmap[13] = 14
	lmap[15] = 16
	for i, v := range lmap {
		fmt.Printf("%d\n", i)
		fmt.Printf("%d\n", v)
	}
}

func f4() {
	var lmap map[int]int = map[int]int{
		7:  17,
		9:  10,
		11: 12,
		0:  18,
	}

	fmt.Printf("%d\n", lmap[7]) // 17
	fmt.Printf("%d\n", lmap[0]) // 18

	/*
	for i, v := range debug {
		fmt.Printf("debug[1%d]=%d\n", i,v)
	}
	*/
}

/*
func f9() {
	var lmap map[string]int
	value := lmap["hello"]
	fmt.Printf("%d\n", value + 13)
	//lmap["x"] = 1
	value = lmap["hello"]
	fmt.Printf("%d\n", value + 14)
}
*/

func main() {
	f1()
	f2()
	f3()
	f4()
	//f9()
}
