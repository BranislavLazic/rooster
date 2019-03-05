# Opcodes

Rooster assembly uses following set of opcodes

## ICONST

ICONST opcode is used to push integer value on the stack.
The value of the last executed ICONST opcode will be on top of the stack.
Example:

```
ICONST 3
ICONST 6
```

Stack will look like this:

| Pos | Value |
| --: | ----: |
|   1 |     6 |
|   0 |     3 |
