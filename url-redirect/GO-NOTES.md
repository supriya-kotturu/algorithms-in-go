# Go Topics in URL Redirect

This project demonstrates the following Go concepts:

## Web Development

### HTTP Server Basics

- Using `net/http` package to create web servers
- `http.ListenAndServe()` function for starting HTTP servers
- Request and response handling with `http.ResponseWriter` and `*http.Request`

### Routing and Multiplexing

- `http.ServeMux` for request routing/multiplexing
- Registering handlers with `HandleFunc`
- Pattern matching for URL paths

### Middleware Pattern

- Creating middleware functions in Go
- Chaining middleware through handler wrapping
- Request flow through middleware to handlers
- Control flow: `ListenAndServe → Middleware → ServeMux → Handlers`

### HTTP Handlers

- `http.Handler` interface and `http.HandlerFunc` type
- Creating custom handlers
- Handler function signatures `func(w http.ResponseWriter, r *http.Request)`
- Redirecting requests with `http.Redirect()`

## Concurrency and Synchronization

### Mutex and RWMutex

- `sync.RWMutex` for read-write locking (multiple readers, exclusive writers)
- Protecting shared resources (maps) from concurrent access
- Lock patterns: `RLock()/RUnlock()` for reads and `Lock()/Unlock()` for writes
- Using `defer` with unlocking operations

## Data Formats

### YAML and JSON Handling

- Parsing YAML with `gopkg.in/yaml.v2`
- Parsing JSON with `encoding/json`
- Unmarshaling data into Go structs
- Error handling for parsing operations

## Package Organization

- Multiple files in the same package
- Function visibility across package files
- Package-level variables and constants
- Import paths and package naming

## Command Line Interface

- Flag parsing with the `flag` package
- Command-line arguments for configuration files

For more detailed Go notes, see the [main Go Notes](../GO-NOTES.md).
