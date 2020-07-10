package e2e

import (
	"bytes"
	"testing"

	"github.com/branislavlazic/rooster"
)

func TestGlobalVars(t *testing.T) {
	content, err := rooster.LoadRCodeFile("./testdata/global_vars.rcode")
	if err != nil {
		t.Fatalf("cannot load an rcode file")
	}

	output := assertOutput(func(buf *bytes.Buffer) {
		rooster.RunVM(content, nil, buf)
	})
	expectedRes := "42\n42\n3\n"
	if output != expectedRes {
		t.Fatalf("test - wrong output. expected=%s, got=%s", expectedRes, output)
	}
}
