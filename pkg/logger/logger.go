package logger

import (
	"fmt"
	"io"
	"runtime"
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
	e := z.addKeyValues(z.zlog.Debug(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) Error(msg string, keyAndValues ...any) {
	e := z.addKeyValues(z.zlog.Error(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) Fatal(msg string, keyAndValues ...any) {
	e := z.addKeyValues(z.zlog.Fatal(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) Info(msg string, keyAndValues ...any) {
	e := z.addKeyValues(z.zlog.Info(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) Trace(msg string, keyAndValues ...any) {
	e := z.addKeyValues(z.zlog.Trace(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) Warning(msg string, keyAndValues ...any) {
	e := z.addKeyValues(z.zlog.Warn(), keyAndValues...)
	e.Msg(msg)
}

func (z *Zerolog) addKeyValues(e *zerolog.Event, keyAndValues ...any) *zerolog.Event {
	if z.isCallerEnabled {
		_, fp, line, ok := runtime.Caller(2)
		if ok {
			e = e.Str("caller", fmt.Sprintf("%s:%d", fp, line))
		}
	}
	if len(z.layers) > 0 {
		e = e.Any("layer", strings.Join(z.layers, "/"))
	}
	for i := range keyAndValues {
		if i%2 == 1 {
			continue
		}
		if i+1 >= len(keyAndValues) {
			continue
		}
		e = e.Any(fmt.Sprintf("%v", keyAndValues[i]), keyAndValues[i+1])
	}

	return e
}
