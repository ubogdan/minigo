package main

import (
	"os"
	"strings"
)

var GENERATION int = 1

var debugMode = false // execute debugf() or not
var debugToken = false

var debugAst = false
var debugParser = false
var tokenizeOnly = false
var parseOnly = false
var resolveOnly = false

func printVersion() {
	println("minigo 0.1.0")
	println("Copyright (C) 2019 @DQNEO")
}

func parseOpts(args []string) []string {
	var r []string

	for _, opt := range args {
		if opt == "--version" {
			printVersion()
			return nil
		}
		if opt == "-t" {
			debugToken = true
		}
		if opt == "-a" {
			debugAst = true
		}
		if opt == "-p" {
			debugParser = true
		}
		if opt == "-d" {
			debugMode = true
		}
		if opt == "--tokenize-only" {
			tokenizeOnly = true
		}
		if opt == "--parse-only" {
			parseOnly = true
		}
		if opt == "--resolve-only" {
			resolveOnly = true
		}
		if strings.HasSuffix(opt, ".go") {
			r = append(r, opt)
		} else if opt == "-" {
			return []string{"/dev/stdin"}
		}
	}

	return r
}

func main() {
	// parsing arguments
	var sourceFiles []string

	assert(len(os.Args) > 0, nil,"os.Args should not be empty")
	if len(os.Args) > 1 {
		sourceFiles = parseOpts(os.Args[1:len(os.Args)])
	}

	if len(sourceFiles) == 0 {
		println("No input files.")
		return
	}

	if tokenizeOnly {
		dumpTokenForFiles(sourceFiles)
		return
	}

	// parser starts
	p := &parser{}
	p.scopes = map[identifier]*Scope{}
	p.initPackage("")

	allScopes = p.scopes

	// setup universe scopes
	universe := newUniverse()
	var globalStringLiterals []*ExprStringLiteral

	// inject builtin functions into the universe
	p.stringLiterals = nil
	internalUniverse := p.parseString("internal_universe.go", internalUniverseCode, universe, false)
	p.resolve(nil)
	for _, sl := range p.stringLiterals {
		globalStringLiterals = append(globalStringLiterals, sl)
	}

	// inject runtime things into the universe
	p.stringLiterals = nil
	internalRuntime := p.parseString("internal_runtime.go", internalRuntimeCode, universe, false)
	p.resolve(nil)
	for _, sl := range p.stringLiterals {
		globalStringLiterals = append(globalStringLiterals, sl)
	}

	// compile stdlibs which are imporetd from userland
	imported := parseImports(sourceFiles)
	stdlibs := compileStdLibs(p, universe, imported)

	// compile the main package
	mainPkg := ParseSources(p, identifier("main"), sourceFiles, false)
	if parseOnly {
		if debugAst {
			mainPkg.dump()
		}
		return
	}
	p.resolve(universe)
	if debugAst {
		mainPkg.dump()
	}

	if resolveOnly {
		return
	}

	setTypeIds(p.allNamedTypes)
	ir := makeIR(internalUniverse, internalRuntime, stdlibs, mainPkg , globalStringLiterals, p.allDynamicTypes)
	ir.emit()
}
