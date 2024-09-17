package logger

import (
	"io"
	"strings"

	"github.com/rs/zerolog"

	"github.com/ksckaan1/apiredator/internal/core/port"
)

var _ port.Logger = (*Zerolog)(nil)

type Zerolog struct {
	zlog            *zerolog.Logger
	layers          []string
	isCallerEnabled bool
	isTimeEnabled   bool
}

func NewZerolog() *Zerolog {
	zlog := zerolog.New(io.Discard)
	return &Zerolog{
		zlog: &zlog,
	}
}

func (z *Zerolog) WithWriter(w io.Writer) *Zerolog {
	nl := z.zlog.Output(w)
	z.zlog = &nl
	return z
}

func (z *Zerolog) WithLevel(l Level) *Zerolog {
	nl := z.zlog.Level(zerolog.Level(l))
	z.zlog = &nl
	return z
}

func (z *Zerolog) WithLayer(name string) *Zerolog {
	z.layers = append(z.layers, name)
	if len(z.layers) > 0 {
		newZlog := z.zlog.With().Str("layer", strings.Join(z.layers, "/")).Logger()
		z.zlog = &newZlog
	}
	return z
}

func (z *Zerolog) WithCaller() *Zerolog {
	z.isCallerEnabled = true
	return z
}

func (z *Zerolog) WithTime() *Zerolog {
	z.isTimeEnabled = true
	return z
}

func (z *Zerolog) Debug(msg string, keyAndValues ...any) {
	z.zlog.Debug().Fields(keyAndValues).Msg(msg)
}

func (z *Zerolog) Error(msg string, keyAndValues ...any) {
	z.zlog.Error().Fields(keyAndValues).Msg(msg)
}

func (z *Zerolog) Fatal(msg string, keyAndValues ...any) {
	z.zlog.Fatal().Fields(keyAndValues).Msg(msg)
}

func (z *Zerolog) Info(msg string, keyAndValues ...any) {
	z.zlog.Info().Fields(keyAndValues).Msg(msg)
}

func (z *Zerolog) Trace(msg string, keyAndValues ...any) {
	z.zlog.Trace().Fields(keyAndValues).Msg(msg)
}

func (z *Zerolog) Warning(msg string, keyAndValues ...any) {
	z.zlog.Warn().Fields(keyAndValues).Msg(msg)
}
