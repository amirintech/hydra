# 🐍 Hydra - A Simple Interpreter in Go

**Hydra** is an interpreter for a programming language implemented in Go. This project serves as a learning exercise in building interpreters and understanding programming language concepts.

## 🚀 Project Overview

- **Lexical Analysis**: Tokenizes source code into meaningful symbols.
- **Parsing**: Constructs an Abstract Syntax Tree (AST) from tokens.
- **Evaluation**: Executes the AST to produce results.

## 🛠️ Technologies Used

- **Go**: The primary programming language for implementation.

## 📂 Project Structure

```
hydra/
│── ast/               # Abstract Syntax Tree structures
│── eval/              # Evaluation logic
│── lexer/             # Lexical analyzer
│── object/            # Object system definitions
│── parser/            # Parser for syntax analysis
│── repl/              # Read-Eval-Print Loop implementation
│── token/             # Token definitions
│── .gitignore         # Git ignore file
│── go.mod             # Go module file
│── main.go            # Entry point of the interpreter
│── Makefile           # Build and run commands
```

## 🛠️ Features

- **Variables and Bindings**: Supports variable declarations and scope management.
- **Arithmetic Operations**: Handles basic arithmetic like addition, subtraction, etc.
- **Boolean Logic**: Implements logical operations and conditionals.
- **Functions**: Allows function definitions and calls.

## 🚧 Status

This project is a work in progress and currently **incomplete**. Contributions and suggestions are welcome to enhance its functionality.
