package logger

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	zlog := NewZerolog().
		WithWriter(os.Stdout).
		WithLevel(TraceLevel).
		WithCaller()
	// WithLayer("base").
	// WithLayer("sub")

	zlog.Trace("example trace msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)

	zlog.Debug("example debug msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)

	zlog.Info("example info msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)

	zlog.Warning("example warning msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)

	zlog.Error("example error msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)

	zlog.Fatal("example fatal msg",
		"key1", "value1",
		"key2", 2,
		"key3", true,
	)
}
