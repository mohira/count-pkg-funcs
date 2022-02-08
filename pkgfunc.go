package pkgfunc

import (
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

	return nil, nil
}
