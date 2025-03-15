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
	// Fast path for common small values with width=2 (hours, minutes, seconds)
	if width == 2 && u < 100 {
		// Ensure we have enough space for two digits
		if cap(b)-len(b) < 2 {
			newB := make([]byte, len(b), len(b)+2)
			copy(newB, b)
			b = newB
		}

		if u < 10 {
			// Single digit with zero padding
			b = append(b, '0')
			b = append(b, byte('0'+u))
			return b
		} else {
			// Two digits
			tens := u / 10
			ones := u % 10
			b = append(b, byte('0'+tens))
			b = append(b, byte('0'+ones))
			return b
		}
	}

	// Regular path for other values or widths
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
	// Fast path for common case with width=2 (days, etc.)
	if width == 2 && u < 100 {
		// Ensure we have enough space
		if cap(b)-len(b) < 2 {
			newB := make([]byte, len(b), len(b)+2)
			copy(newB, b)
			b = newB
		}

		if u < 10 {
			// Single digit with space padding
			b = append(b, ' ')
			b = append(b, byte('0'+u))
			return b
		} else {
			// Two digits
			tens := u / 10
			ones := u % 10
			b = append(b, byte('0'+tens))
			b = append(b, byte('0'+ones))
			return b
		}
	}

	// Regular path for other values or widths
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
