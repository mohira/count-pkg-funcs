package main

import (
	"pkgfunc"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(pkgfunc.Analyzer) }
