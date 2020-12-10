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
	Hex(key string, val []byte) ZeroEvent
	Int(key string, val int) ZeroEvent
	Ints(key string, val []int) ZeroEvent
	Int64(key string, val int64) ZeroEvent
	Interface(key string, val interface{}) ZeroEvent
	Str(key, val string) ZeroEvent
	Strs(key string, val []string) ZeroEvent
	Stringer(key string, val fmt.Stringer) ZeroEvent
	Time(key string, val time.Time) ZeroEvent
	Uint(key string, val uint) ZeroEvent
	Uints(key string, val []uint) ZeroEvent
	Uint64(key string, val uint64) ZeroEvent
}

var _ ZeroEvent = &zeroEvent{}

type zeroEvent zerolog.Event

func (ze *zeroEvent) Send() {
	(*zerolog.Event)(ze).Send()
}

func (ze *zeroEvent) Msg(s string) {
	(*zerolog.Event)(ze).Msg(s)
}

func (ze *zeroEvent) Msgf(format string, v ...interface{}) {
	(*zerolog.Event)(ze).Msgf(format, v...)
}

//-------------------------------------------------------------------------------------------------

// AnErr adds the field key with serialized err to the ZeroEvent context.
// If err is nil, no field is added.
func (ze *zeroEvent) AnErr(key string, err error) ZeroEvent {
	ev := (*zerolog.Event)(ze).AnErr(key, err)
	return (*zeroEvent)(ev)
}

// Bool adds the field key with val as a bool to the ZeroEvent context.
func (ze *zeroEvent) Bool(key string, val bool) ZeroEvent {
	ev := (*zerolog.Event)(ze).Bool(key, val)
	return (*zeroEvent)(ev)
}

// Bytes adds the field key with val as a string to the ZeroEvent context.
//
// Runes outside of normal ASCII ranges will be hex-encoded in the resulting
// JSON.
func (ze *zeroEvent) Bytes(key string, val []byte) ZeroEvent {
	ev := (*zerolog.Event)(ze).Bytes(key, val)
	return (*zeroEvent)(ev)
}

// Dur adds the field key with duration d stored as zerolog.DurationFieldUnit.
// If zerolog.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func (ze *zeroEvent) Dur(key string, val time.Duration) ZeroEvent {
	ev := (*zerolog.Event)(ze).Dur(key, val)
	return (*zeroEvent)(ev)
}

// Err adds the field "error" with serialized err to the ZeroEvent context.
// If err is nil, no field is added.
//
// To customize the key name, change zerolog.ErrorFieldName.
//
// If Stack() has been called before and zerolog.ErrorStackMarshaler is defined,
// the err is passed to ErrorStackMarshaler and the result is appended to the
// zerolog.ErrorStackFieldName.
func (ze *zeroEvent) Err(err error) ZeroEvent {
	ev := (*zerolog.Event)(ze).Err(err)
	return (*zeroEvent)(ev)
}

// Hex adds the field key with val as a hex string to the ZeroEvent context.
func (ze *zeroEvent) Hex(key string, val []byte) ZeroEvent {
	ev := (*zerolog.Event)(ze).Hex(key, val)
	return (*zeroEvent)(ev)
}

// Int adds the field key with i as a int to the ZeroEvent context.
func (ze *zeroEvent) Int(key string, val int) ZeroEvent {
	ev := (*zerolog.Event)(ze).Int(key, val)
	return (*zeroEvent)(ev)
}

// Ints adds the field key with i as a int to the ZeroEvent context.
func (ze *zeroEvent) Ints(key string, val []int) ZeroEvent {
	ev := (*zerolog.Event)(ze).Ints(key, val)
	return (*zeroEvent)(ev)
}

// Int64 adds the field key with i as a int64 to the ZeroEvent context.
func (ze *zeroEvent) Int64(key string, val int64) ZeroEvent {
	ev := (*zerolog.Event)(ze).Int64(key, val)
	return (*zeroEvent)(ev)
}

// Interface adds the field key with i marshaled using reflection.
func (ze *zeroEvent) Interface(key string, val interface{}) ZeroEvent {
	ev := (*zerolog.Event)(ze).Interface(key, val)
	return (*zeroEvent)(ev)
}

// Str adds the field key with val as a string to the ZeroEvent context.
func (ze *zeroEvent) Str(key string, val string) ZeroEvent {
	ev := (*zerolog.Event)(ze).Str(key, val)
	return (*zeroEvent)(ev)
}

// Strs adds the field key with vals as a []string to the ZeroEvent context.
func (ze *zeroEvent) Strs(key string, val []string) ZeroEvent {
	ev := (*zerolog.Event)(ze).Strs(key, val)
	return (*zeroEvent)(ev)
}

// Stringer adds the field key with val.String() (or null if val is nil) to the ZeroEvent context.
func (ze *zeroEvent) Stringer(key string, val fmt.Stringer) ZeroEvent {
	ev := (*zerolog.Event)(ze).Stringer(key, val)
	return (*zeroEvent)(ev)
}

// Time adds the field key with t formated as string using zerolog.TimeFieldFormat.
func (ze *zeroEvent) Time(key string, val time.Time) ZeroEvent {
	ev := (*zerolog.Event)(ze).Time(key, val)
	return (*zeroEvent)(ev)
}

// Uint adds the field key with i as a uint to the ZeroEvent context.
func (ze *zeroEvent) Uint(key string, val uint) ZeroEvent {
	ev := (*zerolog.Event)(ze).Uint(key, val)
	return (*zeroEvent)(ev)
}

// Uint adds the field key with i as a uint to the ZeroEvent context.
func (ze *zeroEvent) Uints(key string, val []uint) ZeroEvent {
	ev := (*zerolog.Event)(ze).Uints(key, val)
	return (*zeroEvent)(ev)
}

// Uint64 adds the field key with i as a uint to the ZeroEvent context.
func (ze *zeroEvent) Uint64(key string, val uint64) ZeroEvent {
	ev := (*zerolog.Event)(ze).Uint64(key, val)
	return (*zeroEvent)(ev)
}
