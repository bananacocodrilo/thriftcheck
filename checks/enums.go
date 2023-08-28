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
	"github.com/pinterest/thriftcheck"
	"go.uber.org/thriftrw/ast"
)

// CheckEnumSize returns a thriftcheck.Check that warns or errors if an
// enumeration's element size grows beyond a limit.
func CheckEnumSize(warningLimit, errorLimit int) *thriftcheck.Check {
	return thriftcheck.NewCheck("enum.size", func(c *thriftcheck.C, e *ast.Enum) {
		size := len(e.Items)
		if errorLimit > 0 && size > errorLimit {
			c.Errorf(e, "enumeration %q has more than %d items", e.Name, errorLimit)
		} else if warningLimit > 0 && size > warningLimit {
			c.Warningf(e, "enumeration %q has more than %d items", e.Name, warningLimit)
		}
	})
}

// CheckEnumExplicit returns a thriftcheck.Check that warns or errors if an
// enumeration's element uses implicit values.
func CheckEnumExplicit() *thriftcheck.Check {
	return thriftcheck.NewCheck("enum.explicit", func(c *thriftcheck.C, e *ast.Enum) {

		negativeValueFound := false
		for _, item := range e.Items {
			if item.Value != nil && *item.Value < 0 {
				negativeValueFound = true
			}
		}
		for _, item := range e.Items {
			if item.Value == nil {
				if negativeValueFound {
					c.Errorf(item, "enumeration %q contains a negative value. In this case is mandatory to use explicit values in the rest of the entries.", e.Name)
				} else {
					c.Warningf(item, "enumeration %q has implicit value for %q. It is recommended to use explicit values in all enums", e.Name, item.Name)
				}

				break
			}
		}

	})
}
