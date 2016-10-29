// This package provides thin wrappers for the clock functions specified by
// POSIX-1.2001.
package clock

// #include <time.h>
import "C"

import (
	"fmt"
)

// A Go flavoured version of the timespec struct defined by POSIX.1-2001.
type Timespec struct {
	// Seconds since the epoch. The epoch is clock specific.
	Sec int64
	// Nanoseconds: this value will be between 0 and 999999999, inclusive.
	Nsec int64
}

func newTimespec(c_ts *C.struct_timespec) *Timespec {
	if c_ts == nil {
		return nil
	}

	return &Timespec{
		Sec:  (int64)(c_ts.tv_sec),
		Nsec: (int64)(c_ts.tv_nsec),
	}
}

// A wrapper for the clock_getres() function. If clock_getres() fails, the
// returned Timespec pointer will be nil.
func Getres(clockID int) (*Timespec, error) {
	var c_ts C.struct_timespec

	rv, err := C.clock_getres((C.clockid_t)(clockID), &c_ts)
	if err != nil {
		return nil, err
	} else if rv != 0 {
		return nil, fmt.Errorf("unknown error; return value %d", rv)
	}

	return newTimespec(&c_ts), nil
}

// A wrapper for the clock_gettime() function. If clock_gettime() fails, the
// returned Timespec pointer will be nil.
func Gettime(clockID int) (*Timespec, error) {
	var c_ts C.struct_timespec

	rv, err := C.clock_gettime((C.clockid_t)(clockID), &c_ts)
	if err != nil {
		return nil, err
	} else if rv != 0 {
		return nil, fmt.Errorf("unknown error; return value %d", rv)
	}

	return newTimespec(&c_ts), nil
}

// A wrapper for the clock_settime() function.
func Settime(clockID int, ts *Timespec) error {
	c_ts := &C.struct_timespec{(C.__time_t)(ts.Sec), (C.__syscall_slong_t)(ts.Nsec)}

	rv, err := C.clock_settime((C.clockid_t)(clockID), c_ts)
	if err != nil {
		return err
	} else if rv != 0 {
		return fmt.Errorf("unknown error; return value %d", rv)
	}

	return nil
}
