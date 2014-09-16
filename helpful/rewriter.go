package helpful

import (
	"fmt"
	"go/ast"
	"strings"
)

const (
	AnyName string = "Any"
)

type Rewriter struct {
	TypeName string
	Names    map[string]string
}

func (v Rewriter) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		for _, m := range n.Decls {
			if d, ok := m.(*ast.GenDecl); ok {
				for _, k := range d.Specs {
					if t, ok := k.(*ast.TypeSpec); ok {
						prev := t.Name.Name
						next := rename(prev, v.TypeName)

						t.Name.Name = next
						v.Names[prev] = next
					}
				}
			}
		}
	case *ast.Field:
		if f, ok := n.Type.(*ast.Ident); ok {
			if f.Name == AnyName {
				f.Name = v.TypeName
			}
			if val, ok := v.Names[f.Name]; ok {
				f.Name = val
			}
		}
	}
	return v
}

func rename(n string, t string) string {
	return fmt.Sprintf("%s%s", n, clean(strings.Title(t)))
}

func clean(n string) string {
	// remove pointers
	p := strings.Replace(n, "*", "Ptr", -1)
	// remove slices
	s := strings.Replace(p, "[]", "Slice", -1)

	return s
}
