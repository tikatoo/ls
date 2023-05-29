package tog

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log logs a message at the specified level.
// The destination logger is the same as what would be returned
// by a call to FromContext(ctx).
//
// See also zap.Logger.Log.
func Log(ctx context.Context, lvl zapcore.Level, msg string, fields ...zapcore.Field) {
	get(ctx).Log(lvl, msg, fields...)
}

// Debug logs a message at DebugLevel.
func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.DebugLevel, msg, fields...)
}

// Info logs a message at InfoLevel.
func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.InfoLevel, msg, fields...)
}

// Warn logs a message at WarnLevel.
func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.WarnLevel, msg, fields...)
}

// Error logs a message at ErrorLevel.
func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.ErrorLevel, msg, fields...)
}

// DPanic logs a message at DPanicLevel.
//
// If the logger is in development mode,
// it then panics (DPanic means "development panic").
// This is useful for catching errors that are recoverable,
// but shouldn't ever happen.
func DPanic(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.DPanicLevel, msg, fields...)
}

// Panic logs a message at PanicLevel.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.PanicLevel, msg, fields...)
}

// Fatal logs a message at FatalLevel.
//
// The logger then calls os.Exit(1),
// even if logging at FatalLevel is disabled.
func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	get(ctx).Log(zapcore.FatalLevel, msg, fields...)
}

// With creates a child logger with the given fields,
// then returns a child context associated with that child logger.
func With(ctx context.Context, fields ...zap.Field) context.Context {
	return set(ctx, get(ctx).With(fields...))
}

// WithOptions clones the logger that is returned by FromContext(ctx),
// applies the supplied options,
// then returns a child context associated with that logger.
func WithOptions(ctx context.Context, opts ...zap.Option) context.Context {
	return set(ctx, get(ctx).WithOptions(opts...))
}

// Sugar returns a logger with a more ergonomic API,
// derived from the logger returned by FromContext(ctx).
func Sugar(ctx context.Context) *zap.SugaredLogger {
	return FromContext(ctx).Sugar()
}
