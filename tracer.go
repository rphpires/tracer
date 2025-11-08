package tracer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	logMaxSize = 5_000_000 // 5MB
	maxFiles   = 15
)

var (
	globalMutex sync.Mutex
)

const htmlPageHeader = `<!DOCTYPE html>
<meta content="text/html;charset=utf-8" http-equiv="Content-Type">
<script>
var original_html = null;
var filter = '';
function filter_log()
{
    document.body.style.cursor = 'wait';
    if (original_html == null) {
        original_html = document.body.innerHTML;
    }
    if (filter == '') {
        document.body.innerHTML = original_html;
    } else {
        l = original_html.split("<br>");
        var pattern = new RegExp(".*" + filter.replace('"', '\\"') + ".*", "i");
        final_html = '';
        for(var i=0; i<l.length; i++){
            if (pattern.test(l[i]))
                final_html += l[i] + '<br>';
        }
        document.body.innerHTML = final_html;
    }
    document.body.style.cursor = 'default';
}

document.onkeydown = function(event) {
    if (event.keyCode == 76) {
        var ret = prompt("Enter the filter regular expression. Examples:\\n\\n\
CheckFirmwareUpdate'\\n\\n'ID=1 |ID=2 \\n\\nID=2 .*Got message\\n\\n2012-08-31 16:.*(ID=1 |ID=2 )\\n\\n", filter);
        if (ret != null) {
            filter = ret;
            filter_log();
        }
        return false;
    }
}
</script>
<STYLE TYPE="text/css">
<!--
BODY
{
  color:white;
  background-color:black;
  font-family:monospace, sans-serif;
}
-->
</STYLE>
<body bgcolor="black" text="white">
<font color="white">`

// LogFile represents a log file with rotation capabilities
type LogFile struct {
	filename    string
	maxSize     int64
	currentSize int64
	file        *os.File
	mutex       sync.Mutex
}

// Config holds the tracer configuration
type Config struct {
	ExecutableName string
	UserID         string
	MaxSize        int64
	MaxFiles       int
}

var defaultConfig = Config{
	ExecutableName: "Integra",
	UserID:         "",
	MaxSize:        logMaxSize,
	MaxFiles:       maxFiles,
}

// SetConfig allows customization of the tracer configuration
func SetConfig(cfg Config) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	if cfg.MaxSize > 0 {
		defaultConfig.MaxSize = cfg.MaxSize
	}
	if cfg.MaxFiles > 0 {
		defaultConfig.MaxFiles = cfg.MaxFiles
	}
	if cfg.ExecutableName != "" {
		defaultConfig.ExecutableName = cfg.ExecutableName
	}
	if cfg.UserID != "" {
		defaultConfig.UserID = cfg.UserID
	}
}

// SetUserID sets the user ID that will appear in log entries
func SetUserID(userID string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	defaultConfig.UserID = userID
}

// NewLogFile creates a new LogFile instance
func newLogFile(filename string, maxSize int64) *LogFile {
	currentSize := int64(0)
	if info, err := os.Stat(filename); err == nil {
		currentSize = info.Size()
	}

	return &LogFile{
		filename:    filename,
		maxSize:     maxSize,
		currentSize: currentSize,
	}
}

func (lf *LogFile) openFile() error {
	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	if lf.file != nil {
		return nil
	}

	file, err := os.OpenFile(lf.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	lf.file = file
	return nil
}

func (lf *LogFile) closeFile() {
	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	if lf.file != nil {
		lf.file.Close()
		lf.file = nil
	}
}

func (lf *LogFile) rotateFile() error {
	lf.closeFile()

	currentDate := time.Now().Format("2006-01-02_15_04_05")
	dir := filepath.Dir(lf.filename)
	base := filepath.Base(lf.filename)
	newFilename := filepath.Join(dir, fmt.Sprintf("%s_%s", currentDate, base))

	if err := os.Rename(lf.filename, newFilename); err != nil {
		return err
	}

	lf.currentSize = 0
	return lf.openFile()
}

func (lf *LogFile) write(data string) error {
	if lf.file == nil {
		if err := lf.openFile(); err != nil {
			return err
		}
	}

	dataBytes := []byte(data)
	dataLen := int64(len(dataBytes))

	if lf.currentSize+dataLen > lf.maxSize {
		if err := lf.rotateFile(); err != nil {
			return err
		}
	}

	lf.mutex.Lock()
	defer lf.mutex.Unlock()

	n, err := lf.file.WriteString(data)
	if err != nil {
		return err
	}

	lf.currentSize += int64(n)
	return nil
}

func (lf *LogFile) close() {
	lf.closeFile()
	lf.currentSize = 0
}

func createHTMLLogFile(logFilename string) error {
	file, err := os.Create(logFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(htmlPageHeader)
	return err
}

func getLogFiles(folderName string) ([]string, error) {
	pattern := filepath.Join(folderName, "*.html")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// Sort by creation time
	sort.Slice(files, func(i, j int) bool {
		info1, err1 := os.Stat(files[i])
		info2, err2 := os.Stat(files[j])
		if err1 != nil || err2 != nil {
			return false
		}
		return info1.ModTime().Before(info2.ModTime())
	})

	return files, nil
}

func removeOldestLogFile(folderName string) error {
	logFiles, err := getLogFiles(folderName)
	if err != nil {
		return err
	}

	if len(logFiles) >= defaultConfig.MaxFiles {
		oldestFile := logFiles[0]
		return os.Remove(oldestFile)
	}

	return nil
}

func isTraceEnabled() bool {
	enableFiles := []string{"TraceEnable.txt", "TraceIntegraEnable.txt", "Trace.txt"}
	for _, file := range enableFiles {
		if _, err := os.Stat(file); err == nil {
			return true
		}
	}
	return false
}

// Trace writes a message to the trace log with white color
func Trace(message string) {
	TraceWithColor(message, "white")
}

// Tracef writes a formatted message to the trace log with white color (like fmt.Printf)
func Tracef(format string, a ...any) {
	TraceWithColor(fmt.Sprintf(format, a...), "white")
}

// TraceWithColor writes a message to the trace log with a specified color
func TraceWithColor(message, color string) {
	fmt.Println(message)

	if !isTraceEnabled() {
		return
	}

	globalMutex.Lock()
	executableName := defaultConfig.ExecutableName
	userID := defaultConfig.UserID
	globalMutex.Unlock()

	folderName := "Trace " + executableName

	// Create folder if it doesn't exist
	if err := os.MkdirAll(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder: %v\n", err)
		return
	}

	logFilename := filepath.Join(folderName, "trace.html")

	// Create HTML log file if it doesn't exist
	if _, err := os.Stat(logFilename); os.IsNotExist(err) {
		if err := createHTMLLogFile(logFilename); err != nil {
			fmt.Printf("Error creating log file: %v\n", err)
			return
		}
	}

	// Remove oldest log file if needed
	if err := removeOldestLogFile(folderName); err != nil {
		fmt.Printf("Error removing old log files: %v\n", err)
	}

	logFile := newLogFile(logFilename, defaultConfig.MaxSize)

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	userIDPart := ""
	if userID != "" {
		userIDPart = userID + " - "
	}
	logEntry := fmt.Sprintf("\n<br></font><font color=\"%s\">%s - %s%s", color, timestamp, userIDPart, message)

	if err := logFile.write(logEntry); err != nil {
		logFile.close()
		currentDate := time.Now().Format("2006-01-02_15_04_05")
		newFilename := filepath.Join(folderName, fmt.Sprintf("%s - trace.html", currentDate))

		if err := os.Rename(logFilename, newFilename); err != nil {
			fmt.Printf("Error renaming log file: %v\n", err)
			return
		}

		if err := createHTMLLogFile(logFilename); err != nil {
			fmt.Printf("Error creating new log file: %v\n", err)
			return
		}

		logFile = newLogFile(logFilename, defaultConfig.MaxSize)
		logFile.write(logEntry)
	}
}

// TraceWithColorf writes a formatted message to the trace log with a specified color (like fmt.Printf)
func TraceWithColorf(color string, format string, a ...any) {
	TraceWithColor(fmt.Sprintf(format, a...), color)
}

// ReportException reports a panic/exception with stack trace
func ReportException(err interface{}) {
	stackTrace := string(debug.Stack())

	// Clean up stack trace for HTML
	stackTrace = strings.ReplaceAll(stackTrace, "<", "&lt;")
	stackTrace = strings.ReplaceAll(stackTrace, ">", "&gt;")

	TraceWithColor(fmt.Sprintf("Bypassing exception (%v)", err), "red")
	TraceWithColor(fmt.Sprintf("**** Exception: <code>%s</code>", stackTrace), "red")
}

// Error writes an error message to the trace log in red
func Error(message string) {
	TraceWithColor(fmt.Sprintf("** %s", message), "red")
}

// TraceSessionError writes a session error message to the trace log in LightSalmon color
func TraceSessionError(message string) {
	TraceWithColor(fmt.Sprintf("** %s", message), "LightSalmon")
}

// RecoverPanic should be used with defer to catch panics and log them
func RecoverPanic() {
	if r := recover(); r != nil {
		ReportException(r)
	}
}
