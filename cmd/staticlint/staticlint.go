package staticlint

import (
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var osExitAnalyzer = &analysis.Analyzer{
	Name:     "don`t use os.Exit() in main func",
	Doc:      "don't allow os.Exit() in your go code",
	Run:      osExit,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func osExit(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ce := n.(*ast.CallExpr)
		ident, ok := ce.Fun.(*ast.Ident)
		if !ok {
			return
		}
		if ident.Name == "Exit" {
			pass.Reportf(ce.Pos(), "don't use os.Exit()")
		}
	})

	return nil, nil
}

var duplicateAnalyzer = &analysis.Analyzer{
	Name:     "imports_duplicate",
	Doc:      "detect duplicate package imports",
	Run:      duplicate,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func duplicate(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		packageAll := make(map[string]bool, 0)
		for _, i := range f.Imports {
			if i.Name.String() == "_" {
				continue
			}
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				pass.Reportf(i.Path.Pos(), "import not quoted")
			}
			if packageAll[path] {
				pass.Reportf(i.Path.Pos(), "duplicate import %s", path)
			}
			packageAll[path] = true
		}
	}
	return nil, nil
}
