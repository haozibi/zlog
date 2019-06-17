package zlog

import "os"

func Example() {
	NewBasicLog(os.Stdout, WithNoColor(true), WithDebug(true))
	// zlog.NewJSONLog(os.Stdout)

	ZInfo().
		Int("z", 100-1).
		Msg("just do it")

	ZDebug().
		Float64("f", 3.1415926).
		Msgf("hello %s", "zlog")
	//Output:
	//2019-06-17 16:10:18.998152 INF example_test.go:11 > just do it z=99
	//2019-06-17 16:10:18.998741 DBG example_test.go:15 > hello zlog f=3.1415926
}
