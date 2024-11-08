# bantam-go

This is my little `learn-by-doing` project to demonstrate and understand the basic concept of **Pratt parsing**.

This app is a `Go` port of `Bantam`, originally written in [Java](https://github.com/munificent/bantam).

If you want to see the full explanation, check out this [awesome blog post](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/) by **Bob Nystrom**.

Making my `Go` version here, I tried to:
- keep the original structure of the project
- change some code to fit into [idiomatic way](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/) of `Go` programming

## Usage
Run the tests:
```bash
go run main.go
```

## Project Structure
```
bantam-go/
├── tokentype/     # Token type definitions
├── token/         # Token structure
├── parser/        # Core parsing logic
├── parselet/      # Parsing rules for different expressions
├── lexer/         # Lexical analyzer
├── expression/    # Expression AST nodes
├── bantamparser/  # Bantam-specific parser configuration
└── main.go        # Test runner
```

## Learning Resources
- [Blog Post: Pratt Parsing: Expression Parsing Made Easy, by Bob Nystrom](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)
- [Original Bantam Repository](https://github.com/munificent/bantam)
- [Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)
