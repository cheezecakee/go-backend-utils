package logger_test

import (
	"testing"

	"github.com/cheezecakee/go-backend-utils/pkg/logger"
)

func BenchmarkMapCopy(b *testing.B) {
	// Create a sample logger
	l := logger.NewLogger(logger.INFO, "app.log")

	// Sample context tags
	tags := map[string]string{"key1": "value1", "key2": "value2"}

	// Benchmark the WithTags method
	b.ResetTimer()
	for b.Loop() {
		l.WithTags(tags)
	}
}
