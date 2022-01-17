// Package errorfmt provides a helper function to decorate errors
// with additional context.
package errorfmt

import "fmt"

// Handlef decorates an error with additional context.
//
//   func Frob(name string) (err error) {
//       defer errorfmt.Handlef("frob widget %s: %w", name, &err)
//       ...
//   }
//
// Its behavior is almost the same as fmt.Errorf,
// with a few differences:
//
//   1. There must be exactly one *error argument.
//   2. Instead of returning a value, Handlef assigns its
//      result via this pointer, to update the error in place.
//   3. The *error is dereferenced before formatting.
//      (Therefore, verb %w also accepts *error.)
//   4. If it points to a nil error, Handlef does nothing.
//
// These differences make it more convenient to use Handlef
// in a defer statement at the top of a block.
func Handlef(format string, a ...interface{}) {
	for i := range a {
		if p, ok := a[i].(*error); ok {
			if *p == nil {
				return // success case for caller, leave nil error alone
			}
			a2 := make([]interface{}, len(a))
			copy(a2, a)
			a2[i] = *p // Errorf wants error, not *error
			*p = fmt.Errorf(format, a2...)
			return
		}
	}
	panic("errorfmt: Handlef must be called with a *error arg")
}
