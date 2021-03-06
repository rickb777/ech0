package testlogger

//go:generate runtemplate -v -tpl threadsafe/list.tpl Type=*TestLogEvent

import (
	"fmt"
	"github.com/rickb777/ech0/v3"
	"strings"
	"time"
)

// TestLogEvent describes one item in a linked list that holds a single log message.
type TestLogEvent struct {
	realEvent   ech0.ZeroEvent
	Method, Key string
	Val         interface{}
	Next        *TestLogEvent
	done        func(msg string)
}

var _ ech0.ZeroEvent = &TestLogEvent{}

func (ev *TestLogEvent) String() string {
	return ev.buildString(&strings.Builder{})
}

func (ev *TestLogEvent) buildString(buf *strings.Builder) string {
	if ev == nil {
		return ""
	}

	if ev.Key == "" {
		fmt.Fprintf(buf, "%s(%v)", ev.Method, ev.Val)
	} else {
		fmt.Fprintf(buf, "%s(%s, %v)", ev.Method, ev.Key, ev.Val)
	}

	if ev.Next == nil {
		return buf.String()
	}

	fmt.Fprintf(buf, ".")
	return ev.Next.buildString(buf)
}

func (ev *TestLogEvent) Send() {
	if ev.realEvent != nil {
		ev.realEvent.Send()
	}
	ev.Next = &TestLogEvent{Method: "Send"}
	if ev.done != nil {
		ev.done("")
	}
}

func (ev *TestLogEvent) Msg(s string) {
	if ev.realEvent != nil {
		ev.realEvent.Msg(s)
	}
	ev.Next = &TestLogEvent{Method: "Msg", Val: s}
	if ev.done != nil {
		ev.done(s)
	}
}

func (ev *TestLogEvent) Msgf(format string, v ...interface{}) {
	ev.Msg(fmt.Sprintf(format, v...))
}

func (ev *TestLogEvent) add(re ech0.ZeroEvent, method, key string, val interface{}) ech0.ZeroEvent {
	next := &TestLogEvent{realEvent: re, Method: method, Key: key, Val: val, done: ev.done}
	ev.Next = next
	return next
}

func (ev *TestLogEvent) AnErr(key string, val error) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.AnErr(key, val)
	}
	return ev.add(re, "AnErr", key, val)
}

func (ev *TestLogEvent) Bool(key string, val bool) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Bool(key, val)
	}
	return ev.add(re, "Bool", key, val)
}

func (ev *TestLogEvent) Bools(key string, val []bool) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Bools(key, val)
	}
	return ev.add(re, "Bools", key, val)
}

func (ev *TestLogEvent) Bytes(key string, val []byte) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Bytes(key, val)
	}
	return ev.add(re, "Bytes", key, val)
}

func (ev *TestLogEvent) Dur(key string, val time.Duration) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Dur(key, val)
	}
	return ev.add(re, "Dur", key, val)
}

func (ev *TestLogEvent) Dict(key string, dict ech0.ZeroEvent) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Dict(key, dict)
	}
	return ev.add(re, "Dict", key, dict)
}

func (ev *TestLogEvent) Err(err error) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Err(err)
	}
	return ev.add(re, "Err", "error", err)
}

func (ev *TestLogEvent) Hex(key string, val []byte) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Hex(key, val)
	}
	return ev.add(re, "Hex", key, val)
}

func (ev *TestLogEvent) Int(key string, val int) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Int(key, val)
	}
	return ev.add(re, "Int", key, val)
}

func (ev *TestLogEvent) Ints(key string, val []int) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Ints(key, val)
	}
	return ev.add(re, "Ints", key, val)
}

func (ev *TestLogEvent) Int64(key string, val int64) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Int64(key, val)
	}
	return ev.add(re, "Int64", key, val)
}

func (ev *TestLogEvent) Interface(key string, val interface{}) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Interface(key, val)
	}
	return ev.add(re, "Interface", key, val)
}

func (ev *TestLogEvent) Str(key, val string) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Str(key, val)
	}
	return ev.add(re, "Str", key, val)
}

func (ev *TestLogEvent) Strs(key string, val []string) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Strs(key, val)
	}
	return ev.add(re, "Strs", key, val)
}

func (ev *TestLogEvent) Stringer(key string, val fmt.Stringer) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Stringer(key, val)
	}
	return ev.add(re, "Stringer", key, val)
}

func (ev *TestLogEvent) Time(key string, val time.Time) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Time(key, val)
	}
	return ev.add(re, "Time", key, val)
}

func (ev *TestLogEvent) Timestamp() ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Timestamp()
	}
	return ev.add(re, "Timestamp", "", nil)
}

func (ev *TestLogEvent) Uint(key string, val uint) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Uint(key, val)
	}
	return ev.add(re, "Uint", key, val)
}

func (ev *TestLogEvent) Uints(key string, val []uint) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Uints(key, val)
	}
	return ev.add(re, "Uints", key, val)
}

func (ev *TestLogEvent) Uint64(key string, val uint64) ech0.ZeroEvent {
	var re ech0.ZeroEvent
	if ev.realEvent != nil {
		re = ev.realEvent.Uint64(key, val)
	}
	return ev.add(re, "Uint64", key, val)
}

// FindByKey searches through the linked list of TestLogEvents to find the (first)
// one with a given key, or the end of the list (nil). Use with Value.
func (ev *TestLogEvent) FindByKey(key string) *TestLogEvent {
	if ev == nil {
		return nil
	}
	if ev.Key == key {
		return ev
	}
	return ev.Next.FindByKey(key)
}

// Value returns the value of one list item. The item may be nil, in which
// case Value returns nil.
func (ev *TestLogEvent) Value() interface{} {
	if ev == nil {
		return nil
	}
	return ev.Val
}

//-------------------------------------------------------------------------------------------------

// First returns the first event.
// This returns nil if the list is empty.
func (list *TestLogEventList) First() *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil
	}

	// return Next because the lists always start with the blank level setting event
	return list.m[0].Next
}

// Last returns the last event.
// This returns nil if the list is empty.
func (list *TestLogEventList) Last() *TestLogEvent {
	list.s.RLock()
	defer list.s.RUnlock()

	if len(list.m) == 0 {
		return nil
	}

	// return Next because the lists always start with the blank level setting event
	return list.m[len(list.m)-1].Next
}
