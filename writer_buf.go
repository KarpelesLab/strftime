// +build !go1.10

package strftime

import "bytes"

func makeWriterBuf() strftimeWriterBuf {
	return &bytes.Buffer{}
}
