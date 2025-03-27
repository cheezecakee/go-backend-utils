package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"runtime"
	"strings"
	"time"
)

// Severity of log entries
type LogLevel int

// Log Levels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	CRITICAL
)

type Logger struct {
	logger      *log.Logger
	logLevel    LogLevel
	contextTags map[string]string
}

type LogEntry struct {
	Timestamp  time.Time
	Level      LogLevel
	Message    string
	Context    map[string]any
	File       string
	Line       int
	StackTrace string
}

func NewLogger(logLevel LogLevel, logFilePath string) *Logger {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stderr, file)

	return &Logger{
		logger:      log.New(multiWriter, "", log.LstdFlags),
		logLevel:    logLevel,
		contextTags: make(map[string]string),
	}
}

func (l *Logger) WithTags(tags map[string]string) *Logger {
	newLogger := *l

	newLogger.contextTags = make(map[string]string)

	maps.Copy(newLogger.contextTags, l.contextTags)

	maps.Copy(newLogger.contextTags, tags)

	return &newLogger
}

func (l *Logger) log(level LogLevel, message string, ctx context.Context, context ...map[string]any) {
	// Skip logging if the level is below the configured threshold
	if level < l.logLevel {
		return
	}

	var stackTrace string
	if level == ERROR || level == CRITICAL {
		stackTrace = captureStackTrace()
	}

	// Capture caller information
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}

	logContext := make(map[string]any)

	requestID := GetRequestID(ctx)
	if requestID != "unknown" {
		logContext["request_id"] = requestID
	}

	// Add global context tags
	for k, v := range l.contextTags {
		logContext[k] = v
	}

	// Merge additional context if provided
	if len(context) > 0 {
		for _, ctx := range context {
			maps.Copy(logContext, ctx)
		}
	}

	// Format log entry
	entry := &LogEntry{
		Timestamp:  time.Now(),
		Level:      level,
		Message:    message,
		Context:    logContext,
		File:       file,
		Line:       line,
		StackTrace: stackTrace,
	}

	// Output formatting
	output := l.formatLogEntry(entry)
	l.logger.Println(output)
}

func (l *Logger) formatLogEntry(entry *LogEntry) string {
	// Sophisticated formatting with level, timestamp, message, and context
	levelStr := l.getLevelString(entry.Level)
	contextStr := l.formatContext(entry.Context)

	return fmt.Sprintf(
		"%s | %s | %s | %s | %s:%d %s",
		entry.Timestamp.Format(time.RFC3339),
		levelStr,
		entry.Message,
		contextStr,
		shortFilePath(entry.File),
		entry.Line,
		entry.StackTrace,
	)
}

func (l *Logger) formatContext(context map[string]any) string {
	var builder strings.Builder

	for key, value := range context {
		builder.WriteString(fmt.Sprintf("%s=%v", key, value))
	}

	return strings.TrimSpace(builder.String())
}

// Helper methods for formatting and utility
func (l *Logger) getLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// Convenience methos for different log levels
func (l *Logger) Debug(ctx context.Context, message string, context ...map[string]any) {
	l.log(DEBUG, message, ctx, context...)
}

func (l *Logger) Info(ctx context.Context, message string, context ...map[string]any) {
	l.log(INFO, message, ctx, context...)
}

func (l *Logger) Warn(ctx context.Context, message string, context ...map[string]any) {
	l.log(WARN, message, ctx, context...)
}

func (l *Logger) Error(ctx context.Context, message string, context ...map[string]any) {
	l.log(ERROR, message, ctx, context...)
}

func (l *Logger) Critical(ctx context.Context, message string, context ...map[string]any) {
	l.log(CRITICAL, message, ctx, context...)
}

// Utility to shorten file path
func shortFilePath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return strings.Join(parts[len(parts)-2:], "/")
	}
	return path
}

func captureStackTrace() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

// Global logger instace with default configuration
var Log = NewLogger(INFO, "app.log")
