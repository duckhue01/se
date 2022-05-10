# Organizing code into package

## naming conventions for go packages
- go package names should be short and concise provide a clear indication of their purpose to the intended users of the package
- better to  abbreviate the package name if possible

## circular dependency
- for a go program to be well formed, its import graph must be acyclic; in other words, it must not contain any loops. Any violation of this predicate will cause the Go complier to emit an error. As the systems you are building grow in complexity, so does the probability of eventually hitting the dreaded import cycle detected error.

- usually, import cycles are an indication of a fault in the design of a software solution.

### breaking circular dependencies via implicit interface
- ./circular-dependency

### sometimes, code repetition is not a bad idea!

