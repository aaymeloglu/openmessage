// Package racetest lets tests opt out of `go test -race` runs when they
// cannot find a race.
//
// The race detector instruments every Go package in the binary, including
// modernc.org/sqlite, which is SQLite transpiled to Go. Database-heavy tests
// therefore run 10-25x slower under -race. A test that never runs code
// concurrently pays that cost for nothing, so it calls SkipSequential and
// still runs in full in the normal (non-race) test job.
package racetest

import "testing"

// SkipSequential skips t when the binary is built with -race. Use it only in
// tests whose code under test starts no goroutines; anything with concurrent
// code must keep running under the race detector.
func SkipSequential(t testing.TB) {
	t.Helper()
	if Enabled {
		t.Skip("sequential test: nothing for the race detector to find; covered by the non-race job")
	}
}
