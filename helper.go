// Package strftime implements C-like strftime functionality.
package strftime

// appendInt appends a decimal representation of int x to byte slice b,
// ensuring minimum width with zero padding.
//
// Parameters:
//   - b: Destination byte slice to append to
//   - x: Integer value to format
//   - width: Minimum width of the result (zero padded if needed)
//
// Returns: The extended byte slice with the formatted integer appended
func appendInt(b []byte, x int, width int) []byte {
	u := uint(x)
	if x < 0 {
		b = append(b, '-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
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
		b = append(b, '0')
	}

	return append(b, buf[i:]...)
}

// appendInt64 appends a decimal representation of int64 x to byte slice b,
// ensuring minimum width with zero padding.
//
// Parameters:
//   - b: Destination byte slice to append to
//   - x: 64-bit integer value to format
//   - width: Minimum width of the result (zero padded if needed)
//
// Returns: The extended byte slice with the formatted integer appended
func appendInt64(b []byte, x int64, width int) []byte {
	u := uint64(x)
	if x < 0 {
		b = append(b, '-')
		u = uint64(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
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
		b = append(b, '0')
	}

	return append(b, buf[i:]...)
}

// appendUint8 appends a decimal representation of uint8 u to byte slice b,
// ensuring minimum width with zero padding.
//
// Parameters:
//   - b: Destination byte slice to append to
//   - u: 8-bit unsigned integer value to format
//   - width: Minimum width of the result (zero padded if needed)
//
// Returns: The extended byte slice with the formatted integer appended
func appendUint8(b []byte, u uint8, width int) []byte {
	// Assemble decimal in reverse order.
	var buf [3]byte
	i := len(buf)
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
		b = append(b, '0')
	}

	return append(b, buf[i:]...)
}

// appendUint8Sp appends a decimal representation of uint8 u to byte slice b,
// ensuring minimum width with space padding (instead of zeros).
//
// Parameters:
//   - b: Destination byte slice to append to
//   - u: 8-bit unsigned integer value to format
//   - width: Minimum width of the result (space padded if needed)
//
// Returns: The extended byte slice with the formatted integer appended
func appendUint8Sp(b []byte, u uint8, width int) []byte {
	// Assemble decimal in reverse order.
	var buf [3]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add space-padding.
	for w := len(buf) - i; w < width; w++ {
		b = append(b, ' ')
	}

	return append(b, buf[i:]...)
}
