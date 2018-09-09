# strftime

strftime in Go, with locale support

[![GoDoc](https://godoc.org/github.com/MagicalTux/strftime?status.svg)](https://godoc.org/github.com/MagicalTux/strftime)

# Usage

```go
f := strftime.New(locale.French)
out.WriteString(f.Format(`pattern`, time.Now()));
```

# Description

There are many strftime libraries out there in Go, but locale support isn't commonly found. This library aims at providing a simple implementation with locale support.
