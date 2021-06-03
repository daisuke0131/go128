package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main(){
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset,"testdata/a.go",nil,0)
	if err != nil{
		panic(err)
	}

	for _,d := range f.Decls{
		fdecl,ok :=  d.(*ast.FuncDecl)
		if !ok{
			continue
		}
		if len(fdecl.Type.Params.List) >= 5{
			fmt.Printf("%sは引数が5個以上あります！!\n", fdecl.Name.Name)
		}
	}

}
