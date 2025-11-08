// Package tracer provides a simple and elegant logging library for Go that generates
// colorful HTML log files with automatic rotation and filtering capabilities.
//
// Features:
//   - HTML-based log files with color support
//   - Automatic log rotation when files reach maximum size (default: 5MB)
//   - Keeps a maximum number of log files (default: 15 files)
//   - Thread-safe operations
//   - Interactive JavaScript-based log filtering
//   - Panic/exception recovery and logging
//   - Enable/disable tracing via configuration files
//
// Quick Start:
//
// To enable tracing, create one of these files in your application's directory:
//   - TraceEnable.txt
//   - TraceIntegraEnable.txt
//   - Trace.txt
//
// Basic usage:
//
//	package main
//
//	import "github.com/rphpires/tracer"
//
//	func main() {
//	    // Simple trace message
//	    tracer.Trace("Application started")
//
//	    // Multiple arguments (like fmt.Println)
//	    tracer.Trace("Processing", 10, "items from database")
//
//	    // Formatted trace
//	    tracer.Tracef("Processing %d items", 10)
//
//	    // Trace with custom color
//	    tracer.TraceWithColor("lightgreen", "Success!", "Data saved")
//
//	    // Formatted trace with color
//	    tracer.TraceWithColorf("yellow", "Memory: %.2f MB", 125.67)
//
//	    // Error logging
//	    tracer.Error("Connection failed:", "timeout")
//	}
//
// Configuration:
//
//	tracer.SetConfig(tracer.Config{
//	    ExecutableName: "MyApp",
//	    UserID:         "User123",
//	    MaxSize:        5_000_000,  // 5MB
//	    MaxFiles:       15,
//	})
//
// Panic Recovery:
//
//	func riskyOperation() {
//	    defer tracer.RecoverPanic()
//	    // Your code here
//	}
package tracer
