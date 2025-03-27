package logger_test

import (
	"testing"

	"github.com/your-username/logger"
)

func BenchmarkMapCopy(b *testing.B) {
	// Create a sample logger
	l := logger.NewLogger(logger.INFO, "app.log")

	// Sample context tags
	tags := map[string]string{"key1": "value1", "key2": "value2"}

	// Benchmark the WithTags method
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.WithTags(tags)
	}
}
