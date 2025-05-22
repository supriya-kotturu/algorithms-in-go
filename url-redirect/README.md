# URL Shortener/Redirector

A simple URL redirection service built in Go that maps paths to destination URLs.

## Project Structure

This project demonstrates Go module organization with multiple packages:

```bash
url-redirect/
├── go.mod                 # Main module definition
├── main.go                # Entry point application
├── main/                  # Executable package
│   └── go.mod             # Module for main package - not needed, just to demonstrate that we can import packages in same level
├── redirect/              # Package for redirection logic
│   ├── go.mod             # Module for redirect package
│   └── handler.go         # Redirection handlers
└── README.md              # This file
```

## Go Modules Setup

### Understanding Module Organization

This project uses Go modules to manage dependencies and package organization. We have three separate `go.mod` files:

1. **Root module** (`/url-redirect/go.mod`):
   - Defines the parent module
   - Module path: `github.com/supriya-kotturu/algorithms-in-go/url-redirect`

2. **Redirect package** (`/url-redirect/redirect/go.mod`):
   - Contains the redirection logic
   - Module path: `github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect`

3. **Main package** (`/url-redirect/main/go.mod`):
   - Contains the executable code
   - Module path: `github.com/supriya-kotturu/algorithms-in-go/url-redirect/main`

### Common Issues and Solutions

#### Issue: "Relative import paths are not supported in module mode"

When you see this error, it means you're trying to use a relative import like `./redirect` in a Go module.

**Solution**: Use the full module path for imports:

```go
import "github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
```

#### Issue: "Module found but does not contain package"

This happens when Go tries to find your package in a remote repository but can't locate it.

**Solutions**:

1. Make sure all intermediate packages are properly initialized with `go mod init`
2. Ensure your module path matches your repository structure
3. If working locally, use a simple module name:

   ```go
   go mod init url-redirect
   ```

   Then import using:

   ```go
   import "url-redirect/redirect"
   ```

### Setting Up Modules Correctly

1. **Initialize the root module**:

   ```bash
   cd url-redirect
   go mod init github.com/supriya-kotturu/algorithms-in-go/url-redirect
   ```

2. **Initialize the redirect package**:

   ```bash
   cd redirect
   go mod init github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect
   ```

3. **Initialize the main package**:

   ```bash
   cd main
   go mod init github.com/supriya-kotturu/algorithms-in-go/url-redirect/main
   ```

## Running the Application

```bash
go run main.go
```

Then visit:

- [http://localhost:8080/google](http://localhost:8080/google) - Redirects to Google
- [http://localhost:8080/ddg](http://localhost:8080/ddg) - Redirects to DuckDuckGo
- Any other path shows "Not found!"