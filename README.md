# An example of emulation of tagged unions and of exhaustive pattern matching in Go

## Tagged unions

A tagged union can be represented as a closed interface.

In the provided example, there is a closed interface `isf.T`, which holds either an int, a string or a float. Well, an `isf.Int`, an `isf.String`, or an `isf.Float`, technically.

## Exhaustive pattern matching

To match on the union, we need a matching construct.

A pattern matching construct can be represented as dispatching function, which takes our union and a handling closure for each possible case in the union.

In the provided example, the function is named `isf.Match`. It receives an `isf.T` and a bunch of handlers.

## Usage

The file `main.go` is provided with usage examples.

## Performance

No measurements have been conducted.

I doubt this will perform well, unless inlined by the compiler.

Since performance requirements vary depending on the task at hand, this may be acceptable.

## License

The expat license.
