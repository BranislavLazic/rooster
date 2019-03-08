package e2e

import (
	"bytes"
)

func assertOutput(f func(buf *bytes.Buffer)) string {
	buf := new(bytes.Buffer)
	f(buf)
	return buf.String()
}
