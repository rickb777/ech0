package testlogger

import (
	. "github.com/onsi/gomega"
	"github.com/rickb777/ech0/v3"
	"github.com/rs/zerolog"
	"testing"
)

func Test1(t *testing.T) {
	g := NewGomegaWithT(t)
	zl := zerolog.New(zerolog.NewConsoleWriter())
	zl.Info().Msg("m1")
	z := ech0.Wrap(zl)
	z.Info().Msg("m2")
	tl := New(z)

	tl.Warn().Int("b", 2).Msg("m3")
	tl.Int("a", 100).Level(zerolog.InfoLevel).Warn().Int("c", 3).Msg("m4")
	tl.Int("a", 101).Warn().Int("d", 4).Msg("m5")

	g.Expect(tl.Infos.Len()).To(Equal(0))
	//g.Expect(tl.Infos.Drop(1).Last().FindByKey("c").Value()).To(Equal(3))
	g.Expect(tl.Warns.Len()).To(Equal(3))
	g.Expect(tl.Warns.First().FindByKey("b").Value()).To(Equal(2))
	g.Expect(tl.Warns.Drop(1).First().FindByKey("c").Value()).To(Equal(3))
	g.Expect(tl.Warns.DropLast(1).Last().FindByKey("").Value()).To(Equal("m4"))
	g.Expect(tl.LastWarn().FindByKey("").Value()).To(Equal("m5"))
	g.Expect(tl.LastWarn().FindByKey("a").Value()).To(BeNil())

	g.Expect(tl.Warns.First().String()).To(Equal("Int(b, 2).Msg(m3)"))
	g.Expect(tl.Warns.Drop(1).First().String()).To(Equal("Int(c, 3).Msg(m4)"))
	g.Expect(tl.Warns.Drop(2).First().String()).To(Equal("Int(d, 4).Msg(m5)"))

	tl.Reset()

	g.Expect(tl.Warns.IsEmpty()).To(BeTrue())
	g.Expect(tl.LastWarn()).To(BeNil())
}
