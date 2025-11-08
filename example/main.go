package main

import (
	"time"

	"github.com/rphpires/tracer"
)

func main() {
	// Configure the tracer (optional)
	tracer.SetConfig(tracer.Config{
		ExecutableName: "MyApp",
		UserID:         "User123",
		MaxSize:        5_000_000, // 5MB
		MaxFiles:       15,
	})

	// Or just set the user ID
	tracer.SetUserID("User456")

	// Basic trace (white color) - single argument
	tracer.Trace("Application started")

	// Trace with multiple arguments (like fmt.Println)
	count := 10
	tracer.Trace("Processing", count, "items from database")
	tracer.Trace("User", "John", "logged in at", time.Now().Format("15:04:05"))

	// Trace with formatting (like fmt.Printf)
	tracer.Tracef("Processing %d items", count)
	tracer.Tracef("User %s logged in at %v", "John", time.Now().Format("15:04:05"))

	// Trace with custom color - multiple arguments
	port := 8080
	tracer.TraceWithColor("lightblue", "Server listening on port", port)
	tracer.TraceWithColor("lightgreen", "Connection", "established", "successfully")

	// Trace with custom color and formatting
	tracer.TraceWithColorf("lightgreen", "Server started on port %d", port)
	tracer.TraceWithColorf("yellow", "Memory usage: %.2f MB", 125.67)

	// Error message (red color) - multiple arguments
	tracer.Error("Failed to connect to database:", "timeout")
	tracer.Error("Connection error:", port, "unreachable")

	// Session error (LightSalmon color)
	tracer.TraceSessionError("Session timeout occurred")
	tracer.TraceSessionError("User", "session", "expired after", 30, "minutes")

	// Example with panic recovery
	riskyOperation()

	// More traces with different colors
	tracer.TraceWithColor("Success! Data saved", "lightgreen")
	tracer.TraceWithColor("Warning: Low disk space", "yellow")
	tracer.TraceWithColor("Info: Cache cleared", "cyan")

	tracer.Trace("Application finished")
}

func riskyOperation() {
	// Use defer to catch panics
	defer tracer.RecoverPanic()

	// Simulate some work
	tracer.Trace("Starting risky operation...")
	time.Sleep(100 * time.Millisecond)

	// Uncomment to test panic recovery
	// panic("Something went wrong!")

	tracer.Trace("Risky operation completed successfully")
}

// Example of a goroutine with panic recovery
func workerWithRecovery(id int) {
	defer tracer.RecoverPanic()

	tracer.TraceWithColorf("lightgreen", "Worker %d started", id)

	// Do some work...
	time.Sleep(100 * time.Millisecond)

	// If panic occurs, it will be logged
	if id == 3 {
		// Uncomment to test
		// panic(fmt.Sprintf("Worker %d encountered an error", id))
	}

	tracer.TraceWithColorf("lightgreen", "Worker %d finished", id)
}
