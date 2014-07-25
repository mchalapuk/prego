precond
=======

A library that contains utilities for checking functions' preconditions
in [Go programming language](http://golang.org). Precondition checking
is a concept intruduced in code design approach called [Design
by Contract](http://en.wikipedia.org/wiki/Contract_programming).
If preconditions of a function are not satisfied then, according
to [Crash Early Principle](http://pragmatictips.com/32), program
is invalid and should be immediately terminated. Termination should
be preceeded with proper error message.

Fetch gosmos/precond library to your go workspace!

```bash
go get github.com/gosmos/precond
```

How To Use It
-------------

As many other languages, Go does not provide nillability information
in its type system. Precond fills this gap by providing *precond.NotNil*
function. Checking if argument is not nil is the most simple and most common
way of using the library.

```go
import "github.com/gosmos/precond"

func trim(noBeTrimmed string) string {
  // if toBeTrimmed is nil, precond.NotNil will panic
  // with error passed as second argument
  precond.NotNil(toBeTrimmed, "trimmed string must not be null");

  ...
}
```

All methods in the *precond* and *check* namespace support passing arguments
to be used when formatting error message.

```go
import "github.com/gosmos/precond"

func (set *Set) get(index int) interface{} {
  // if index is not contained in <0, length> precond.InRange
  // eill panic with properly formatted error
  precond.InRange(index, 0, set.Length()-1, "index %v out of bounds", index);

  ...
}
```

Other checks can be found in
[panic.go](https://github.com/gosmos/precond/blob/master/panic.go).

Contradicting Official Documentation
------------------------------------

In "Effective Go", which is a part of official documentation
of Go language, it is suggested that
[functions should avoid panic](http://golang.org/doc/effective_go.html#panic).
This library contradicts that guideline.

Avoiding panic and trying to silently ignore errors is a form of
[Defensive Programming](http://en.wikipedia.org/wiki/Defensive_programming),
which is a fine technique, that should be used when dealing with unpredictable
inputs. As not all programs deal with such situations and certainly not all
parts of the program handle unpredictable input, usage of defensive
programming should be limited.

*Design by Contract* and *Crash Early*, on the other hand, are proven
to be very useful techniques that significantly reduce number of bugs.
They should be used extensively in most of today's software.

