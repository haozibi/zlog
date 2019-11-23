package main

import (
	"fmt"
	"os"

	"github.com/haozibi/zlog"
	"github.com/pkg/errors"
)

func init() {

	zlog.NewBasic(os.Stdout, zlog.WithColor(), zlog.WithDebug())
	// zlog.NewJSONLog(os.Stdout)
}

func main() {
	zlog.ZInfo().
		Int("z", 100-1).
		Msg("just do it")

	zlog.ZDebug().
		Float64("f", 3.1415926).
		Msgf("hello %s", "zlog")

	var err error

	err = doit()
	if err != nil {
		zlog.ZError().Stack().Err(err).Msg("[doit] some error")
	}

	err = doErr()
	if err != nil {
		zlog.ZError().Err(err).Msg("[do error]")
		zlog.ZError().Stack().Err(err).Msg("[do error]")
	}

}

func doit() error {
	err := doErr()
	if err != nil {
		return errors.Wrap(err, "just error")
	}
	return nil
}

func doErr() error {
	return fmt.Errorf("some error")
}
