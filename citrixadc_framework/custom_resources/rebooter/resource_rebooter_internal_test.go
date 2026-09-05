package rebooter

import (
	"errors"
	"testing"
)

// TestRebooterIsConnectionTeardownError verifies the GitHub #980 guard: the two
// transport errors the appliance produces when it reboots mid-request must be
// recognized (and therefore tolerated), while genuine errors and "never reachable"
// errors must NOT be swallowed.
func TestRebooterIsConnectionTeardownError(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		// The two exact transport errors reported in GitHub issue #980.
		{"eof", errors.New(`Post "https://10.0.0.1/nitro/v1/config/reboot": EOF`), true},
		{"connreset", errors.New(`Post "https://10.0.0.1/nitro/v1/config/reboot": read tcp 10.0.0.2:60542->10.0.0.1:443: read: connection reset by peer`), true},
		// Genuine errors that must still fail the apply (not masked by the guard).
		{"nitro_errorcode", errors.New(`[ERROR] nitro-go: failed: {"errorcode": 354, "message": "Not authorized"}`), false},
		{"conn_refused", errors.New(`Post "https://10.0.0.1/nitro/v1/config/reboot": dial tcp 10.0.0.1:443: connect: connection refused`), false},
		{"timeout", errors.New(`Post "https://10.0.0.1/nitro/v1/config/reboot": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := rebooterIsConnectionTeardownError(tc.err); got != tc.want {
				t.Fatalf("rebooterIsConnectionTeardownError(%q) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
