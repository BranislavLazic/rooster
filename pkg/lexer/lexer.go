package lexer

import (
	"github.com/branislavlazic/rooster/pkg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	lineNumber   int
	index        int
	ch           byte
}

// NewLexer returns an initialized Lexer value
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input, lineNumber: 1, index: -1}
	lexer.readChar()
	return lexer
}

// NextToken is returning a token lexing a given sequence of characters
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	// Comments
	case '#':
		tok.Literal = l.readComment()
		tok.Type = token.COMMENT
		tok.LineNumber = l.lineNumber
	case '"':
		tok.Literal = l.readString()
		tok.Type = token.STRING
		tok.LineNumber = l.lineNumber
		l.index++
		tok.Index = l.index
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
		if isUppercaseLetter(l.ch) {
			tok.Literal = l.readInstruction()
			tok.Type = token.LookupInstruction(tok.Literal)
			tok.LineNumber = l.lineNumber
			l.index++
			tok.Index = l.index
			return tok
		} else if isLowercaseLetter(l.ch) {
			l.index++
			tok = l.readLabelToken()
		} else if isDigit(l.ch) {
			tok.Literal, tok.Type = l.readNumber()
			tok.LineNumber = l.lineNumber
			l.index++
			tok.Index = l.index
			return tok
		} else {
			l.index++
			tok = token.Token{
				Type:       token.ILLEGAL,
				Literal:    string(l.ch),
				LineNumber: l.lineNumber,
				Index:      l.index,
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
	for isUppercaseLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readLabelToken() token.Token {
	var tokenType token.TokType
	tokenType = token.LabelName
	position := l.position
	for isLetter(l.ch) {
		if l.nextChar() == ':' {
			tokenType = token.LABEL
		}
		l.readChar()
	}
	literal := l.input[position:l.position]
	tok := token.Token{
		Type:       tokenType,
		Literal:    literal,
		LineNumber: l.lineNumber,
		Index:      l.index,
	}
	// Decrement the index if it is a label so it contains an index
	// of the next operand. Label is passed to the instruction list.
	// It only "carries" an index of the first operand in the procedure.
	if tokenType == token.LABEL {
		l.index--
	}
	return tok
}

func (l *Lexer) readComment() string {
	position := l.position
	for l.nextChar() != '\n' && l.nextChar() != 0 {
		l.readChar()
	}
	return l.input[position : l.position+1]
}

func (l *Lexer) readString() string {
	position := l.position
	for l.nextChar() != '"' && l.nextChar() != '\n' {
		l.readChar()
	}
	result := l.input[position+1 : l.position+1]
	l.position = l.readPosition
	l.readPosition++
	return result
}

// Read either an integer or a floating point number.
// Initially, the token type will be declared as INT.
// The token will be declared as FLOAT only if the sequence of characters
// contains ".". If the "." is repeated again, then the ILLEGAL token type
// will be returned. Every floating point number must begin with number.
// E.g: 0.5
// Values such as ".35" are not allowed.
func (l *Lexer) readNumber() (string, token.TokType) {
	var tokType token.TokType = token.INT
	dotRepeated := false
	position := l.position
	for isDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if !dotRepeated {
				tokType = token.FLOAT
				dotRepeated = true
			} else {
				return string(l.ch), token.ILLEGAL
			}
		}
		l.readChar()
	}
	return l.input[position:l.position], tokType
}

func isUppercaseLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}

func isLowercaseLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}
}
