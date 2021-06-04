package go128_test

import (
	"github.com/daisuke0131/go128"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func init() {
	go128.Analyzer.Flags.Set("num", "5")
}

func TestFromFileSystem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, go128.Analyzer, "a")
}