package ech0

import (
	"github.com/rs/zerolog"
	"io"
)

// Zero mimics a zerolog.Logger.
// It excludes Trace because it is a non-requirement here.
type Zero interface {
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

	Output(w io.Writer) Zero
	Level(lvl zerolog.Level) Zero

	Str(key, val string) Zero
	Int(key string, val int) Zero
	RawJSON(key string, b []byte) Zero
}

var _ Zero = &zeroFacade{}

type zeroFacade zerolog.Logger

// Wrap wraps an existing logger.
func Wrap(z zerolog.Logger) Zero {
	return (*zeroFacade)(&z)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Debug() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Debug()}
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Info() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Info()}
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Warn() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Warn()}
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Error() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Error()}
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Err(err error) ZeroEvent {
	return &zeroEvent{ev: z.Zero().Err(err)}
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Fatal() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Fatal()}
}

// Panic starts a new message with panic level. The panic() function
// is called by the Msg method, which stops the ordinary flow of a goroutine.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) Panic() ZeroEvent {
	return &zeroEvent{ev: z.Zero().Panic()}
}

// WithLevel starts a new message with level. Unlike Fatal and Panic
// methods, WithLevel does not terminate the program or stop the ordinary
// flow of a gourotine when used with their respective levels.
//
// You must call Msg on the returned event in order to send the event.
func (z *zeroFacade) WithLevel(level zerolog.Level) ZeroEvent {
	return &zeroEvent{ev: z.Zero().WithLevel(level)}
}

// Zero unwraps the actual logger.
func (z *zeroFacade) Zero() *zerolog.Logger {
	return (*zerolog.Logger)(z)
}

func (z *zeroFacade) Output(w io.Writer) Zero {
	return Wrap(z.Zero().Output(w))
}

func (z *zeroFacade) Level(lvl zerolog.Level) Zero {
	return Wrap(z.Zero().Level(lvl))
}

func (z *zeroFacade) Str(key, val string) Zero {
	return Wrap(z.Zero().With().Str(key, val).Logger())
}

func (z *zeroFacade) Int(key string, val int) Zero {
	return Wrap(z.Zero().With().Int(key, val).Logger())
}

func (z *zeroFacade) RawJSON(key string, val []byte) Zero {
	return Wrap(z.Zero().With().RawJSON(key, val).Logger())
}

//func (z *zeroFacade) Timestamp() Zero {
//	return Wrap(z.Zero().With().Timestamp().Logger())
//}
