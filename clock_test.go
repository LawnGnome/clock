package clock

import (
	"math"
	"testing"
	"time"
)

func assertTimespecValid(t *testing.T, ts *Timespec) {
	if ts.Sec < 0 || ts.Nsec < 0 || ts.Nsec > 999999999 {
		t.Errorf("Unexpected ts value: %v", ts)
	}
}

func TestGetres(t *testing.T) {
	ts, err := Getres(CLOCK_REALTIME)
	if err != nil {
		t.Error(err)
	}
	assertTimespecValid(t, ts)
}

func TestGettime(t *testing.T) {
	now := time.Now()
	ts, err := Gettime(CLOCK_REALTIME)
	if err != nil {
		t.Error(err)
	}
	if math.Abs((float64)(now.Unix()-ts.Sec)) > 1 {
		t.Errorf("Unexpected difference in CLOCK_REALTIME: now=%d ts.Sec=%d", now.Unix(), ts.Sec)
	}
	assertTimespecValid(t, ts)
}

// TODO: Settime() testing.
