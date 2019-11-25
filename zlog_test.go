package zlog

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestNewBasicLog(t *testing.T) {

	w := &bytes.Buffer{}

	NewJSON(w)
	NewBasic(w)

	ZDebug().Str("Test", "ok").Msg("[Test]")
	ZInfo().Str("Test", "ok").Msg("[Test]")
	ZWarn().Str("Test", "ok").Msg("[Test]")
	ZError().Str("Test", "ok").Msg("[Test]")

	got := w.String()
	fmt.Println(got)
}

func TestBasicWithConfig(t *testing.T) {

	w := os.Stdout

	NewBasic(w)
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasic(w, WithColor())
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasic(w, WithTimeFormat("2006-01-02"), WithColor())
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasic(w, WithTimeFormat("2006-01-02"), WithColor(), WithDebug())
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasic(w, WithTimeFormat("2006-01-02"), WithColor(), WithDebug(), WithDeep(1))
	ZDebug().Str("Test", "ok").Msg("[With]")
}

func TestLogStack(t *testing.T) {

	out := &bytes.Buffer{}
	NewBasic(out)

	err := errors.Wrap(errors.New("error message"), "from error")
	ZInfo().Stack().Err(err).Msg("[err]")

	got := out.String()
	fmt.Println(got)

	out2 := &bytes.Buffer{}
	NewBasic(out2)
	err2 := fmt.Errorf("just error")
	ZInfo().Stack().Err(err2).Msg("[err2]")
	fmt.Println(out2.String())
}

func TestLogLevel(t *testing.T) {

	out := &bytes.Buffer{}
	NewBasic(out, WithLevel(InfoLevel))

	ZDebug().Msg("debug")
	ZInfo().Msg("info")

	got := out.String()
	fmt.Println(got)
}
