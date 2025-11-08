# Tracer - HTML Logging Library for Go

A simple and elegant logging library for Go that generates colorful HTML log files with automatic rotation and filtering capabilities.

## Features

- HTML-based log files with color support
- Automatic log rotation when files reach maximum size (default: 5MB)
- Keeps a maximum number of log files (default: 15 files)
- Thread-safe operations
- Interactive JavaScript-based log filtering
- Panic/exception recovery and logging
- Enable/disable tracing via configuration files
- Customizable user ID and executable name

## Installation

```bash
go get github.com/rphpires/tracer
```

## Quick Start

### 1. Enable Tracing

Create one of these files in your application's directory to enable tracing:
- `TraceEnable.txt`
- `TraceIntegraEnable.txt`
- `Trace.txt`

### 2. Basic Usage

```go
package main

import "github.com/rphpires/tracer"

func main() {
    // Simple trace message (white color)
    tracer.Trace("Application started")

    // Trace with custom color
    tracer.TraceWithColor("Processing...", "lightblue")

    // Error message (red color)
    tracer.Error("Something went wrong")

    // Session error (LightSalmon color)
    tracer.TraceSessionError("Session timeout")
}
```

## Configuration

### Set User ID

```go
tracer.SetUserID("User123")
```

### Custom Configuration

```go
tracer.SetConfig(tracer.Config{
    ExecutableName: "MyApp",       // Folder name will be "Trace MyApp"
    UserID:         "User123",     // Appears in log entries
    MaxSize:        5_000_000,     // 5MB per file
    MaxFiles:       15,            // Keep maximum 15 files
})
```

## API Reference

### Main Functions

#### `Trace(message string)`
Logs a message with white color.

```go
tracer.Trace("Operation completed")
```

#### `Tracef(format string, a ...any)`
Logs a formatted message with white color (like `fmt.Printf`).

```go
tracer.Tracef("Processing %d items", 10)
tracer.Tracef("User %s logged in at %v", "John", time.Now())
```

#### `TraceWithColor(message, color string)`
Logs a message with a specified HTML color.

```go
tracer.TraceWithColor("Warning: Low memory", "yellow")
tracer.TraceWithColor("Success!", "lightgreen")
tracer.TraceWithColor("Processing...", "cyan")
```

#### `TraceWithColorf(color string, format string, a ...any)`
Logs a formatted message with a specified HTML color (like `fmt.Printf`).

```go
tracer.TraceWithColorf("lightgreen", "Server started on port %d", 8080)
tracer.TraceWithColorf("yellow", "Memory usage: %.2f MB", 125.67)
tracer.TraceWithColorf("cyan", "Connected to %s:%d", "localhost", 5432)
```

Supported colors include: `white`, `red`, `lightgreen`, `lightblue`, `yellow`, `cyan`, `orange`, `pink`, `LightSalmon`, etc. Any valid HTML color name or hex code works.

#### `Error(message string)`
Logs an error message in red color with "**" prefix.

```go
tracer.Error("Database connection failed")
```

#### `TraceSessionError(message string)`
Logs a session error message in LightSalmon color with "**" prefix.

```go
tracer.TraceSessionError("Session expired")
```

#### `ReportException(err interface{})`
Logs an exception/panic with stack trace in red color.

```go
if err != nil {
    tracer.ReportException(err)
}
```

#### `RecoverPanic()`
Use with `defer` to catch and log panics automatically.

```go
func riskyOperation() {
    defer tracer.RecoverPanic()

    // Your code here
    // If panic occurs, it will be logged
}
```

### Configuration Functions

#### `SetUserID(userID string)`
Sets the user ID that appears in log entries.

```go
tracer.SetUserID("User123")
```

#### `SetConfig(cfg Config)`
Sets the complete configuration.

```go
tracer.SetConfig(tracer.Config{
    ExecutableName: "MyApp",
    UserID:         "User123",
    MaxSize:        10_000_000,  // 10MB
    MaxFiles:       20,
})
```

## Log File Structure

Logs are stored in a folder named `Trace [ExecutableName]` (default: `Trace Integra`).

The main log file is `trace.html`, which rotates when it reaches the maximum size. Rotated files are named with timestamps:
- `2024-11-08_14_30_45_trace.html`

## Interactive Log Filtering

Open any `.html` log file in a web browser and press **L** to activate the filter dialog. You can use regular expressions to filter log entries.

### Filter Examples:
- `CheckFirmwareUpdate` - Show only lines containing this text
- `ID=1 |ID=2` - Show lines with ID=1 or ID=2
- `2024-11-08 16:.*(ID=1 |ID=2 )` - Combine timestamp and ID filters

## Advanced Examples

### Using with Goroutines

```go
func worker(id int) {
    defer tracer.RecoverPanic()

    tracer.TraceWithColor(fmt.Sprintf("Worker %d started", id), "lightgreen")

    // Do work...

    tracer.TraceWithColor(fmt.Sprintf("Worker %d finished", id), "lightgreen")
}

func main() {
    for i := 0; i < 5; i++ {
        go worker(i)
    }
    time.Sleep(time.Second)
}
```

### Error Handling

```go
func connectDatabase() error {
    defer tracer.RecoverPanic()

    tracer.Trace("Connecting to database...")

    conn, err := db.Connect("localhost:5432")
    if err != nil {
        tracer.Error(fmt.Sprintf("Failed to connect: %v", err))
        return err
    }

    tracer.TraceWithColor("Database connected successfully", "lightgreen")
    return nil
}
```

### Conditional Tracing

```go
func processRequest(req Request) {
    tracer.Trace("Processing request: " + req.ID)

    if req.Priority == "high" {
        tracer.TraceWithColor("High priority request", "orange")
    }

    // Process...

    if err := process(req); err != nil {
        tracer.Error(fmt.Sprintf("Request failed: %v", err))
    } else {
        tracer.TraceWithColor("Request completed", "lightgreen")
    }
}
```

## Thread Safety

All functions are thread-safe and can be called from multiple goroutines simultaneously.

## Disabling Tracing

Simply remove or rename the enable files:
- `TraceEnable.txt`
- `TraceIntegraEnable.txt`
- `Trace.txt`

When disabled, trace calls only print to stdout but don't write to files.

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
