package ech0

import (
	"errors"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"strings"
	"testing"
	"time"
)

func TestWithLevel(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	z.WithLevel(zerolog.ErrorLevel).Int("a", 1).Send()

	g.Expect(buf.String()).To(Equal(`{"level":"error","a":1`+newline), buf.String())
}

func TestLevel(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	z.Level(zerolog.ErrorLevel).Info().Int("a", 1).Send()
	g.Expect(buf.String()).To(Equal(""), buf.String())

	z.Level(zerolog.ErrorLevel).Error().Int("a", 1).Send()
	g.Expect(buf.String()).To(Equal(`{"level":"error","a":1`+newline), buf.String())
}

func TestOutput(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(nil)
	z := Wrap(zl)

	apply(z.Output(buf).Error()).Send()

	g.Expect(buf.String()).To(Equal(`{"level":"error",`+core+newline), buf.String())
}

func TestStrIntJSON(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(nil)
	z := Wrap(zl).Output(buf)

	z.Str("a", "1").Int("b", 2).Str("c", "3").RawJSON("r", []byte(`{}`)).Info().Str("x", "y").Send()

	g.Expect(buf.String()).To(Equal(`{"level":"info","a":"1","b":2,"c":"3","r":{},"x":"y"`+newline), buf.String())
}

func TestError(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Error()).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"error",`+core+msg+newline), buf.String())
}

func TestWarn(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Warn()).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"warn",`+core+msg+newline), buf.String())
}

func TestInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Info()).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"info",`+core+msg+newline), buf.String())
}

func TestDebug(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Debug()).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"debug",`+core+msg+newline), buf.String())
}

func TestErr_not_nil(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Err(e1)).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"error","error":"x1",`+core+msg+newline), buf.String())
}

func TestErr_nil(t *testing.T) {
	g := NewGomegaWithT(t)
	buf := &strings.Builder{}
	zl := zerolog.New(buf)
	z := Wrap(zl)

	apply(z.Err(nil)).Msg("msg")

	g.Expect(buf.String()).To(Equal(`{"level":"info",`+core+msg+newline), buf.String())
}

func apply(in ZeroEvent) ZeroEvent {
	return in.Str("s1", "v1").
		Stringer("s2", time.Second).
		Bytes("b2", []byte("v2")).
		Hex("h3", []byte("~")).
		Int("i2", 2).
		Int64("i3", 3).
		Uint("u4", 4).
		Uint64("u5", 5).
		Bool("b3", true).
		Err(e2).
		Time("t1", t1).
		Dur("d4", time.Second)
}

const core = `"s1":"v1","s2":"1s","b2":"v2","h3":"7e","i2":2,"i3":3,"u4":4,"u5":5,"b3":true,"error":"x2","t1":"2020-12-01T13:14:15Z","d4":1000`
const msg = `,"message":"msg"`
const newline = "}\n"

var e1 = errors.New("x1")
var e2 = errors.New("x2")
var t1 = time.Date(2020, 12, 1, 13, 14, 15, 0, time.UTC)
