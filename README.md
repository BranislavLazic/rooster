# Rooster

Example of very primitive virtual machine

How to execute programs:

`go build main.go`

`./main program.rcode`

Example of program:

```
ICONST 7     // 0 - push 7 in stack
ICONST 5     // 2 - push 5 in stack
IADD         // 4 - add numbers
PRINT        // 5 - print the result of addition
HALT         // 6 - stop the program
```