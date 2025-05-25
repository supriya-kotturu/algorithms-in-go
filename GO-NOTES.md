# Go Topics Learned

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
middleware = redirect.MapHandler(routeMap, mux)
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

## Future Topics

- YAML parsing and handling
- More advanced middleware patterns
- Authentication and authorization
- Static file serving
- RESTful API design