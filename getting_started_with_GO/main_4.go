package main

import (
	"fmt"
)

//  BASIC INTERFACE DEFINITION
// An interface defines a CONTRACT (method signatures) without implementation
// type Logger interface {
//     Log(message string)
// }

// UNCOMMENT above to define a basic interface
type Logger interface {
	Log(message string)
}

//  MULTIPLE IMPLEMENTATIONS OF THE SAME INTERFACE
// In Go, a type implements an interface IMPLICITLY by having all required methods
// You don't need to explicitly declare "implements Logger"

// ConsoleLogger - logs to console
type ConsoleLogger struct {
	// prefix string // UNCOMMENT to add custom fields
}

// Method to satisfy Logger interface
func (l ConsoleLogger) Log(message string) {
	// fmt.Println("[CONSOLE]", message) // UNCOMMENT to see output
}

// FileLogger - logs to file (simulated)
type FileLogger struct {
	// filename string // UNCOMMENT to add custom fields
	// buffer   []string // UNCOMMENT to add a buffer
}

func (l FileLogger) Log(message string) {
	// fmt.Println("[FILE]", message) // UNCOMMENT to see output
	// l.buffer = append(l.buffer, message) // UNCOMMENT to buffer messages
}

// SqlLogger - logs to database (simulated)
type SqlLogger struct {
	// connectionString string // UNCOMMENT to add custom fields
}

func (l SqlLogger) Log(message string) {
	// fmt.Println("[SQL]", message) // UNCOMMENT to see output
}

// USING INTERFACES AS FUNCTION PARAMETERS
// This is the POWER of interfaces - decouple from concrete types

// Process function accepts ANY type that implements Logger
func process(logger Logger) {
	// logger.Log("Processing started...") // UNCOMMENT to see execution
	// logger.Log("Doing some work...") // UNCOMMENT to see execution
	// logger.Log("Processing complete!") // UNCOMMENT to see execution
}

// USING INTERFACES AS STRUCT FIELDS
// This allows dependency injection

type Server struct {
	// logger Logger // UNCOMMENT to use interface as field
	// name   string // UNCOMMENT
}

// Method on Server that uses the logger
func (s Server) Start() {
	// s.logger.Log("Server starting...") // UNCOMMENT
	// fmt.Println("Server name:", s.name) // UNCOMMENT
}

// INTERFACE COMPOSITION
// 
// Interfaces can be composed of other interfaces

// Reader interface - similar to io.Reader
type Reader interface {
	// Read(p []byte) (n int, err error) // UNCOMMENT
}

// Writer interface - similar to io.Writer
type Writer interface {
	// Write(p []byte) (n int, err error) // UNCOMMENT
}

// Closer interface - similar to io.Closer
type Closer interface {
	// Close() error // UNCOMMENT
}

// ReadWriter is composed of Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser is composed of Reader and Closer
type ReadCloser interface {
	Reader
	Closer
}

//  PRACTICAL EXAMPLE - Data Processing Pipeline
// 

type DataProcessor interface {
	Process(data string) string
}

// ConvertProcessor - converts data to uppercase
type ConvertProcessor struct{}

func (cp ConvertProcessor) Process(data string) string {
	// return strings.ToUpper(data) // UNCOMMENT
	return data
}

// TrimProcessor - removes whitespace
type TrimProcessor struct{}

func (tp TrimProcessor) Process(data string) string {
	// return strings.TrimSpace(data) // UNCOMMENT
	return data
}

// LowercaseProcessor - converts to lowercase
type LowercaseProcessor struct{}

func (lp LowercaseProcessor) Process(data string) string {
	// return strings.ToLower(data) // UNCOMMENT
	return data
}

// Pipeline function that uses ANY DataProcessor implementation
func runPipeline(processor DataProcessor, data string) string {
	// return processor.Process(data) // UNCOMMENT
	return data
}

// 
// SECTION 7: INTERFACE ADVANTAGES DEMO
// 
// Shows how you can switch implementations without changing code

func logWithDifferentImplementations() {
	// Create different logger implementations
	consoleLog := ConsoleLogger{}
	fileLog := FileLogger{}
	sqlLog := SqlLogger{}

	// All can be used as Logger interface
	// The function doesn't care which concrete type it is
	process(consoleLog) 
	process(fileLog) // UNCOMMENT
	process(sqlLog) // UNCOMMENT

	// Easy to test - you can create a MockLogger that implements Logger
	// mockLog := MockLogger{} // UNCOMMENT
	// process(mockLog) // UNCOMMENT
}

// ============================================================================
// SECTION 8: EMPTY INTERFACE
// ============================================================================
// interface{} is the empty interface - all types implement it!

func acceptAnything(i interface{}) {
	// fmt.Println("Accepted:", i) // UNCOMMENT
	// switch v := i.(type) { // UNCOMMENT - type assertion
	// case string:
	//     fmt.Println("String value:", v)
	// case int:
	//     fmt.Println("Integer value:", v)
	// case Logger:
	//     fmt.Println("This is a Logger!")
	//     v.Log("Hello from type assertion!")
	// default:
	//     fmt.Println("Unknown type")
	// }
}

// ============================================================================
// MAIN FUNCTION - UNCOMMENT SECTIONS TO RUN EXAMPLES
// ============================================================================

func main() {
	fmt.Println("=== Go Interfaces Learning Script ===\n")

	// Example 1: Basic interface usage
	fmt.Println("Example 1: Basic Interface Usage")
	// consoleLogger := ConsoleLogger{} // UNCOMMENT
	// process(consoleLogger) // UNCOMMENT
	fmt.Println("")

	// Example 2: Switch implementations
	fmt.Println("Example 2: Switching Implementations")
	// logWithDifferentImplementations() // UNCOMMENT
	fmt.Println("")

	// Example 3: Using interfaces in struct fields
	fmt.Println("Example 3: Interface as Struct Field")
	// server := Server{ // UNCOMMENT
	//     logger: ConsoleLogger{},
	//     name:   "WebServer",
	// }
	// server.Start() // UNCOMMENT
	fmt.Println("")

	// Example 4: Data processing pipeline
	fmt.Println("Example 4: Data Processing Pipeline")
	testData := "  Hello World  "
	uppercase := runPipeline(ConvertProcessor{}, testData) // UNCOMMENT
	fmt.Println("Uppercase:", uppercase) // UNCOMMENT
	//
	// lowercase := runPipeline(LowercaseProcessor{}, testData) // UNCOMMENT
	// fmt.Println("Lowercase:", lowercase) // UNCOMMENT
	//
	// trimmed := runPipeline(TrimProcessor{}, testData) // UNCOMMENT
	// fmt.Println("Trimmed:", trimmed) // UNCOMMENT
	fmt.Println("")

	// Example 5: Empty interface and type assertion
	fmt.Println("Example 5: Empty Interface & Type Assertion")
	// acceptAnything("hello") // UNCOMMENT
	// acceptAnything(42) // UNCOMMENT
	// acceptAnything(ConsoleLogger{}) // UNCOMMENT
	fmt.Println("")

	fmt.Println("Learning Tips")
	fmt.Println("1. Uncomment sections one by one and run to see output")
	fmt.Println("2. Try creating new implementations of Logger interface")
	fmt.Println("3. Notice how function signatures only depend on interfaces, not concrete types")
	fmt.Println("4. Interfaces enable loose coupling and easy testing")
	fmt.Println("5. Go's implicit interface satisfaction is very different from Java/C#")
}

/*
KEY CONCEPTS:

1. INTERFACE DEFINITION:
   - Defines METHOD SIGNATURES only (contract, no implementation)
   - type Logger interface { Log(message string) }

2. IMPLICIT SATISFACTION:
   - A type implements an interface by HAVING all required methods
   - No explicit "implements" keyword needed (unlike Java/C#)
   - This promotes small, focused interfaces

3. WHY USE INTERFACES?
   - Decouple code from specific implementations
   - Easy to test (create mock implementations)
   - Easy to switch implementations
   - Promote code reusability

4. INTERFACE COMPOSITION:
   - Interfaces can embed other interfaces
   - io.ReadCloser = io.Reader + io.Closer

5. EMPTY INTERFACE interface{}:
   - Implemented by ALL types
   - Use with type assertions: v.(Type)
   - Often used for flexible function parameters

6. GOOD PRACTICES:
   - Accept interfaces, return concrete types
   - Make interfaces small and focused
   - Interface names often end with "-er" (Reader, Writer, Logger)
   - Check standard library for interface patterns (io package)
*/
