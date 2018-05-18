package lexer

import (
	"testing"

	"github.com/BranislavLazic/rooster/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `ICONST 7 # Push number 7 on the stack
	ICONST 10 # Push number 10 on the stack
	IADD
	HALT`
	lexer := NewLexer(input)
	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{token.ICONST, "ICONST"},
		{token.INT, "7"},
		{token.COMMENT, "# Push number 7 on the stack"},
		{token.EOL, ""},
		{token.ICONST, "ICONST"},
		{token.INT, "10"},
		{token.COMMENT, "# Push number 10 on the stack"},
		{token.EOL, ""},
		{token.IADD, "IADD"},
		{token.HALT, "HALT"},
		{token.EOF, ""},
	}

	for i, tokenType := range tests {
		tok := lexer.NextToken()
		if tok.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tokenType.expectedType, tok.Type)
		}

		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, tok.Literal)
		}
	}
}
