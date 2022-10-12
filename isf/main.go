// Package isf implements a type, which can be either an int, a string or a float.
package isf

type (
	// T is a closed interface, which is only implemented by Int, String and Float32.
	T       interface{ m() }
	Int     int
	String  string
	Float64 float64
)

// m makes receiver implement T.

func (_ Int) m()     {}
func (_ String) m()  {}
func (_ Float64) m() {}

// Match exhaustively switches on all possibles types it interface can contain and passes underlying type to the corresponding handler.
func Match(it T, i func(int), s func(string), f func(float64)) {
	switch it := it.(type) {
	case Int:
		i(int(it))
	case String:
		s(string(it))
	case Float64:
		f(float64(it))
	}
}
