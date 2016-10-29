# clock

This Go package provides an extremely thin wrapper over the POSIX
[clock API](http://pubs.opengroup.org/onlinepubs/009695399/functions/clock_getres.html).

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
