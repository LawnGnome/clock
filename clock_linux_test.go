package clock

import (
	"testing"
)

const CLOCK_INVALID = 16

func TestGetresLinux(t *testing.T) {
	// Test with an invalid clock ID.
	ts, err := Getres(CLOCK_INVALID)
	if err == nil || ts != nil {
		t.Errorf("Invalid clock was successful:\nerr: %v\nts: %v\n", err, ts)
	}
	if err.Error() != "invalid argument" {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test with a valid, Linux only clock ID.
	ts, err = Getres(CLOCK_MONOTONIC)
	if err != nil {
		t.Error(err)
	}
	assertTimespecValid(t, ts)
}

func TestGettimeLinux(t *testing.T) {
	// Test with an invalid clock ID.
	ts, err := Gettime(CLOCK_INVALID)
	if err == nil || ts != nil {
		t.Errorf("Invalid clock was successful:\nerr: %v\nts: %v\n", err, ts)
	}
	if err.Error() != "invalid argument" {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test with a valid, Linux only clock ID.
	ts, err = Gettime(CLOCK_MONOTONIC)
	if err != nil {
		t.Error(err)
	}
	assertTimespecValid(t, ts)
}
