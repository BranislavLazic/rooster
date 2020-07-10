package e2e

import (
	"bytes"
	"testing"

	"github.com/branislavlazic/rooster"
)

func TestEquality(t *testing.T) {
	content, err := rooster.LoadRCodeFile("./testdata/equality.rcode")
	if err != nil {
		t.Fatalf("cannot load an rcode file")
	}

	output := assertOutput(func(buf *bytes.Buffer) {
		rooster.RunVM(content, nil, buf)
	})
	expectedRes := "0\n1\n"
	if output != expectedRes {
		t.Fatalf("test - wrong output. expected=%s, got=%s", expectedRes, output)
	}
}
