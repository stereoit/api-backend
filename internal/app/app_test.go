package app

import "testing"

func TestGetEnv(t *testing.T) {
	want := "default"
	if got := GetEnv("NONEXISTEN", "default"); got != want {
		t.Errorf("getEnv() = %q, want %q", got, want)
	}
}
