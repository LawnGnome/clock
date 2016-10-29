# clock

This Go package provides an extremely thin wrapper over the POSIX
[clock API](http://pubs.opengroup.org/onlinepubs/009695399/functions/clock_getres.html).

If you would prefer a slightly thicker wrapper that integrates with Go's
`time.Time` type, you may want to check out Dave Cheney's
https://godoc.org/github.com/davecheney/junk/clock instead. This package is
very much intended for Go code that has to integrate with non-Go code using
non-realtime clocks.

At present, only Linux is supported. Adding additional platforms should be as
simple as defining the available `CLOCK_*` constants with their values:
contributions are welcome, and [`clock_linux.go`](clock_linux.go) should
provide a good template to start with.

## Usage

A basic example of getting the current value of the `CLOCK_MONOTONIC` clock:

```go
package main

import (
	"fmt"
	"os"

	"github.com/LawnGnome/clock"
)

func main() {
	ts, err := clock.Gettime(clock.CLOCK_MONOTONIC)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("The current monotonic time is: %d.%d\n", ts.Sec, ts.Nsec)
}
```
