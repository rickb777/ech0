package ech0

import (
	"github.com/rs/zerolog"
	"io"
)

// Zero mimics a zerolog.Logger.
// It excludes Trace because it is a non-requirement here.
type Zero interface {
	// Log starts a new message with no level. Setting GlobalLevel to Disabled
	// will still disable events produced by this method.
	Log() ZeroEvent
	// Debug starts a new message with debug level.
	Debug() ZeroEvent
	// Info starts a new message with info level.
	Info() ZeroEvent
	// Warn starts a new message with warn level.
	Warn() ZeroEvent
	// Error starts a new message with error level.
	Error() ZeroEvent
	// Fatal starts a new message with fatal level; the program will terminate.
	Fatal() ZeroEvent
	// Panic starts a new message with panic level; there will be a panic.
	Panic() ZeroEvent
	// Err starts a new message with error level with err as a field if not nil or
	// with info level if err is nil.
	Err(err error) ZeroEvent
	// WithLevel starts a new message with level. Unlike Fatal and Panic
	// methods, WithLevel does not terminate the program or stop the ordinary
	// flow of a gourotine when used with their respective levels.
	//
	// You must call Msg on the returned event in order to send the event.
	WithLevel(level zerolog.Level) ZeroEvent

	// Output duplicates the current logger and sets w as its output.
	Output(w io.Writer) Zero
	// Level creates a child logger with the minimum accepted level set to level.
	Level(lvl zerolog.Level) Zero

	// Str creates a child logger with the field key and with val as a string to the logger context.
	Str(key, val string) Zero
	// Int creates a child logger with the field key and with val as an int to the logger context.
	Int(key string, val int) Zero
	// RawJSON creates a child logger with the field key with val as already encoded JSON to context.
	//
	// No sanity check is performed on b; it must not contain carriage returns and be valid JSON.
	RawJSON(key string, b []byte) Zero
	// Timestamp adds the current local time as UNIX timestamp to the logger context with the "time" key.
	// To customize the key name, change zerolog.TimestampFieldName.
	//
	// NOTE: It won't dedupe the "time" key if the internal context has one already.
	Timestamp() Zero
}

var _ Zero = &zeroFacade{}

type zeroFacade zerolog.Logger

// Wrap wraps an existing logger.
func Wrap(z zerolog.Logger) Zero {
	return (*zeroFacade)(&z)
}

// Log starts a new message with no level. Setting GlobalLevel to Disabled
// will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Log() ZeroEvent {
	return (*zeroEvent)(z.Zero().Log())
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Debug() ZeroEvent {
	return (*zeroEvent)(z.Zero().Debug())
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Info() ZeroEvent {
	return (*zeroEvent)(z.Zero().Info())
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Warn() ZeroEvent {
	return (*zeroEvent)(z.Zero().Warn())
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Error() ZeroEvent {
	return (*zeroEvent)(z.Zero().Error())
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Err(err error) ZeroEvent {
	return (*zeroEvent)(z.Zero().Err(err))
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Fatal() ZeroEvent {
	return (*zeroEvent)(z.Zero().Fatal())
}

// Panic starts a new message with panic level. The panic() function
// is called by the Msg method, which stops the ordinary flow of a goroutine.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Panic() ZeroEvent {
	return (*zeroEvent)(z.Zero().Panic())
}

// WithLevel starts a new message with level. Unlike Fatal and Panic
// methods, WithLevel does not terminate the program or stop the ordinary
// flow of a gourotine when used with their respective levels.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) WithLevel(level zerolog.Level) ZeroEvent {
	return (*zeroEvent)(z.Zero().WithLevel(level))
}

// Output duplicates the current logger and sets w as its output.
func (z *zeroFacade) Output(w io.Writer) Zero {
	return Wrap(z.Zero().Output(w))
}

// Level creates a child logger with the minimum accepted level set to level.
func (z *zeroFacade) Level(lvl zerolog.Level) Zero {
	return Wrap(z.Zero().Level(lvl))
}

// Str creates a child logger with the field key and with val as a string to the logger context.
func (z *zeroFacade) Str(key, val string) Zero {
	return Wrap(z.Zero().With().Str(key, val).Logger())
}

// Int creates a child logger with the field key and with val as an int to the logger context.
func (z *zeroFacade) Int(key string, val int) Zero {
	return Wrap(z.Zero().With().Int(key, val).Logger())
}

// RawJSON creates a child logger with the field key with val as already encoded JSON to context.
//
// No sanity check is performed on b; it must not contain carriage returns and be valid JSON.
func (z *zeroFacade) RawJSON(key string, val []byte) Zero {
	return Wrap(z.Zero().With().RawJSON(key, val).Logger())
}

// Timestamp creates a child logger and adds the current local time as UNIX timestamp to the
// logger context with the "time" key. To customize the key name, change zerolog.TimestampFieldName.
//
// NOTE: It won't dedupe the "time" key if the internal context has one already.
func (z *zeroFacade) Timestamp() Zero {
	return Wrap(z.Zero().With().Timestamp().Logger())
}

// Zero unwraps the actual logger.
func (z *zeroFacade) Zero() *zerolog.Logger {
	return (*zerolog.Logger)(z)
}
