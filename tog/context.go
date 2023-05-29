package tog

import (
	"context"

	"go.uber.org/zap"
)

// &ctxKeyLogger is the context key for the associated logger.
var ctxKeyLogger int

// set is an internal version of WithLogger.
// set doesn't modify AddCallerSkip,
// so is useful if you need to get+set in one go.
func set(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, &ctxKeyLogger, log)
}

// get is an internal version of FromContext.
// get doesn't modify AddCallerSkip,
// so is more appropriate for use within this package.
func get(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if log, ok := ctx.Value(&ctxKeyLogger).(*zap.Logger); ok && log != nil {
			return log
		}
	}

	return logTopLevel
}

// WithLogger returns a child context of ctx,
// which will use the given logger.
func WithLogger(ctx context.Context, log *zap.Logger) context.Context {
	return set(ctx, log.WithOptions(zap.AddCallerSkip(1)))
}

// FromContext gets the logger associated with ctx.
// If ctx is nil or has no associated logger,
// then this returns the top-level logger instead.
func FromContext(ctx context.Context) *zap.Logger {
	if ctx != nil {
		if log, ok := ctx.Value(&ctxKeyLogger).(*zap.Logger); ok && log != nil {
			return log.WithOptions(zap.AddCallerSkip(-1))
		}
	}

	return logTopLevelExt
}
