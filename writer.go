package strftime

import "io"

type strftimeWriter interface {
	io.Writer
	WriteByte(byte) error
	WriteRune(rune) (int, error)
	WriteString(s string) (n int, err error)
}
