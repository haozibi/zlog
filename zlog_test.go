package zlog

import (
	"os"
	"testing"
)

func TestNewBasicLog(t *testing.T) {

	w := os.Stdout

	NewJSONLog(w)
	NewBasicLog(w)

	ZDebug().Str("Test", "ok").Msg("[Test]")
	ZInfo().Str("Test", "ok").Msg("[Test]")
	ZWarn().Str("Test", "ok").Msg("[Test]")
	ZError().Str("Test", "ok").Msg("[Test]")
	// ZFatal().Str("Test", "ok").Msg("[Test]")

	Debugf("%s", "ok")
}

func TestBasicWithConfig(t *testing.T) {

	w := os.Stdout

	NewBasicLog(w)
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasicLog(w, WithNoColor(true))
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasicLog(w, WithTimeFormat("2006-01-02"), WithNoColor(true))
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasicLog(w, WithTimeFormat("2006-01-02"), WithNoColor(true), WithDebug(true))
	ZDebug().Str("Test", "ok").Msg("[With]")

	NewBasicLog(w, WithTimeFormat("2006-01-02"), WithNoColor(true), WithDebug(true), WithDeep(1))
	ZDebug().Str("Test", "ok").Msg("[With]")

}
