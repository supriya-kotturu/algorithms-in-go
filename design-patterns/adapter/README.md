# Adapter pattern

When we have two different variants(legacy and current) of the same entity, and need a consistent format to work together.
An adapter acts as a bridge to convert one interface to another that the target expects.

- Adapter `struct` implements the target interface and has a field which refers/implements the legacy interface.
- It wraps the legacy object and exposes the target interface methods, which internally delegates legacy methods and converts the data.

```go
type Adapter struct {
    legacy LegacySubscription  // <- legacy property
}

// Implements target interface
func (a *Adapter) getDetails() Details {
    // converts legacy format to target format
}
```

To run the `main` function:

```bash
go run *.go
```
