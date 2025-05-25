# Installing and Using Air for Go Hot Reloading

## Installation

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Verify installation
air -v
```

## Usage

Navigate to your Go project directory and run:

```bash
air
```

## Configuration (Optional)

You can create a custom configuration file:

```bash
# Generate default config
air init
```

This creates a `.air.toml` file that you can customize.

## Basic .air.toml Example

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "./tmp/main"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor"]
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_regex = ["_test\\.go"]
```
