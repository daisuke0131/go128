package main

import (
	"github.com/daisuke0131/go128"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(go128.Analyzer)
}
