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

	// Basic trace (white color)
	tracer.Trace("Application started")

	// Trace with formatting (like fmt.Printf)
	count := 10
	tracer.Tracef("Processing %d items", count)
	tracer.Tracef("User %s logged in at %v", "John", time.Now().Format("15:04:05"))

	// Trace with custom color
	tracer.TraceWithColor("Processing request...", "lightblue")

	// Trace with custom color and formatting
	port := 8080
	tracer.TraceWithColorf("lightgreen", "Server started on port %d", port)
	tracer.TraceWithColorf("yellow", "Memory usage: %.2f MB", 125.67)

	// Error message (red color)
	tracer.Error("Failed to connect to database")

	// Session error (LightSalmon color)
	tracer.TraceSessionError("Session timeout occurred")

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
