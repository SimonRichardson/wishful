package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/SimonRichardson/wishful/helpful"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	const (
		fileUsage = "Path to monad template."
		typeUsage = "Type to which to specialize."
	)

	var filename string
	var typename string

	flag.StringVar(&filename, "file", "", fileUsage)
	flag.StringVar(&filename, "f", "", fileUsage)
	flag.StringVar(&typename, "type", "", typeUsage)
	flag.StringVar(&typename, "t", "", typeUsage)
	flag.Parse()

	if filename == "" {
		fmt.Println("-file is required")
		os.Exit(255)
	}
	if typename == "" {
		fmt.Println("-type is required")
		os.Exit(255)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		fmt.Printf("Oops! Can't parse the source: %v\n", err)
		return
	}

	v := helpful.Rewriter{typename, make(map[string]string)}

	ast.Walk(v, f)
	newSrc, err := writeFile(f, fset)
	if err != nil {
		fmt.Printf("Oops! Can't write out AST: %v\n", err)
		os.Exit(1)
	}

	err = checkSyntax(fset, newSrc)
	if err != nil {
		fmt.Printf("Oops! Generated code can't compile: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("%s\n", newSrc)
}

func writeFile(f *ast.File, fset *token.FileSet) ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func checkSyntax(fset *token.FileSet, src []byte) (err error) {
	_, err = parser.ParseFile(fset, "", src, 0)
	return
}
