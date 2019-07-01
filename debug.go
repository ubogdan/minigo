package main

import (
	"os"
)

func debugf(format string, v ...interface{}) {
	if !debugMode {
		return
	}

	var indents []byte

	for i := 0; i < debugNest; i++ {
		indents = append(indents, ' ')
		indents = append(indents, ' ')
	}
	os.Stderr.Write(indents)
	s2 := Sprintf(gostring(format), v...)
	var b []byte = []byte(s2)
	b = append(b, '\n')
	os.Stderr.Write(b)
}

var debugNest int

// States "To Be Implemented"
func TBI(tok *Token, format string, v ...interface{}) {
	errorft(tok, "(To Be Implemented) "+format, v...)
}

// errorf with a position token
func errorft(tok *Token, format string, v ...interface{}) {
	var tokString string
	if tok != nil {
		tokString = tok.String()
	}
	errorf(format+"\n "+tokString, v...)
}

func errorf(format string, v ...interface{}) {
	s := Sprintf(gostring(format), v...)
	panic(s)
}

func assert(cond bool, tok *Token, format string, v ...interface{}) {
	if !cond {
		s := Sprintf(gostring(format), v...)
		msg := concat3(S("assertion failed: "), s,  tok.GoString())
		panic(msg)
	}
}

func assertNotReached(tok *Token) {
	msg := concat(S("assertNotReached "), S(tok.String()))
	panic(msg)
}

func assertNotNil(cond bool, tok *Token) {
	assert(cond, tok, "should not be nil")
}
