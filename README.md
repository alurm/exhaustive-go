# An emulation of exhaustive pattern matching and tagged unions in Go

Tagged unions can be represented as closed interfaces. In the example, it is named `isf.T`, which holds an int, string or a float.

Exhaustive switch can be represented as a function, taking the tagged union instance and a handler function for each match case. In the example, the function is named `isf.Match`. It receives an `isf.T` and a bunch of handlers and does the dispatch internally.

The file `main.go` has usages examples.

I doubt this will perform well, unless inlined by the compiler. Anyway, since performance requirements vary depending on the task at hand, this may be acceptable.

## License

MIT.
