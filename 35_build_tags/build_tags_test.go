package buildtags

import (
	"runtime"
	"strings"
	"testing"
)

func TestGetPlatformInfoNonEmpty(t *testing.T) {
	info := GetPlatformInfo()
	if info == "" {
		t.Fatal("GetPlatformInfo returned empty string")
	}
}

func TestGetPlatformInfoFormat(t *testing.T) {
	info := GetPlatformInfo()

	// The format should be "os/arch".
	parts := strings.SplitN(info, "/", 2)
	if len(parts) != 2 {
		t.Fatalf("expected format os/arch, got %q", info)
	}
	if parts[0] == "" || parts[1] == "" {
		t.Fatalf("os or arch is empty in %q", info)
	}
}

func TestGetPlatformInfoMatchesRuntime(t *testing.T) {
	info := GetPlatformInfo()
	expected := runtime.GOOS + "/" + runtime.GOARCH

	if info != expected {
		t.Fatalf("expected %q, got %q", expected, info)
	}
}

func TestDefaultMessageConstant(t *testing.T) {
	if DefaultMessage == "" {
		t.Fatal("DefaultMessage should not be empty")
	}
}

func TestGetDefaultMessage(t *testing.T) {
	msg := GetDefaultMessage()
	if msg != DefaultMessage {
		t.Fatalf("expected %q, got %q", DefaultMessage, msg)
	}
}

func TestBuildTagExample(t *testing.T) {
	result := BuildTagExample()
	if result == "" {
		t.Fatal("BuildTagExample should return a non-empty string")
	}
}
