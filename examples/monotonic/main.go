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
