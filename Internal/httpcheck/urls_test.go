package input

import "testing"

func TestNormalizeURL_AddsHTTPSWhenMissingScheme(t *testing.T) {
    got := NormalizeURL("example.com")
    want := "https://example.com"
    if got != want {
        t.Fatalf("got %q, want %q", got, want)
    }
}

func TestNormalizeURL_LeavesExistingScheme(t *testing.T) {
    got := NormalizeURL("http://example.com")
    want := "http://example.com"
    if got != want {
        t.Fatalf("got %q, want %q", got, want)
    }
}
