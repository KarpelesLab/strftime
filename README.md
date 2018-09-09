# strftime

strftime in Go, with BCP 47 language tags support via Go's golang.org/x/text/language

[![GoDoc](https://godoc.org/github.com/MagicalTux/strftime?status.svg)](https://godoc.org/github.com/MagicalTux/strftime)

# Usage

```go
f := strftime.New(language.French)
out.WriteString(f.Format(`pattern`, time.Now()));
```

# Description

There are many strftime libraries out there in Go, but locale support isn't commonly found. This library aims at providing a simple implementation with locale support.

## Internals

This strftime implementation works by writing to a io.Writer by default. If using Format instead of FormatF, it will write internally to a buffer, which is then returned as string.

Go provides two methods to accomplish this starting go 1.10: bytes.Buffer (available in earlier versions) or the new strings.Builder. In go 1.10 strings.Builder is still a bit slow, bug go 1.11 benchmark gets much closer to bytes.Buffer, with slightly less memory usage (208 B/op with strings.Builder vs 227 B/op with bytes.Buffer).

Because of this, it has been decided to stay with bytes.Buffer for now, but tests will have to be run again with new Go versions.
