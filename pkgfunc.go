package pkgfunc

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

const doc = "pkgfunc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "pkgfunc",
	Doc:  doc,
	Run:  run,
	//Requires: []*analysis.Analyzer{
	//	//inspect.Analyzer,
	//},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		//ast.Print(pass.Fset, file)
		for _, decl := range file.Decls {
			//fmt.Printf("%T %v\n", decl, decl)
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			fmt.Println(funcDecl.Name)

		}
	}
	return nil, nil
}
