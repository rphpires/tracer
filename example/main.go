package main

import (
	"fmt"
	"time"

	"github.com/yourusername/tracer"
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

	// Trace with custom color
	tracer.TraceWithColor("Processing request...", "lightblue")

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

	tracer.TraceWithColor(fmt.Sprintf("Worker %d started", id), "lightgreen")

	// Do some work...
	time.Sleep(100 * time.Millisecond)

	// If panic occurs, it will be logged
	if id == 3 {
		// Uncomment to test
		// panic(fmt.Sprintf("Worker %d encountered an error", id))
	}

	tracer.TraceWithColor(fmt.Sprintf("Worker %d finished", id), "lightgreen")
}
