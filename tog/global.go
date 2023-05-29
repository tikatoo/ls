package tog

import (
	"math/rand"
	"strconv"

	"go.uber.org/zap"
)

// The top-level logger.
var logTopLevel *zap.Logger = nil
var logTopLevelExt *zap.Logger = nil

// Logger returns the top-level logger for the application.
// See the package documentation for details on
// the initial value of this function.
func Logger() *zap.Logger {
	return logTopLevelExt
}

// SetLogger sets logger as the new top-level logger.
// Note that this will not affect anything
// already derived from the old top-level logger,
// so you will likely want to call this function
// either from some init(),
// or very early on in main().
//
// This package makes some minor adjustments
// to logger before setting it as the top-level,
// so the actual logger returned by Logger
// will not be exactly the same as the argument to this function
// (although it will be derived from the passed logger).
// The return value of this function
// is the logger that would be returned by Logger.
func SetLogger(logger *zap.Logger) *zap.Logger {
	logTopLevel = logger.WithOptions(zap.AddCallerSkip(1))
	logTopLevelExt = logTopLevel.WithOptions(zap.AddCallerSkip(-1))
	return Logger()
}

func getID() string {
	// Uniqueness is only to distinguish entries in a log file.
	// These numbers are not used in a security context,
	// and there is no real need for the order to change between runs.
	uid := rand.Uint64()
	return strconv.FormatUint(uid, 32)
}

// ID returns a log field with a unique-ish value,
// to aid in distinguishing multiple log entries from one conceptual task.
//
// This value is NOT SECURELY GENERATED, may not be entirely unique,
// and may be generated completely deterministically.
// DO NOT rely on it for anything more than a general visual aid.
func ID() zap.Field {
	return zap.String("id", getID())
}

// NYI causes a panic, with the log message "not implemented".
// It is used to mark a function body that needs to be revisited.
func NYI() {
	Logger().Panic("not implemented", zap.StackSkip("at", 1))
}
