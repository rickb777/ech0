package ech0

import (
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

// ZeroEvent mimics the zerolog Event struct.
type ZeroEvent interface {
	Send()
	Msg(string)
	Msgf(format string, v ...interface{})

	AnErr(key string, val error) ZeroEvent
	Bool(key string, val bool) ZeroEvent
	Bytes(key string, val []byte) ZeroEvent
	Dur(key string, val time.Duration) ZeroEvent
	Err(err error) ZeroEvent
	Int(key string, val int) ZeroEvent
	Int64(key string, val int64) ZeroEvent
	Interface(key string, val interface{}) ZeroEvent
	Str(key, val string) ZeroEvent
	Strs(key string, val []string) ZeroEvent
	Stringer(key string, val fmt.Stringer) ZeroEvent
	Time(key string, val time.Time) ZeroEvent
	Uint(key string, val uint) ZeroEvent
	Uint64(key string, val uint64) ZeroEvent
}

var _ ZeroEvent = &zeroEvent{}

type zeroEvent struct {
	ev *zerolog.Event
}

func (ze *zeroEvent) Send() {
	ze.ev.Send()
}

func (ze *zeroEvent) Msg(s string) {
	ze.ev.Msg(s)
}

func (ze *zeroEvent) Msgf(format string, v ...interface{}) {
	ze.ev.Msgf(format, v...)
}

//-------------------------------------------------------------------------------------------------

// AnErr adds the field key with serialized err to the *Event context.
// If err is nil, no field is added.
func (ze *zeroEvent) AnErr(key string, err error) ZeroEvent {
	return &zeroEvent{ev: ze.ev.AnErr(key, err)}
}

// Bool adds the field key with val as a bool to the *Event context.
func (ze *zeroEvent) Bool(key string, val bool) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Bool(key, val)}
}

// Bytes adds the field key with val as a string to the *Event context.
//
// Runes outside of normal ASCII ranges will be hex-encoded in the resulting
// JSON.
func (ze *zeroEvent) Bytes(key string, val []byte) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Bytes(key, val)}
}

// Dur adds the field key with duration d stored as zerolog.DurationFieldUnit.
// If zerolog.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func (ze *zeroEvent) Dur(key string, val time.Duration) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Dur(key, val)}
}

// Err adds the field "error" with serialized err to the *Event context.
// If err is nil, no field is added.
//
// To customize the key name, change zerolog.ErrorFieldName.
//
// If Stack() has been called before and zerolog.ErrorStackMarshaler is defined,
// the err is passed to ErrorStackMarshaler and the result is appended to the
// zerolog.ErrorStackFieldName.
func (ze *zeroEvent) Err(err error) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Err(err)}
}

// Int adds the field key with i as a int to the *Event context.
func (ze *zeroEvent) Int(key string, val int) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Int(key, val)}
}

// Hex adds the field key with val as a hex string to the *Event context.
func (ze *zeroEvent) Hex(key string, val []byte) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Hex(key, val)}
}

// Int64 adds the field key with i as a int64 to the *Event context.
func (ze *zeroEvent) Int64(key string, val int64) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Int64(key, val)}
}

// Interface adds the field key with i marshaled using reflection.
func (ze *zeroEvent) Interface(key string, val interface{}) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Interface(key, val)}
}

// Str adds the field key with val as a string to the *Event context.
func (ze *zeroEvent) Str(key string, val string) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Str(key, val)}
}

// Strs adds the field key with vals as a []string to the *Event context.
func (ze *zeroEvent) Strs(key string, val []string) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Strs(key, val)}
}

// Stringer adds the field key with val.String() (or null if val is nil) to the *Event context.
func (ze *zeroEvent) Stringer(key string, val fmt.Stringer) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Stringer(key, val)}
}

// Time adds the field key with t formated as string using zerolog.TimeFieldFormat.
func (ze *zeroEvent) Time(key string, val time.Time) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Time(key, val)}
}

// Uint adds the field key with i as a uint to the *Event context.
func (ze *zeroEvent) Uint(key string, val uint) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Uint(key, val)}
}

// Uint64 adds the field key with i as a uint to the *Event context.
func (ze *zeroEvent) Uint64(key string, val uint64) ZeroEvent {
	return &zeroEvent{ev: ze.ev.Uint64(key, val)}
}
