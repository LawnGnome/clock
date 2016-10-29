package clock

// This is every constant defined by /usr/include/linux/time.h as of Linux 4.8.
// Some of these constants are hardware specific.
//
// Refer to the clock_gettime(2) manpage for further details.
const (
	CLOCK_REALTIME           = 0
	CLOCK_MONOTONIC          = 1
	CLOCK_PROCESS_CPUTIME_ID = 2
	CLOCK_THREAD_CPUTIME_ID  = 3
	CLOCK_MONOTONIC_RAW      = 4
	CLOCK_REALTIME_COARSE    = 5
	CLOCK_MONOTONIC_COARSE   = 6
	CLOCK_BOOTTIME           = 7
	CLOCK_REALTIME_ALARM     = 8
	CLOCK_BOOTTIME_ALARM     = 9
	CLOCK_SGI_CYCLE          = 10
	CLOCK_TAI                = 11
)
