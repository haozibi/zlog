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
