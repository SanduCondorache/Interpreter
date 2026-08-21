# Interpreter

A small interpreted programming language implemented in **Go**.

This project implements the core components of a programming language from the ground up, including a lexer, Pratt parser, Abstract Syntax Tree (AST), evaluator, runtime objects, and an interactive REPL.

The project is primarily intended as a learning project for understanding how programming languages and interpreters work internally.

## Features

The language currently supports:

- Integer literals
- Boolean values
- String literals
- Arrays
- Variable bindings
- `let` statements
- `return` statements
- Prefix operators
- Infix operators
- Comparison operators
- Conditional expressions
- Functions
- Function calls
- Array indexing
- Lexical environments
- Runtime error handling
- Interactive REPL

### Supported operators

Arithmetic:

```text
+   -   *   /
```

Comparison:

```text
==   !=   <   >
```

Prefix operators:

```text
!   -
```

## Example

A simple program looks like this:

```text
let x = 10;
let y = 20;

let add = fn(a, b) {
    return a + b;
};

add(x, y);
```

Arrays and indexing are also supported:

```text
let numbers = [1, 2, 3, 4];

numbers[0];
numbers[2];
```

Conditional expressions can be used to control execution:

```text
if (5 < 10) {
    100;
} else {
    200;
}
```

Strings are supported as well:

```text
"Hello, world!";
```

## Architecture

The interpreter is divided into several components:

```text
Source Code
    │
    ▼
  Lexer
    │
    ▼
  Tokens
    │
    ▼
  Parser
    │
    ▼
   AST
    │
    ▼
 Evaluator
    │
    ▼
Runtime Objects
    │
    ▼
   Output
```

### Lexer

The lexer converts source code into a stream of tokens.

It handles:

- Identifiers
- Integer literals
- Strings
- Keywords
- Operators
- Delimiters
- Brackets
- Illegal tokens

The lexer is implemented in `internals/lexer`.

### Parser

The parser converts the token stream into an Abstract Syntax Tree.

Expression parsing uses operator precedence, allowing expressions such as:

```text
1 + 2 * 3
```

to be interpreted with the expected precedence.

The parser also handles:

- Variable declarations
- Return statements
- Prefix expressions
- Infix expressions
- Boolean expressions
- Conditional expressions
- Function literals
- Function calls
- String literals
- Array literals
- Array indexing

The parser implementation is located in `parser`.

### Abstract Syntax Tree

The `ast` package contains the nodes used to represent parsed programs.

Examples include:

- `Program`
- `LetStatement`
- `ReturnStatement`
- `ExpressionStatement`
- `InfixExpression`
- `PrefixExpression`
- `IfExpression`
- `FunctionLiteral`
- `CallExpression`
- `ArrayLiteral`
- `IndexExpression`

### Evaluator

The evaluator walks the AST and executes the program.

It is responsible for:

- Evaluating expressions
- Performing arithmetic
- Comparing values
- Evaluating conditionals
- Managing variables
- Creating functions
- Calling functions
- Evaluating arrays
- Performing array indexing
- Propagating return values
- Reporting runtime errors

The evaluator is implemented in `evaluator`.

### Runtime Objects

Runtime values are represented by objects in `internals/object`.

The runtime includes objects for values such as:

```text
Integer
Boolean
String
Array
Function
Null
ReturnValue
Error
```

An environment is used to store variable bindings and provide the scope in which expressions and functions are evaluated.

### REPL

The project includes an interactive Read-Eval-Print Loop.

Running the interpreter starts a session where source code can be entered and evaluated interactively.

## Project Structure

```text
.
├── ast/
│   ├── ast.go
│   └── ast_test.go
│
├── cmd/
│   └── main/
│       └── main.go
│
├── evaluator/
│   ├── builtins.go
│   ├── evaluator.go
│   └── evaluator_test.go
│
├── internals/
│   ├── lexer/
│   ├── object/
│   ├── repl/
│   └── token/
│
├── parser/
│   ├── parser.go
│   └── parser_test.go
│
├── Makefile
├── go.mod
└── README.md
```

## Requirements

- [Go](https://go.dev/) 1.23.1 or newer

The module is configured for Go 1.23.1.

## Getting Started

Clone the repository:

```bash
git clone https://github.com/SanduCondorache/Interpreter.git
cd Interpreter
```

Run the interpreter:

```bash
make run
```

Alternatively, run it directly with Go:

```bash
go run ./cmd/main/main.go
```

The executable starts the language's REPL.

## Building

Build the project with:

```bash
make build
```

Or use Go directly:

```bash
go build ./...
```

The repository's Makefile provides a dedicated `build` target.

## Testing

Tests are organized by interpreter component.

Run the AST tests:

```bash
make test-ast
```

Run the parser tests:

```bash
make test-parser
```

Run the lexer tests:

```bash
make test-lexer
```

Run the object-system tests:

```bash
make test-object
```

Run the evaluator tests:

```bash
make test-evaluator
```

You can also run the complete test suite with:

```bash
go test ./...
```

The repository contains tests for the AST, parser, lexer, object system, and evaluator.
