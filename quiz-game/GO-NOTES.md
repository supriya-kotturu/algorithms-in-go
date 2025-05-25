# Go Topics in Quiz Game

This project demonstrates the following Go concepts:

## Concurrency and Synchronization

### Goroutines

- Lightweight concurrent execution units
- Starting goroutines with `go` keyword for non-blocking I/O operations
- Asynchronous execution model for handling user input

### Channels

- Communication between goroutines
- Sending and receiving values for quiz answers
- Blocking operations and synchronization with timers
- Select statement for managing multiple channels (timer and answer channels)

## File Operations

- Opening files with `os.Open()`
- Closing files with `defer file.Close()`
- Reading file contents for quiz questions
- Error handling for file operations
- CSV parsing with `encoding/csv` package

## Command Line Interface

- Flag parsing with the `flag` package
- Command-line arguments for file paths and time limits
- Help text and usage information

## Error Handling

- Custom error creation with `errors.New()`
- Error propagation through return values
- Graceful error handling and user feedback

## Data Structures

- Custom types for questions and results
- Slices for storing collections of questions
- String manipulation with the `strings` package

For more detailed Go notes, see the [main Go Notes](../GO-NOTES.md).