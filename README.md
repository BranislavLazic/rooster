# Rooster [![Build Status](https://travis-ci.org/BranislavLazic/rooster.svg)](https://travis-ci.org/BranislavLazic/rooster)

Example of very primitive stack based virtual machine

Build executable:

`go build`

and execute programs:

`./rooster -sourceFile=program.rcode`

Example of program that adds two numbers and prints the result:

```
ICONST 7     # 0 - push 7 in stack
ICONST 5     # 2 - push 5 in stack
IADD         # 4 - add numbers
PRINT        # 5 - print the result of addition
HALT         # 6 - stop the program
```

More complex example that demonstrates calculation of product recursively:

```
# Use of recursion to calculate a product of numbers
ICONST 1
GSTORE 0 # Set initial accumulation value to 1

# Push four values on stack
ICONST 1
ICONST 2
ICONST 3
ICONST 4


CALL product 1 # Call "product" procedure recursively until the stack is depleted
GLOAD 0        # After the stack is depleted, load value from global memory
PRINT          # And print it out
HALT

product:
    
    GLOAD  0 # Load initial accumulation value from global memory
    LOAD   0 # Load value from frame stack
    IMUL     # Multiply value from global memory and the one from frame stack
    GSTORE 0 # Put the result in the same address in global memory
    JMP   12 # Repeat
    RET
```

More examples can be found in `programs` directory.

# Working principle

## Virtual machine

Virtual machine itself works on a quite simple principle. It basically ingests a "slice" (an array) of instructions represented with integers.
The "slice" of instructions is then matched against a "switch" statement. Rooster is a stack-based VM. It means that all operations are performed
through a virtual stack. E.g. if we want to add two integer numbers, we "push" those numbers on the stack with `ICONST` stack operation and then
we use `IADD` operation to "pop" them from stack, add them, and then "push" the result again on the stack. Rooster also features a memory for
global variables which can be manipulated via `GLOAD` and `GSTORE` operations. When it comes to procedures, Rooster creates a chunk of special 
local memory called "frame" that contains all local variables for a procedure and an address of instruction where it should return after the 
procedure finishes its computation. For each procedure, new "frame" is created and "pushed" on the stack of frames.

## Lexer

Lexer is a piece of our program that performs so called "tokenization" and checks for code grammar. Tokenization is esentially a process
of converting particular chunks of source code into tokens. And in the end, token is piece of source code that can have some meaning for us.
E.g. token is a comment. Keyword such as `ICONST` is also a token. Even very special characters such as "end of line" or "end of file" 
can be tokens.

## Parser

After lexer performs tokenization, we can parse our code. In parser, we are converting collection of tokens into instruction set which 
will be ingested by our VM. Notice that we are not building AST (abstract syntax tree). The reason for that is that our VM assembly (`rcode`) 
simply doesn't require such a thing. Parser will simply check for valid and relevant tokens. E.g. comment tokens will simply be discarded since 
they are irrelevant for VM. Also, if invalid token is present, parsing will fail. That's simply because our code is syntactically incorrect.
