package e2e

import (
	"bytes"
	"testing"

	"github.com/branislavlazic/rooster"
)

// TestUnconditionalJump should print only a single value (3) since
// first value (6) will be skipped
func TestUnconditionalJump(t *testing.T) {
	content, err := rooster.LoadRCodeFile("./testdata/jump/unconditional_jump.rcode")
	if err != nil {
		t.Fatalf("cannot load an rcode file")
	}

	output := assertOutput(func(buf *bytes.Buffer) {
		rooster.RunVM(content, nil, buf)
	})

	expectedRes := "3\n"
	if output != expectedRes {
		t.Fatalf("test - wrong output. expected=%s, got=%s", expectedRes, output)
	}
}
