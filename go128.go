package go128

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var num int

func init() {
	Analyzer.Flags.IntVar(&num, "num", 5,"Number of parameters")
}

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
			if len(fdecl.Type.Params.List) >= num{
				pass.Reportf(fdecl.Pos(),"%sは引数が%d個以上あります！!\n", fdecl.Name.Name,num)
			}
		}
	}

	return nil, nil
}