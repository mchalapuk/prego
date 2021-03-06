prego
=====

A tiny library that contains utilities for checking functions' preconditions
in [Go programming language](http://golang.org). Precondition checking
is a concept intruduced in code design approach called [Design
by Contract](http://en.wikipedia.org/wiki/Contract_programming).
If&nbsp;preconditions of a function are not satisfied then, according
to [Crash Early Principle](http://pragmatictips.com/32), program
is invalid and&nbsp;should be immediately terminated. Termination should
be preceeded with error message. Precond simplifies doing it all
(contract checking, printing error and crashing) to&nbsp;a single
function call.

Fetch muroc/prego library to your go workspace!

```bash
go get github.com/muroc/prego
```

How To Use It
-------------

As many other languages, Go does not provide nillability information
in its type system. Precond fills this gap by providing *precond.NotNil*
function. Checking if argument is not nil is the most simple and most common
way of using the library.

```go
import "github.com/muroc/prego"

func trim(noBeTrimmed string) string {
  // if toBeTrimmed is nil, precond.NotNil will panic
  // with error passed as second argument
  precond.NotNil(toBeTrimmed, "trimmed string must not be null")

  ... // function body
}
```

### Error Formatting

All methods in the *precond* and *check* namespace support
passing arguments to be used when formatting error message.
[fmt.Errorf](http://golang.org/pkg/fmt/#Errorf) is used to create
error instances. All [fmt formatting rules](http://golang.org/pkg/fmt/)
apply.

```go
import "github.com/muroc/prego"

func (set *Set) get(index int) interface{} {
  // if index is not contained in <0, length> precond.InRange
  // eill panic with properly formatted error
  precond.InRange(index, 0, set.Length()-1, "index %v out of bounds", index)

  ... // function body
}
```

Other panic checks can be found in
[panic.go](https://github.com/gosmos/precond/blob/master/panic.go).

### Defensive Programming

Package *precond/check* contain function that does the same checks,
but return encountered errors instead of&nbsp;panic. They will come in handy
when writing code that uses technique called [Defensive
Programming](http://en.wikipedia.org/wiki/Defensive_programming).

```go
import "github.com/muroc/prego/check"

func handleMessage(msg string) string {
  // if isProerFormat(msg) is false, check.True will return error
  // with message passed as second argument
  err := check.True(isProperFormat(msg), "ill-formated message: %v", message)
  if err != nil {
    log.Error(err)
    return
  }

  ... // function body
}
```

Other error checks can be found in
[check/check.go](check/check.go).

License
-------

Copyright &copy; 2016 Maciej Chałapuk. Released under [Apache License 2.0](LICENSE).

