package testlogger

import (
	. "github.com/onsi/gomega"
	"github.com/rickb777/ech0"
	"github.com/rs/zerolog"
	"testing"
)

func Test1(t *testing.T) {
	g := NewGomegaWithT(t)
	zl := zerolog.New(zerolog.NewConsoleWriter())
	zl.Info().Msg("foooo")
	z := ech0.Wrap(zl)
	z.Info().Msg("bar")
	tl := New(z)

	tl.Int("a", 1).Level(zerolog.InfoLevel).Warn().Int("b", 2).Msg("baz")

	g.Expect(tl.Warns).To(HaveLen(1))
	g.Expect(tl.LastWarn().FindByKey("b").Value()).To(Equal(2))
	g.Expect(tl.LastWarn().FindByKey("").Value()).To(Equal("baz"))

	tl.Reset()

	g.Expect(tl.Warns).To(BeEmpty())
	g.Expect(tl.LastWarn()).To(BeNil())
}
