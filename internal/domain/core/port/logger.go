package port

import "context"

type Logger interface {
	Trace(ctx context.Context, message string, fields ...any)
	Debug(ctx context.Context, message string, fields ...any)
	Info(ctx context.Context, message string, fields ...any)
	Warn(ctx context.Context, message string, fields ...any)
	Error(ctx context.Context, message string, fields ...any)
	Fatal(ctx context.Context, message string, fields ...any)
	Panic(ctx context.Context, message string, fields ...any)
}
