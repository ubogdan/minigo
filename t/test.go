package main

var ga int

func fa() {
	printf("%d\n", ga)
	printf("")// this is comment
	printf("") // this is comment
}


/* this
is
comment
* /
*
/

*/

func fb() {
	printf("%d\n", 2 - 1)
	printf("%d\n", 1 + 1)
	printf("%d\n", 1 + 1 + 1)
	printf("%d\n", 2 * 2)
	printf("%d\n", 2 * 3 - 1)
	printf("%d\n", 1 + 1 * 5)
}

func fc() {
	var i int
	i = 3
	printf("%d\n", i + 4)
}

func fd() {
	var j int = 2
	printf("%d\n", j + 6)
}

func fe() {
	var a int = 5
	var b int = 4
	printf("%d\n", a + b)
}

var gb int = 10

func ff() {
	printf("%d\n", gb)
}


func main() {
	fa()
	fb()
	fc()
	fd()
	fe()
	ff()
	println("hello world")
}

