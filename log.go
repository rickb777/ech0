package ech0

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

var (
	zlvlToGomLvl = map[zerolog.Level]log.Lvl{
		zerolog.DebugLevel: log.DEBUG,
		zerolog.InfoLevel:  log.INFO,
		zerolog.WarnLevel:  log.WARN,
		zerolog.ErrorLevel: log.ERROR,
		zerolog.NoLevel:    log.OFF,
	}

	gomLvlToZlvl = map[log.Lvl]zerolog.Level{
		log.DEBUG: zerolog.DebugLevel,
		log.INFO:  zerolog.InfoLevel,
		log.WARN:  zerolog.WarnLevel,
		log.ERROR: zerolog.ErrorLevel,
		log.OFF:   zerolog.NoLevel,
	}
)

// Log wraps a zerolog.Logger to provide an `echo.Logger` implementation
type Log struct {
	prefix   string
	zl       zerolog.Logger
	out      io.Writer
	lvl      zerolog.Level
	callsite bool
}

// New returns a new Log instance with the given output.
// Pass in your own zerolog logger if required.
// No prefix is set; use SetPrefix if required.
func New(out io.Writer, z ...zerolog.Logger) *Log {
	var zl zerolog.Logger
	if len(z) == 0 {
		zl = zerolog.New(out).With().Timestamp().Logger()
	} else {
		zl = z[0]
	}
	return &Log{
		prefix: "",
		out:    out,
		zl:     zl,
		lvl:    zerolog.GlobalLevel(),
	}
}

func (l Log) logWithFields() zerolog.Logger {
	if !l.callsite {
		return l.zl
	}

	ll := l.zl
	_, f, no, ok := runtime.Caller(2)
	if ok {
		f = filepath.Base(f)
		ll = l.zl.With().Str("file", f).Int("line", no).Logger()
	}

	return ll
}

// Debug satisfies the echo.Logger interface
func (l Log) Debug(i ...interface{}) {
	ll := l.logWithFields()
	ll.Debug().Msg(fmt.Sprint(i...))
}

// Debugf satisfies the echo.Logger interface
func (l Log) Debugf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Debug().Msgf(format, i...)
}

// Debugj satisfies the echo.Logger interface
func (l Log) Debugj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Debug().Msg("")

}

// Info satisfies the echo.Logger interface
func (l Log) Info(i ...interface{}) {
	ll := l.logWithFields()
	ll.Info().Msg(fmt.Sprint(i...))
}

// Infof satisfies the echo.Logger interface
func (l Log) Infof(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Info().Msgf(format, i...)
}

// Infoj satisfies the echo.Logger interface
func (l Log) Infoj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Info().Msg("")

}

// Warn satisfies the echo.Logger interface
func (l Log) Warn(i ...interface{}) {
	ll := l.logWithFields()
	ll.Warn().Msg(fmt.Sprint(i...))
}

//Warnf satisfies the echo.Logger interface
func (l Log) Warnf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Warn().Msgf(format, i...)
}

// Warnj satisfies the echo.Logger interface
func (l Log) Warnj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Warn().Msg("")

}

// Error satisfies the echo.Logger interface
func (l Log) Error(i ...interface{}) {
	ll := l.logWithFields()
	ll.Error().Msg(fmt.Sprint(i...))
}

// Errorf satisfies the echo.Logger interface
func (l Log) Errorf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Error().Msgf(format, i...)
}

// Errorj satisfies the echo.Logger interface
func (l Log) Errorj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Error().Msg("")

}

// Fatal satisfies the echo.Logger interface
func (l Log) Fatal(i ...interface{}) {
	ll := l.logWithFields()
	ll.Fatal().Msg(fmt.Sprint(i...))
}

// Fatalf satisfies the echo.Logger interface
func (l Log) Fatalf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Fatal().Msgf(format, i...)
}

// Fatalj satisfies the echo.Logger interface
func (l Log) Fatalj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Fatal().Msg("")

}

// Panic satisfies the echo.Logger interface
func (l Log) Panic(i ...interface{}) {
	ll := l.logWithFields()
	ll.Panic().Msg(fmt.Sprint(i...))
}

// Panicf satisfies the echo.Logger interface
func (l Log) Panicf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.Panic().Msgf(format, i...)
}

// Panicj satisfies the echo.Logger interface
func (l Log) Panicj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}

	ll.Panic().Msg("")

}

// Print satisfies the echo.Logger interface
func (l Log) Print(i ...interface{}) {
	ll := l.logWithFields()
	ll.WithLevel(zerolog.NoLevel).Str("level", "-").Msg(fmt.Sprint(i...))
}

// Printf satisfies the echo.Logger interface
func (l Log) Printf(format string, i ...interface{}) {
	ll := l.logWithFields()
	ll.WithLevel(zerolog.NoLevel).Str("level", "-").Msgf(format, i...)
}

// Printj satisfies the echo.Logger interface
func (l Log) Printj(j log.JSON) {
	ll := l.logWithFields()
	for k, v := range j {
		j, _ := json.Marshal(v)
		ll = ll.With().RawJSON(k, j).Logger()
	}
	ll.WithLevel(zerolog.NoLevel).Str("level", "-").Msg("")
}

// Output satisfies the echo.Logger interface
func (l Log) Output() io.Writer {
	return l.out
}

// SetOutput satisfies the echo.Logger interface
func (l *Log) SetOutput(w io.Writer) {
	l.zl = l.zl.Output(w)
	l.out = w
}

// Level satisfies the echo.Logger interface
func (l Log) Level() log.Lvl {
	return zlvlToGomLvl[l.lvl]
}

// SetLevel satisfies the echo.Logger interface
func (l *Log) SetLevel(v log.Lvl) {
	zlvl := gomLvlToZlvl[v]
	l.zl = l.zl.Level(zlvl)
	l.lvl = zlvl
}

// Prefix satisfies the echo.Logger interface
func (l Log) Prefix() string {
	return l.prefix
}

// SetPrefix satisfies the echo.Logger interface.
func (l *Log) SetPrefix(prefix string) {
	// Have to create a brand-new logger, since zero-log doesn't dedup fields. "prefix" would appear twice in the log output.
	var z zerolog.Logger
	if prefix != "" {
		z = zerolog.New(l.Output()).With().Str("prefix", prefix).Timestamp().Logger().Level(l.lvl)
	} else {
		z = zerolog.New(l.Output()).With().Timestamp().Logger().Level(l.lvl)
	}
	l.zl = z
	l.prefix = prefix
}

// SetHeader satisfies the echo.Logger interface. It does nothing.
// Within `echo`, this is used to set the template for formatting log messages,
// but it is not required when using zerolog.
func (l *Log) SetHeader(h string) {
	// no-op
}

// SetCallsite controls whether file and line numbers are emitted with every
// log output. Set this true to enable these items.
func (l *Log) SetCallsite(enabled bool) {
	l.callsite = enabled
}

var _ echo.Logger = (*Log)(nil)
