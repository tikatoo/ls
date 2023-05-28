// Package tl implements a context-based logging API.
// The ctx argument to logging functions in this package
// will have a value that determines the destination Logger.
//
// This package is powered by go.uber.org/zap.
// See that package for more usage details.
//
// This package configures a default "top-level" logger
// (which is used for contexts that do not have a logger set),
// which has different behaviours depending on build tags.
// If the tag `tlProduction` is used at build time,
// the top-level logger is set to log at InfoLevel,
// and outputs to stderr in JSON format.
//
// Otherwise, if this tag is not specified,
// a logger more suitable for debugging is set up.
// This logger still outputs to stderr,
// but uses a more human-readable format,
// and defaults to logging at DebugLevel.
// This logger can also be tweaked with the environment variables
// LOG_LEVEL (which sets the logging level),
// and LOG_FILE (which adds a file path to output logs to,
// in addition to stderr).
//
// You can also manually set the top-level logger,
// using SetLogger.

package tl
