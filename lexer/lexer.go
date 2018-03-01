package lexer

import "github.com/BranislavLazic/rooster/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	lineNumber   int
	ch           byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input, lineNumber: 1}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	// Comments
	case '#':
		tok.Literal = l.readComment()
		tok.Type = token.COMMENT
		tok.LineNumber = l.lineNumber
	case '\n':
		tok.Literal = ""
		tok.Type = token.EOL
		tok.LineNumber = l.lineNumber
		l.lineNumber++
	case '\r':
		tok.Literal = ""
		tok.Type = token.EOL
		tok.LineNumber = l.lineNumber
		l.lineNumber++
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.LineNumber = l.lineNumber
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readInstruction()
			tok.Type = token.LookupInstruction(tok.Literal)
			tok.LineNumber = l.lineNumber
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.LineNumber = l.lineNumber
			return tok
		} else {
			tok = token.Token{
				Type:       token.ILLEGAL,
				Literal:    string(l.ch),
				LineNumber: l.lineNumber,
			}
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) nextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readInstruction() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readComment() string {
	position := l.position
	for l.nextChar() != '\n' && l.nextChar() != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}
