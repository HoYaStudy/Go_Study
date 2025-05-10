package lexer

import "monkey/token"

type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	character    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			literal := string(ch) + string(l.character)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			literal := string(ch) + string(l.character)
			tok = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.character)
		}
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '*':
		tok = newToken(token.ASTERISK, l.character)
	case '/':
		tok = newToken(token.SLASH, l.character)
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case ':':
		tok = newToken(token.COLON, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.character)
	case ']':
		tok = newToken(token.RBRACKET, l.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.character) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.character)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.character = 0 // NUL (ASCII 0)
	} else {
		l.character = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currPosition
	for isLetter(l.character) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currPosition
	for isDigit(l.character) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) readString() string {
	position := l.currPosition + 1
	for {
		l.readChar()
		if l.character == '"' || l.character == 0 {
			break
		}
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
