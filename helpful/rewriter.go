package helpful

import (
	"fmt"
	"go/ast"
)

const (
	AnyValName string = "AnyVal"
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
						prev := fmt.Sprintf("%s", t.Name.Name)
						next := fmt.Sprintf("%s_%s", prev, v.TypeName)

						t.Name.Name = next
						v.Names[prev] = next
					}
				}
			}
		}
	case *ast.Field:
		if f, ok := n.Type.(*ast.Ident); ok {
			if f.Name == AnyValName {
				f.Name = v.TypeName
			}
			if val, ok := v.Names[f.Name]; ok {
				f.Name = val
			}
		}
	}
	return v
}
