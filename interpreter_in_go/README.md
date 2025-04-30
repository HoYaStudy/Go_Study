# Interpreter in Go

## REPL (Read, Eval, Print, Loop)

- Run
  > $ go run main.go

## Lexer (Lexical Analysis)

```
Source Code -> Token
```

### ✅ Test

> $ go test ./lexer

## Parser

```
Token -> AST (Abstract Syntax Tree)
```

Programming언어를 parsing할 때 2가지 전략이 있다.

- Top-Down 전략
- Bottom-Up 전략

  - **Recursive descent parsing**
    - Top-down operator precedence parser (Pratt parser)
      - http://crockford.com/javascript/tdop/tdop.htm
      - http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy
  - Earley parsing
    - 주어진 CFG(문맥 무관 문법)에 속한 문자열을 parsing 하는 algorithm
  - Predictive parsing

### ✅ Test

> $ go test ./parser

with Tracing

> $ go test -v -run <TEST_NAME> ./parser
