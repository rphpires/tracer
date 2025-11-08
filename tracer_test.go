package tracer

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestSetConfig verifies that configuration can be set
func TestSetConfig(t *testing.T) {
	cfg := Config{
		ExecutableName: "TestApp",
		UserID:         "TestUser",
		MaxSize:        1000000,
		MaxFiles:       10,
	}

	SetConfig(cfg)

	if defaultConfig.ExecutableName != "TestApp" {
		t.Errorf("Expected ExecutableName to be 'TestApp', got '%s'", defaultConfig.ExecutableName)
	}

	if defaultConfig.UserID != "TestUser" {
		t.Errorf("Expected UserID to be 'TestUser', got '%s'", defaultConfig.UserID)
	}
}

// TestSetUserID verifies that user ID can be set
func TestSetUserID(t *testing.T) {
	SetUserID("User123")

	if defaultConfig.UserID != "User123" {
		t.Errorf("Expected UserID to be 'User123', got '%s'", defaultConfig.UserID)
	}
}

// TestTraceWithoutEnableFile verifies that tracing works without enable file
func TestTraceWithoutEnableFile(t *testing.T) {
	// Clean up any enable files
	os.Remove("TraceEnable.txt")
	os.Remove("TraceIntegraEnable.txt")
	os.Remove("Trace.txt")

	// This should not panic, just print to stdout
	Trace("Test message")
	Tracef("Test %s", "formatted")
	TraceWithColor("red", "Test colored")
	TraceWithColorf("blue", "Test %d", 123)
	Error("Test error")
	TraceSessionError("Test session error")
}

// TestTraceWithEnableFile verifies that tracing creates files when enabled
func TestTraceWithEnableFile(t *testing.T) {
	// Create enable file
	enableFile := "TraceEnable.txt"
	if err := os.WriteFile(enableFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove(enableFile)

	// Set test config
	SetConfig(Config{
		ExecutableName: "TestTrace",
		UserID:         "TestUser",
	})

	folderName := "Trace TestTrace"
	defer os.RemoveAll(folderName)

	// Test basic trace
	Trace("Test message")
	time.Sleep(100 * time.Millisecond)

	// Verify folder was created
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		t.Errorf("Expected folder '%s' to be created", folderName)
	}

	// Verify log file was created
	logFile := filepath.Join(folderName, "trace.html")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Errorf("Expected log file '%s' to be created", logFile)
	}

	// Read log file and verify content
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if len(content) == 0 {
		t.Error("Expected log file to have content")
	}

	// Check for HTML header
	if !contains(string(content), "<!DOCTYPE html>") {
		t.Error("Expected log file to contain HTML header")
	}

	// Check for test message
	if !contains(string(content), "Test message") {
		t.Error("Expected log file to contain 'Test message'")
	}
}

// TestTracef verifies formatted tracing
func TestTracef(t *testing.T) {
	enableFile := "TraceEnable.txt"
	if err := os.WriteFile(enableFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove(enableFile)

	SetConfig(Config{
		ExecutableName: "TestTracef",
	})

	folderName := "Trace TestTracef"
	defer os.RemoveAll(folderName)

	Tracef("Processing %d items", 42)
	time.Sleep(100 * time.Millisecond)

	logFile := filepath.Join(folderName, "trace.html")
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !contains(string(content), "Processing 42 items") {
		t.Error("Expected log file to contain 'Processing 42 items'")
	}
}

// TestTraceWithColorf verifies formatted colored tracing
func TestTraceWithColorf(t *testing.T) {
	enableFile := "TraceEnable.txt"
	if err := os.WriteFile(enableFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove(enableFile)

	SetConfig(Config{
		ExecutableName: "TestColorf",
	})

	folderName := "Trace TestColorf"
	defer os.RemoveAll(folderName)

	TraceWithColorf("lightgreen", "Server on port %d", 8080)
	time.Sleep(100 * time.Millisecond)

	logFile := filepath.Join(folderName, "trace.html")
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !contains(string(content), "Server on port 8080") {
		t.Error("Expected log file to contain 'Server on port 8080'")
	}

	if !contains(string(content), "lightgreen") {
		t.Error("Expected log file to contain 'lightgreen' color")
	}
}

// TestRecoverPanic verifies panic recovery
func TestRecoverPanic(t *testing.T) {
	enableFile := "TraceEnable.txt"
	if err := os.WriteFile(enableFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove(enableFile)

	SetConfig(Config{
		ExecutableName: "TestPanic",
	})

	folderName := "Trace TestPanic"
	defer os.RemoveAll(folderName)

	// This should not panic the test
	func() {
		defer RecoverPanic()
		panic("test panic")
	}()

	time.Sleep(100 * time.Millisecond)

	logFile := filepath.Join(folderName, "trace.html")
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !contains(string(content), "test panic") {
		t.Error("Expected log file to contain 'test panic'")
	}
}

// TestIsTraceEnabled verifies enable file detection
func TestIsTraceEnabled(t *testing.T) {
	// Clean up
	os.Remove("TraceEnable.txt")
	os.Remove("TraceIntegraEnable.txt")
	os.Remove("Trace.txt")

	if isTraceEnabled() {
		t.Error("Expected tracing to be disabled")
	}

	// Create enable file
	if err := os.WriteFile("TraceEnable.txt", []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove("TraceEnable.txt")

	if !isTraceEnabled() {
		t.Error("Expected tracing to be enabled")
	}
}

// TestTraceMultipleArgs verifies that multiple arguments work like fmt.Println
func TestTraceMultipleArgs(t *testing.T) {
	enableFile := "TraceEnable.txt"
	if err := os.WriteFile(enableFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create enable file: %v", err)
	}
	defer os.Remove(enableFile)

	SetConfig(Config{
		ExecutableName: "TestMultiArgs",
	})

	folderName := "Trace TestMultiArgs"
	defer os.RemoveAll(folderName)

	// Test Trace with multiple arguments
	Trace("Processing", 42, "items", "from", "database")
	time.Sleep(100 * time.Millisecond)

	logFile := filepath.Join(folderName, "trace.html")
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected := "Processing 42 items from database"
	if !contains(string(content), expected) {
		t.Errorf("Expected log file to contain '%s'", expected)
	}

	// Test TraceWithColor with multiple arguments
	TraceWithColor("lightgreen", "Server", "started", "on", "port", 8080)
	time.Sleep(100 * time.Millisecond)

	content, err = os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected2 := "Server started on port 8080"
	if !contains(string(content), expected2) {
		t.Errorf("Expected log file to contain '%s'", expected2)
	}

	// Test Error with multiple arguments
	Error("Connection", "failed:", "timeout")
	time.Sleep(100 * time.Millisecond)

	content, err = os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected3 := "** Connection failed: timeout"
	if !contains(string(content), expected3) {
		t.Errorf("Expected log file to contain '%s'", expected3)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) >= len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr)))
}
