package main

import (
	"os"
)

func myPrintf(format gostring, a []interface{}) {
	var s gostring = Sprintf(gostring(format), a...)
	os.Stdout.Write(s)
}

func f0() {
	var a []interface{}
	myPrintf(S("hello\n"), a)
}

func f1() {
	var a []interface{}
	var i int = 123
	var ifc interface{}
	ifc = i
	a = append(a, ifc)
	myPrintf(S("%d\n"), a)
}

func f2() {
	var a []interface{}
	var i int = 123
	var i2 int = 456
	var ifc interface{}
	var ifc2 interface{}
	ifc = i
	ifc2 = i2
	a = nil
	a = append(a, ifc)
	a = append(a, ifc2)
	myPrintf(S("%d %d\n"), a)
}

func f3() {
	var a []interface{}
	var s gostring = S("hello")
	var s2 gostring = S("world")
	var ifc interface{}
	var ifc2 interface{}
	ifc = s
	ifc2 = s2
	a = append(a, ifc)
	a = append(a, ifc2)
	myPrintf(S("%s %s\n"), a)
}

func f4() {
	var a []interface{}
	var s gostring = S("hello")
	var i int = 123
	var ifc interface{}
	var ifc2 interface{}
	ifc = s
	ifc2 = i
	a = append(a, ifc)
	a = append(a, ifc2)
	myPrintf(S("%s %d\n"), a)
}

func f5() {
	var a []interface{}
	var s gostring = S("hello")
	var i int = 123
	var i2 int = 456
	var ifc interface{}
	var ifc2 interface{}
	var ifc3 interface{}
	ifc = s
	ifc2 = i
	ifc3 = i2
	a = append(a, ifc)
	a = append(a, ifc2)
	a = append(a, ifc3)
	myPrintf(S("%s %d %d\n"), a)
}

func test_dumpToken() {
	format := "string=%s,int=%d\n" // "string=abcdefg,int=12345"
	var s1 gostring = S("abcdefg")
	var s2 int = 12345
	var ifcs []interface{}
	var ifc1 interface{} = s1
	var ifc2 interface{} = s2
	ifcs = append(ifcs, ifc1)
	ifcs = append(ifcs, ifc2)
	b := Sprintf(gostring(format), ifcs...)
	os.Stdout.Write(b)
}

func main() {
	f0()
	f1()
	f2()
	f3()
	f4()
	f5()
	test_dumpToken()
}
