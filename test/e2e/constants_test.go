package e2e

import (
	"bytes"
	"testing"

	"github.com/branislavlazic/rooster"
)

func TestConstants(t *testing.T) {
	content, err := rooster.LoadRCodeFile("./testdata/constants.rcode")
	if err != nil {
		t.Fatalf("cannot load an rcode file")
	}

	output := assertOutput(func(buf *bytes.Buffer) {
		rooster.RunVM(content, nil, buf)
	})
	expectedRes := "5\njohn doe\n55.3\n"
	if output != expectedRes {
		t.Fatalf("test - wrong output. expected=%s, got=%s", expectedRes, output)
	}
}
