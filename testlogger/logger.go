package testlogger

import (
	"github.com/rickb777/ech0"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strconv"
)

// TestLogger captures log messages, organised by level: Infos, Warns and Errors.
// It deliberately ignores Debug level messages.
// Note that Fatal will call os.Exit so cannot usefully be tested.
type TestLogger struct {
	Infos  []*TestLogEvent
	Warns  []*TestLogEvent
	Errors []*TestLogEvent
	Panics []*TestLogEvent
	// note that debug messages are deliberately ignored
	// and fatal messages cannot be captured
}

var _ ech0.Zero = &TestLogger{}

func (l *TestLogger) Debug() ech0.ZeroEvent {
	return &TestLogEvent{} // will be discarded after use
}

func (l *TestLogger) Info() ech0.ZeroEvent {
	first := &TestLogEvent{}
	l.Infos = append(l.Infos, first)
	return first
}

func (l *TestLogger) Warn() ech0.ZeroEvent {
	first := &TestLogEvent{}
	l.Warns = append(l.Warns, first)
	return first
}

func (l *TestLogger) Error() ech0.ZeroEvent {
	first := &TestLogEvent{}
	l.Errors = append(l.Errors, first)
	return first
}

func (l *TestLogger) Panic() ech0.ZeroEvent {
	return &TestLogEvent{done: func(s string) { panic(s) }}
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
// Therefore, this should be avoided during testing.
func (l *TestLogger) Fatal() ech0.ZeroEvent {
	return &TestLogEvent{done: func(string) { os.Exit(1) }}
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
	//case zerolog.PanicLevel:
	//	return l.newEvent(zerolog.PanicLevel, nil)
	//case zerolog.NoLevel:
	//	return l.Log()

	case zerolog.Disabled:
		return nil
	default:
		panic("zerolog: WithLevel(): invalid level: " + strconv.Itoa(int(level)))
	}
}

func (l *TestLogger) Output(w io.Writer) ech0.Zero {
	return l
}

func (l *TestLogger) Level(lvl zerolog.Level) ech0.Zero {
	return l
}

func (l *TestLogger) Str(key, val string) ech0.Zero {
	return l
}

func (l *TestLogger) Int(key string, val int) ech0.Zero {
	return l
}

func (l *TestLogger) RawJSON(key string, b []byte) ech0.Zero {
	return l
}

//-------------------------------------------------------------------------------------------------

func (l *TestLogger) LastInfo() *TestLogEvent {
	if len(l.Infos) == 0 {
		return nil
	}
	return l.Infos[len(l.Infos)-1].Next
}

func (l *TestLogger) LastWarn() *TestLogEvent {
	if len(l.Warns) == 0 {
		return nil
	}
	return l.Warns[len(l.Warns)-1].Next
}

func (l *TestLogger) LastError() *TestLogEvent {
	if len(l.Errors) == 0 {
		return nil
	}
	return l.Errors[len(l.Errors)-1].Next
}

func (l *TestLogger) Reset() {
	l.Infos = nil
	l.Warns = nil
	l.Errors = nil
}
