# Showcase of exhaustive pattern matching and tagged unions in Go

## Introduction

Welcome!

I come from functional programming background (Haskell, OCaml, etc.), so I missed exhaustive pattern matching and tagged unions in Go. So I've been wondering how to implement them.

Luckly, it's a possible task. You don't need reflection or generics for that. It's not a lot of code, simple stuff.

This module shows one of the ways to implement and use these things in Go. This is not a library, it's a pattern.

Note: this pattern should not be used in performance sensitive code.

License: MIT.

## Features

- No reflection or generics used.
- Not too much boilerplate required, implementation is easy to understand.
- Type safe.

## Details

`isf.T` is a tagged union, which holds an int, a string or a float. `isf.Match` is used for pattern matching. It receives a `isf.T` and a handler for each of the possible cases. `main.go` has an example how this should be used, check it out.

I'm open for suggestions on how to improve this :)

## Future goals

Add more examples. Add an example of a simple enum. Add an example of a more complicated tagged union (e.g. a binary tree). Figure out how to generate implementation boilerplate programmatically.
