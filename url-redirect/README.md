# URL Shortener/Redirector

A simple URL redirection service built in Go that maps paths to destination URLs.

## Project Structure

This project demonstrates Go module organization with multiple packages:

```bash
url-redirect/
├── go.mod                 # Main module definition
├── main.go                # Entry point application
├── main/                  # Executable package
│   └── go.mod             # Module for main package
├── redirect/              # Package for redirection logic
│   ├── go.mod             # Module for redirect package
│   └── redirect.go        # Redirection handlers
├── ISSUES.md              # Common issues and solutions
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

### Setting Up Modules

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
- Any other path shows "Hello, world!"

## Troubleshooting

For common issues and their solutions, please see [ISSUES.md](./ISSUES.md).