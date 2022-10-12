// Package isf implements a type, which can be either an int, a string or a float.
package isf

type (
	// T is a closed interface, which is only implemented by intType, stringType and float64Type.
	T interface{ m() }
	// intType holds an int.
	intType int
	// stringType holds a string.
	stringType string
	// float64Type holds a float64.
	float64Type float64
)

// m makes receiver implement T.

func (_ intType) m()     {}
func (_ stringType) m()  {}
func (_ float64Type) m() {}

// Constructors

// Int wraps an int.
func Int(x int) intType { return intType(x) }

// String wraps a string.
func String(x string) stringType { return stringType(x) }

// Float64 wraps a float64.
func Float64(x float64) float64Type { return float64Type(x) }

// Match exhaustively switches on all possibles types it interface can contain and passes underlying type to the corresponding handler.
func Match(it T, i func(int), s func(string), f func(float64)) {
	switch it := it.(type) {
	case intType:
		i(int(it))
	case stringType:
		s(string(it))
	case float64Type:
		f(float64(it))
	}
}
