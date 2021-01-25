package testlogger

import (
	"github.com/rickb777/ech0/v2"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strconv"
	"sync"
)

// TestLogger captures log messages, organised by level: Infos, Warns, Errors and Panics.
// It deliberately ignores Debug level messages.
//
// Note that Fatal will call os.Exit so cannot usefully be tested.
type TestLogger struct {
	realLogger ech0.Zero
	Infos      TestLogEvents
	Warns      TestLogEvents
	Errors     TestLogEvents
	Panics     TestLogEvents
	mu         *sync.Mutex
	// note that debug messages are deliberately ignored
	// and fatal messages cannot be captured
}

var _ ech0.Zero = &TestLogger{}

func New(realLogger ech0.Zero) *TestLogger {
	return &TestLogger{realLogger: realLogger, mu: &sync.Mutex{}}
}

// NewWithConsoleLogger creates a new test logger with a wrapped console logger.
func NewWithConsoleLogger() *TestLogger {
	return New(ech0.Wrap(zerolog.New(zerolog.NewConsoleWriter())))
}

func (l *TestLogger) Log() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Log()
	}
	return &TestLogEvent{realEvent: ze} // will be discarded after use
}

func (l *TestLogger) Debug() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Debug()
	}
	return &TestLogEvent{realEvent: ze} // will be discarded after use
}

func (l *TestLogger) Info() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Info()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	first := &TestLogEvent{realEvent: ze}
	l.Infos = append(l.Infos, first)
	return first
}

func (l *TestLogger) Warn() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Warn()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	first := &TestLogEvent{realEvent: ze}
	l.Warns = append(l.Warns, first)
	return first
}

func (l *TestLogger) Error() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Error()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	first := &TestLogEvent{realEvent: ze}
	l.Errors = append(l.Errors, first)
	return first
}

func (l *TestLogger) Panic() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Panic()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	first := &TestLogEvent{realEvent: ze, done: func(s string) { panic(s) }}
	l.Panics = append(l.Panics, first)
	return first
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
// Therefore, this should be avoided during testing.
func (l *TestLogger) Fatal() ech0.ZeroEvent {
	var ze ech0.ZeroEvent
	if l.realLogger != nil {
		ze = l.realLogger.Fatal()
	}
	return &TestLogEvent{realEvent: ze, done: func(string) { os.Exit(1) }}
}

func (l *TestLogger) Err(err error) ech0.ZeroEvent {
	if err != nil {
		return l.Error().Err(err)
	}
	return l.Info()
}

func (l *TestLogger) WithLevel(level zerolog.Level) ech0.ZeroEvent {
	switch level {
	case zerolog.DebugLevel:
		return l.Debug()
	case zerolog.InfoLevel:
		return l.Info()
	case zerolog.WarnLevel:
		return l.Warn()
	case zerolog.ErrorLevel:
		return l.Error()

	// zerolog has odd behaviour for fatal and panic that avoids
	// the normal termination behaviour; it's not supported here.
	//case zerolog.FatalLevel:
	//	return l.newEvent(zerolog.FatalLevel, nil)
	case zerolog.PanicLevel:
		return l.Panic()
	case zerolog.NoLevel:
		return l.Log()

	case zerolog.Disabled:
		return nil
	default:
		panic("zerolog: WithLevel(): invalid level: " + strconv.Itoa(int(level)))
	}
}

func (l *TestLogger) Output(w io.Writer) ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.Output(w)
	}
	return l
}

func (l *TestLogger) Level(lvl zerolog.Level) ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.Level(lvl)
	}
	return l
}

func (l *TestLogger) Str(key, val string) ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.Str(key, val)
	}
	return l
}

func (l *TestLogger) Int(key string, val int) ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.Int(key, val)
	}
	return l
}

func (l *TestLogger) RawJSON(key string, val []byte) ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.RawJSON(key, val)
	}
	return l
}

func (l *TestLogger) Timestamp() ech0.Zero {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.realLogger != nil {
		l.realLogger = l.realLogger.Timestamp()
	}
	return l
}

//-------------------------------------------------------------------------------------------------

func (l *TestLogger) LastInfo() *TestLogEvent {
	return l.Infos.Last()
}

func (l *TestLogger) LastWarn() *TestLogEvent {
	return l.Warns.Last()
}

func (l *TestLogger) LastError() *TestLogEvent {
	return l.Errors.Last()
}

func (l *TestLogger) Reset() {
	l.Infos = nil
	l.Warns = nil
	l.Errors = nil
	l.Panics = nil
}
