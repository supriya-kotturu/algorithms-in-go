# Common Issues and Solutions

This document contains solutions to common issues encountered while working with the URL Shortener project.

## Module and Import Issues

### "Relative import paths are not supported in module mode"

When you see this error, it means you're trying to use a relative import like `./redirect` in a Go module.

**Solution**: Use the full module path for imports:

```go
import "github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect"
```

### "Module found but does not contain package"

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

### Working with local packages - stale imports

If you're using a custom package, you might run into issues when your local package is updated, but your main package (which uses it) still refers to the stale package.

**Solutions**:

1. **Use replace directive in go.mod**:

   ```go
   // In your main/go.mod file
   replace github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect => ../redirect
   ```

2. **Clean Go's module cache**:

   ```bash
   go clean -cache -modcache
   ```

3. **Force update the package**:

   ```bash
   go get -u github.com/supriya-kotturu/algorithms-in-go/url-redirect/redirect
   ```

## HTTP and Server Issues

### "Superfluous response.WriteHeader call"

This happens when you try to write to an HTTP response more than once.

**Solution**: Ensure you're not writing to the response before calling http.Redirect:

```go
// Incorrect - writes to response twice
w.Write([]byte("Some text"))
http.Redirect(w, r, url, http.StatusFound)

// Correct - only one write to the response
http.Redirect(w, r, url, http.StatusFound)
return // Important: return after redirect
```

### Browser caching issues with redirects

Sometimes browsers cache redirects, making it difficult to test changes.

**Solutions**:

1. **Use incognito/private browsing mode**
2. **Clear browser cache** or do a hard refresh (Ctrl+F5 or Cmd+Shift+R)
3. **Add a query parameter** to force a fresh request:

   ```bash
   http://localhost:8080/ddg?t=123456789
   ```

4. **Use curl** to test without browser caching:

   ```bash
   curl -v http://localhost:8080/ddg
   ```

5. **Change the port number** to ensure you're testing the latest version:

   ```go
   http.ListenAndServe(":8081", handler)
   ```
