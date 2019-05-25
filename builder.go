// builder builds packages
package main

// analyze imports of given go files
func parseImports(sourceFiles []string) []string {

	pForImport := &parser{}
	// "fmt" depends on "os. So inject it in advance.
	// Actually, dependency graph should be analyzed.
	var imported []string = []string{"os"}
	for _, sourceFile := range sourceFiles {
		astFile := pForImport.parseFile(sourceFile, nil, true)
		for _, importDecl := range astFile.importDecls {
			for _, spec := range importDecl.specs {
				baseName := getBaseNameFromImport(spec.path)
				if !in_array(baseName, imported) {
					imported = append(imported, baseName)
				}
			}
		}
	}

	return imported
}

func compileStdLibs(p *parser, universe *Scope, imported []string) *compiledStdlib {

	// add std packages
	// parse std packages
	var libs *compiledStdlib = &compiledStdlib{
		compiledPackages: map[identifier]*AstPackage{},
		uniqImportedPackageNames:nil,
	}
	stdPkgs := makeStdLib()

	for _, spkgName := range imported {
		pkgName := identifier(spkgName)
		var pkgCode string
		var ok bool
		pkgCode, ok = stdPkgs[pkgName]
		if !ok {
			errorf("package '" + string(pkgName) + "' is not a standard library.")
		}
		pkg := parseStdPkg(p, universe, pkgName, pkgCode)
		libs.AddPackage(pkg)
	}

	return libs
}

func parseStdPkg(p *parser, universe *Scope, pkgname identifier, code string) *AstPackage {
	// initialize a package
	p.initPackage(pkgname)
	p.scopes[pkgname] = newScope(nil, string(pkgname))

	filename := string(pkgname) + ".memory"
	asf := p.parseString(filename, code, p.scopes[pkgname], false)

	p.resolve(universe)
	if debugAst {
		asf.dump()
	}
	return &AstPackage{
		name:  pkgname,
		files: []*AstFile{asf},
	}
}

func compileInputFiles(p *parser, pkgname identifier, sourceFiles []string) *AstPackage {
	p.initPackage(pkgname)
	p.scopes[pkgname] = newScope(nil, string(pkgname))
	var astFiles []*AstFile
	for _, sourceFile := range sourceFiles {
		asf := p.parseFile(sourceFile, p.scopes[pkgname], false)
		astFiles = append(astFiles, asf)
	}

	return &AstPackage{
		name: pkgname,
		files: astFiles,
	}
}

type compiledStdlib struct {
	compiledPackages map[identifier]*AstPackage
	uniqImportedPackageNames []string
}

func (csl *compiledStdlib) AddPackage(pkg *AstPackage) {
	csl.compiledPackages[pkg.name] = pkg
	if !in_array(string(pkg.name), csl.uniqImportedPackageNames) {
		csl.uniqImportedPackageNames = append(csl.uniqImportedPackageNames, string(pkg.name))
	}
}