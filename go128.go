package go128

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

const doc = "go128 is try"

var Analyzer = &analysis.Analyzer{
	Name:             "go128",
	Doc:              doc,
	Run:              run,
	RunDespiteErrors: true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _,f := range pass.Files{
		for _,d := range f.Decls{
			fdecl,ok :=  d.(*ast.FuncDecl)
			if !ok{
				continue
			}
			if len(fdecl.Type.Params.List) >= 5{
				pass.Reportf(fdecl.Pos(),"%sは引数が5個以上あります！!\n", fdecl.Name.Name)
			}
		}
	}

	return nil, nil
}