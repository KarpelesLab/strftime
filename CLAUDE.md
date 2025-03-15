# Build/Test Commands for strftime

```bash
# Build the project
go build -v

# Run all tests
make test

# Run a specific test
go test -run TestFunction

# Run benchmarks
go test -tags bench -benchmem -bench .

# Format code (requires goimports)
$(go env GOPATH)/bin/goimports -w -l .
```

# Code Style Guidelines

- **Module**: Use `github.com/KarpelesLab/strftime` for imports
- **Formatting**: Standard Go formatting with goimports
- **Naming**: CamelCase for public functions/structs, camelCase for private
- **Types**: Strong typing, especially with language.Tag for locales
- **Error Handling**: Return errors explicitly, don't use panic
- **Tests**: Use testify/assert package for test assertions
- **Documentation**: Document public functions with godoc-compatible comments
- **Functions**: Prefer methods that accept io.Writer to enable efficient output