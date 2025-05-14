# Quiz Game

A simple command-line quiz application written in Go.

## Concepts Learned

### 1. Dependency Injection

Dependency injection helps make code more testable by allowing dependencies to be replaced with mocks during testing.

```go
// Define an interface for dependencies
type QuizDependencies interface {
    parseCommand() (string, int, bool, error)
    getQuestions(string) ([]Question, error)
    play([]Question, int) (Result, error)
    showHelp()
    showResult(Result)
}

// Create a struct that uses the dependencies
type Quiz struct {
    deps QuizDependencies
}

// Constructor that accepts dependencies
func NewQuiz(deps QuizDependencies) *Quiz {
    if deps == nil {
        return &Quiz{deps: &DefaultQuizDeps{}}
    }
    return &Quiz{deps: deps}
}
```

### 2. Interface-Based Design

Interfaces define behavior contracts that can be implemented by different structs.

```go
// Default implementation
type DefaultQuizDeps struct{}

func (d *DefaultQuizDeps) parseCommand() (string, int, bool, error) {
    return parseCommand()
}

// Other methods...
```

### 3. Mocking for Tests

Creating mock implementations for testing without external dependencies.

```go
// Mock implementation for testing
type MockQuizDeps struct {
    ParseCommandCalled bool
    FilePath           string
    TimeLimit          int
    HelpFlag           bool
    // Other fields...
}

func (m *MockQuizDeps) parseCommand() (string, int, bool, error) {
    m.ParseCommandCalled = true
    return m.FilePath, m.TimeLimit, m.HelpFlag, m.ParseCommandError
}
```

### 4. Table-Driven Tests

Writing multiple test cases in a concise way.

```go
func TestSomething(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
    }{
        {"Case 1", "input1", "expected1"},
        {"Case 2", "input2", "expected2"},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := functionUnderTest(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %s, got %s", tc.expected, result)
            }
        })
    }
}
```

### 5. Concurrency with Goroutines and Channels

Using goroutines for non-blocking operations and channels for communication.

```go
func play(questions []Question, limit int) (Result, error) {
    timer := time.NewTimer(time.Duration(limit) * time.Second)
    ansCh := make(chan string)
    
    for _, q := range questions {
        fmt.Println(q.question)
        
        // Non-blocking input with goroutine
        go func() {
            var ans string
            fmt.Scanf("%s\n", &ans)
            ansCh <- ans
        }()
        
        // Select from multiple channels
        select {
        case <-timer.C:
            return result, nil
        case ans := <-ansCh:
            // Process answer
        }
    }
    
    return result, nil
}
```

### 6. Command-Line Flag Parsing

Using the flag package to parse command-line arguments.

```go
func parseCommand() (string, int, bool, error) {
    file := flag.String("file", "problem-set.csv", "Path to question bank")
    limit := flag.Int("limit", 30, "Time limit in seconds")
    help := flag.Bool("help", false, "Show help")
    
    flag.Parse()
    
    // Use the parsed flags
    return *file, *limit, *help, nil
}
```

### 7. File I/O and CSV Parsing

Reading and parsing CSV files.

```go
func getQuestions(filePath string) ([]Question, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return []Question{}, err
    }
    defer file.Close()
    
    r := csv.NewReader(file)
    records, err := r.ReadAll()
    
    // Process records into questions
    questions := make([]Question, len(records))
    for idx, record := range records {
        questions[idx] = Question{
            id:       idx,
            question: strings.TrimSpace(record[0]),
            answer:   strings.TrimSpace(record[1]),
        }
    }
    
    return questions, nil
}
```

### 8. Error Handling

Proper error propagation and handling.

```go
questions, err := deps.getQuestions(filepath)
if err != nil {
    log.Fatalf("could not get questions: %v", err)
}
```

### 9. Struct Design

Creating appropriate data structures.

```go
type Question struct {
    id       int
    question string
    answer   string
}

type Result struct {
    totalQuestions int
    correct        int
    attempted      int
}
```