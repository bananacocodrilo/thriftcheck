package thriftcheck

import (
	"reflect"

	"go.uber.org/thriftrw/ast"
)

var nodeInterface = reflect.TypeOf((*ast.Node)(nil)).Elem()

// VisitorFunc adapts a function to the ast.Visitor interface. This differs
// from ast.VisitorFunc in that is supports an ast.Visitor-compativle return
// value.
type VisitorFunc func(ast.Walker, ast.Node) VisitorFunc

func (f VisitorFunc) Visit(w ast.Walker, n ast.Node) ast.Visitor {
	return f(w, n)
}

// Annotations returns an ast.Node's Annotations.
func Annotations(node ast.Node) []*ast.Annotation {
	if v := reflect.ValueOf(node); v.Kind() == reflect.Ptr {
		if f := v.Elem().FieldByName("Annotations"); f.IsValid() {
			return f.Interface().([]*ast.Annotation)
		}
	}
	return nil
}
