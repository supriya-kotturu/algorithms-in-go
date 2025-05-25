# Go Topics Reference

This document serves as a comprehensive reference for Go concepts demonstrated across projects in this repository.

## Project-Specific Topics

- [Quiz Game Topics](./quiz-game/GO-NOTES.md)
- [URL Redirect Topics](./url-redirect/GO-NOTES.md)

## Concurrency and Synchronization

### Mutex and RWMutex

- `sync.Mutex` for exclusive locking (one writer/reader at a time)
- `sync.RWMutex` for read-write locking (multiple readers, exclusive writers)
- Protecting shared resources from concurrent access
- Lock patterns: `Lock()/Unlock()` and `RLock()/RUnlock()`
- Using `defer` with unlocking operations

### Goroutines

- Lightweight concurrent execution units
- Starting goroutines with `go` keyword
- Asynchronous execution model

### Channels

- Communication between goroutines
- Sending and receiving values
- Blocking operations and synchronization
- Select statement for managing multiple channels

## File Operations

- Opening files with `os.Open()`
- Closing files with `defer file.Close()`
- Reading file contents
- Error handling for file operations
- CSV parsing with `encoding/csv`

## Package Organization

- Multiple files in the same package
- Function visibility across package files
- Package-level variables and constants
- Import paths and package naming

## Web Development

### HTTP Server Basics

- Using `net/http` package to create web servers
- `http.ListenAndServe()` function for starting HTTP servers
- Request and response handling with `http.ResponseWriter` and `*http.Request`

#### HTTP Request Flow

```go
HTTP Request
    │
    ▼
http.ListenAndServe(":8080", middleware)
    │
    ▼
middleware = redirect.JSONHandler(shortPathMap, jsonFilePath, middleware)
    │
    ├─── Is path in JSON map? ───Yes──► http.Redirect() ───► HTTP Redirect Response
    │                                                           (e.g., /github → https://github.com)
    │ No
    ▼
middleware = redirect.YAMLHandler(shortPathMap, yamlFilePath, middleware)
    │
    ├─── Is path in YAML map? ───Yes──► http.Redirect() ───► HTTP Redirect Response
    │                                                           (e.g., /yaml → https://yaml.org)
    │ No
    ▼
middleware = redirect.MapHandler(shortPathMap, mux)
    │
    ├─── Is path in routeMap? ───Yes──► http.Redirect() ───► HTTP Redirect Response
    │                                                           (e.g., /google → https://google.com)
    │ No
    ▼
mux.ServeHTTP(w, r)  [fallback handler]
    │
    ├─── Path = "/" ───► defaultHandler() ───► "Hello, world!" Response
    │
    ├─── Path = "/404" ───► errorHandler() ───► Error Response (400)
    │
    └─── No match ───► defaultHandler() ───► "Hello, world!" Response
                        (ServeMux routes unmatched paths to "/" handler)
```

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

### URL Handling

- Accessing URL paths with `r.URL.Path`
- URL redirection implementation
- Path-to-URL mapping

## Data Structures

- Using maps for path-to-URL mapping (`map[string]string`)

## Project Structure

- Organizing code into packages
- Importing custom packages

## Development Tools

### Hot Reloading

- **Air**: Most popular option - `go install github.com/cosmtrek/air@latest`
- **CompileDaemon**: Simple configuration - `go install github.com/githubnemo/CompileDaemon@latest`
- **Gin**: Lightweight utility (not the web framework) - `go install github.com/codegangsta/gin@latest`
- **Fresh**: Minimal auto-reloader - `go install github.com/pilu/fresh@latest`

## Data Formats

### YAML and JSON Handling

- Parsing YAML with `gopkg.in/yaml.v2`
- Parsing JSON with `encoding/json`
- Unmarshaling data into Go structs
- Error handling for parsing operations

## Future Topics

- Authentication and authorization
- Static file serving
- RESTful API design
- Database integration
- Testing HTTP handlers
