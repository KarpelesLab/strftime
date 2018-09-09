package strftime

import "io"

func writeInt(w io.Writer, x int, width int) error {
	u := uint(x)
	var buf [20]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = '0'
	}

	if x < 0 {
		i--
		buf[i] = '-'
	}

	_, err := w.Write(buf[i:])
	return err
}

// optimized writeInt for unsigned values below 256
func writeUint8(w io.Writer, u uint8, width int) error {
	var buf [3]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = '0'
	}

	_, err := w.Write(buf[i:])
	return err
}

// version with space padding
func writeUint8Sp(w io.Writer, u uint8, width int) error {
	var buf [3]byte
	i := len(buf)

	// Assemble decimal in reverse order.
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		i--
		buf[i] = ' '
	}

	_, err := w.Write(buf[i:])
	return err
}
