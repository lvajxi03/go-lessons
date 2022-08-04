# First

First things first!

# Language and syntax

* Program needs to have `main` function defined as a standard entry point.
* Program needs to have `package main` preamble at the top, otherwise it will not be compiled with `go build` or started with `go run`
* Imports can be groupped into single instruction, "to improve readability" (well...)
* Imports make all `exported` symbols available for the program
* Variables or function names shall start with a capital letter, if intended to be exported
* Variables or function names shall start with a lowercase letter, if intended to be internal only
* Imported definitions can be accessed with `modulename.Defname`

# Building and runtime

* Program can be built with `go build filename.go` -- this will produce executable `filename` (or `filename.exe` in Windows)
* Produced binary is a ready-to-run program, holding all the necessary stuff inside.
* Program is statically linked, `ldd` says it's "not a dynamically linked program"

```bash
$ go build first.go
$ ls -lh first
1,7M

$ strip first
$ ls -lh first
1,2M
```

* Program can be run directly with `go run filename.go`
* Command above works only if `package main` preamble is present as well as function `main`
* Otherwise you get:

```bash
package command-line-arguments is not a main package 
```

when package preamble is missing, or:

```bash
untime.main_main·f: function main is undeclared in the main package
```

when no `main` function defined


That's all folks.
