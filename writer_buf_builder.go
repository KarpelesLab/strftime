// +build go1.10

package strftime

import "strings"

func makeWriterBuf() strftimeWriterBuf {
	return &strings.Builder{}
}
