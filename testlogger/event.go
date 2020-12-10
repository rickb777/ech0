package testlogger

import (
	"fmt"
	"github.com/rickb777/ech0"
	"time"
)

type TestLogEvent struct {
	Method, Key string
	Val         interface{}
	Next        *TestLogEvent
	done        func(msg string)
}

var _ ech0.ZeroEvent = &TestLogEvent{}

func (ev *TestLogEvent) Send() {
	ev.Next = &TestLogEvent{Method: "Send"}
	if ev.done != nil {
		ev.done("")
	}
}

func (ev *TestLogEvent) Msg(s string) {
	ev.Next = &TestLogEvent{Method: "Msg", Val: s}
	if ev.done != nil {
		ev.done(s)
	}
}

func (ev *TestLogEvent) Msgf(format string, v ...interface{}) {
	ev.Msg(fmt.Sprintf(format, v...))
}

func (ev *TestLogEvent) add(method, key string, val interface{}) ech0.ZeroEvent {
	next := &TestLogEvent{Method: method, Key: key, Val: val, done: ev.done}
	ev.Next = next
	return next
}

func (ev *TestLogEvent) AnErr(key string, val error) ech0.ZeroEvent {
	return ev.add("AnErr", key, val)
}

func (ev *TestLogEvent) Bool(key string, val bool) ech0.ZeroEvent {
	return ev.add("Bool", key, val)
}

func (ev *TestLogEvent) Bytes(key string, val []byte) ech0.ZeroEvent {
	return ev.add("Bytes", key, val)
}

func (ev *TestLogEvent) Dur(key string, val time.Duration) ech0.ZeroEvent {
	return ev.add("Dur", key, val)
}

func (ev *TestLogEvent) Err(err error) ech0.ZeroEvent {
	return ev.add("Err", "error", err)
}

func (ev *TestLogEvent) Hex(key string, val []byte) ech0.ZeroEvent {
	return ev.add("Hex", key, val)
}

func (ev *TestLogEvent) Int(key string, val int) ech0.ZeroEvent {
	return ev.add("Int", key, val)
}

func (ev *TestLogEvent) Ints(key string, val []int) ech0.ZeroEvent {
	return ev.add("Ints", key, val)
}

func (ev *TestLogEvent) Int64(key string, val int64) ech0.ZeroEvent {
	return ev.add("Int64", key, val)
}

func (ev *TestLogEvent) Interface(key string, val interface{}) ech0.ZeroEvent {
	return ev.add("Interface", key, val)
}

func (ev *TestLogEvent) Str(key, val string) ech0.ZeroEvent {
	return ev.add("Str", key, val)
}

func (ev *TestLogEvent) Strs(key string, val []string) ech0.ZeroEvent {
	return ev.add("Strs", key, val)
}

func (ev *TestLogEvent) Stringer(key string, val fmt.Stringer) ech0.ZeroEvent {
	return ev.add("Stringer", key, val)
}

func (ev *TestLogEvent) Time(key string, val time.Time) ech0.ZeroEvent {
	return ev.add("Time", key, val)
}

func (ev *TestLogEvent) Uint(key string, val uint) ech0.ZeroEvent {
	return ev.add("Uint", key, val)
}

func (ev *TestLogEvent) Uints(key string, val []uint) ech0.ZeroEvent {
	return ev.add("Uints", key, val)
}

func (ev *TestLogEvent) Uint64(key string, val uint64) ech0.ZeroEvent {
	return ev.add("Uint64", key, val)
}

func (ev *TestLogEvent) FindByKey(key string) *TestLogEvent {
	if ev == nil {
		return nil
	}
	if ev.Key == key {
		return ev
	}
	return ev.FindByKey(key)
}

func (ev *TestLogEvent) Value() interface{} {
	if ev == nil {
		return nil
	}
	return ev.Val
}
