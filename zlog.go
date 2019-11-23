package zlog

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// - 推荐错误处理使用 "github.com/pkg/errors" 并配合 `Stack()` 方法使用
// - 错误在调用方进行处理，中间过程直接返回

var zlog zerolog.Logger

var (
	// TimeFieldFormat time format
	TimeFieldFormat = time.RFC3339

	// TimeFormatUnixNano time format
	TimeFormatUnixNano = "2006-01-02 15:04:05.999999999"

	// NoColor if set color
	NoColor = false
)

func newWriter(w io.Writer, opts ...Option) *zerolog.ConsoleWriter {

	d := defaultOptions()
	for _, o := range opts {
		o.apply(d)
	}

	// ConsoleWriter parses the JSON input and writes it in an (optionally) colorized, human-friendly format to Out.
	return &zerolog.ConsoleWriter{
		Out:        w,
		TimeFormat: d.timeFormat,
		NoColor:    d.nocolor,
	}
}

// NewBasic new basic format log
func NewBasic(w io.Writer, opts ...Option) {
	writer := newWriter(w, opts...)
	zlog = newLog(writer, opts...).Logger()
}

// NewJSON new log by json format
func NewJSON(w io.Writer, opts ...Option) {
	zlog = newLog(w, opts...).Logger()
}

func newLog(w io.Writer, opts ...Option) zerolog.Context {

	d := defaultOptions()

	for _, o := range opts {
		o.apply(d)
	}

	zerolog.TimeFieldFormat = TimeFormatUnixNano
	zerolog.ErrorStackMarshaler = d.marshalStack
	z := zerolog.New(w).With().Timestamp()

	if d.debug {
		z = z.CallerWithSkipFrameCount(2)
	}

	return z
}

// ZDebug debug log
func ZDebug() *zerolog.Event {
	return zlog.Debug()
}

// ZWarn warn log
func ZWarn() *zerolog.Event {
	return zlog.Warn()
}

// ZInfo info log
func ZInfo() *zerolog.Event {
	return zlog.Info()
}

// ZError error log
func ZError() *zerolog.Event {
	return zlog.Error()
}

// ZFatal fatal log
func ZFatal() *zerolog.Event {
	return zlog.Fatal()
}

// ZTrace trace log
// func ZTrace() *zerolog.Event {
// 	return zlog.Trace()
// }

func defaultOptions() *options {
	return &options{
		debug:        false,
		deep:         2,
		nocolor:      true,
		timeFormat:   TimeFormatUnixNano,
		marshalStack: pkgerrors.MarshalStack,
	}
}

type options struct {
	debug        bool
	deep         int
	timeFormat   string
	nocolor      bool
	marshalStack func(err error) interface{}
}

// Option overrides behavior of zlog.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// WithDebug set if debug,debug output line num
func WithDebug() Option {
	return optionFunc(func(o *options) {
		o.debug = true
	})
}

// WithDeep set line deep,default eq 2
func WithDeep(n int) Option {
	return optionFunc(func(o *options) {
		o.deep = n
	})
}

// WithColor set if has color
func WithColor() Option {
	return optionFunc(func(o *options) {
		o.nocolor = false
	})
}

// WithTimeFormat set time format when basic format
func WithTimeFormat(format string) Option {
	return optionFunc(func(o *options) {
		o.timeFormat = format
	})
}
