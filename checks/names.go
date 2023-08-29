// Copyright 2021 Pinterest
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checks

import (
	"reflect"
	"regexp"

	"github.com/pinterest/thriftcheck"
	"go.uber.org/thriftrw/ast"
)

// Name returns an ast.Node's Name string.
func nodeName(node ast.Node) string {
	if v := reflect.ValueOf(node); v.Kind() == reflect.Ptr {
		if f := v.Elem().FieldByName("Name"); f.IsValid() {
			return f.Interface().(string)
		}
	}
	return ""
}

var (
	screamingSnakeCaseRE = regexp.MustCompile("^[A-Z]+(_[A-Z]+)*$")
	camelCaseCheck       = regexp.MustCompile(`^[a-z]+([A-Z][a-z]+)*$`)
	pascalCaseCheck      = regexp.MustCompile(`^([A-Z][a-z]+)+$`)
)

// CheckNamesReserved checks if a node's name is in the list of reserved names.
func CheckNamesReserved(names []string) *thriftcheck.Check {
	reserved := make(map[string]bool)
	for _, name := range names {
		reserved[name] = true
	}

	return thriftcheck.NewCheck("names.reserved", func(c *thriftcheck.C, n ast.Node) {
		if name := nodeName(n); name != "" && reserved[name] {
			c.Errorf(n, "%q is a reserved name", name)
		}
	})

}

// CheckNamesCasing checks if each node complies with the casing rules.
func CheckNamesCasing() *thriftcheck.Check {
	return thriftcheck.NewCheck("names.casing", func(c *thriftcheck.C, n ast.Node) {

		switch n.(type) {

		// types that should be pascal case
		case *ast.Enum, *ast.Service, *ast.Struct, *ast.Typedef:
			if !pascalCaseCheck.MatchString(nodeName(n)) {
				c.Errorf(n, "%q is not pascal case", nodeName(n))
			}

		// types that should be camel case
		case *ast.Function, *ast.Field:
			if !camelCaseCheck.MatchString(nodeName(n)) {
				c.Errorf(n, "%q is not camel case", nodeName(n))
			}

		//types that should be screaming snake case
		case *ast.Constant, *ast.EnumItem:
			if !screamingSnakeCaseRE.MatchString(nodeName(n)) {
				c.Errorf(n, "%q is not screaming snake case", nodeName(n))
			}

		// Checked in CheckNamespacePattern ("namespace.patterns")
		case *ast.Namespace:
		// Root node, do nothing
		case *ast.Program:
		// Nothing to be done
		case ast.BaseType, ast.MapType, ast.ListType, ast.SetType:
		// Nothing to be done
		default:
		}

	})
}
