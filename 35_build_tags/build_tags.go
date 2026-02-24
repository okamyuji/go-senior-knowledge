// Package buildtags demonstrates Go's build constraint system for
// conditional compilation. Build constraints (also called build tags)
// control which files are included in a package during compilation.
//
// Syntax (Go 1.17+):
//
//	//go:build linux && amd64
//
// Legacy syntax (still supported but deprecated):
//
//	// +build linux,amd64
//
// The //go:build line uses boolean expressions with &&, ||, and !
// operators. It must appear before the package clause, with a blank line
// separating it from the package statement.
//
// Common build tags:
//   - OS tags: linux, darwin, windows, freebsd, etc.
//   - Architecture tags: amd64, arm64, 386, etc.
//   - Custom tags: enabled with -tags flag (go build -tags mytag)
//   - ignore: prevents a file from being compiled at all
//
// Warning: build tags can silently change program behavior. If a file
// with a build tag provides a different implementation of a function,
// the program compiles and runs without error on all platforms, but the
// behavior varies depending on which files are included. This can cause
// subtle bugs that only manifest on specific platforms or with specific
// build flags.
//
// Another trap is forgetting the blank line between //go:build and the
// package clause. Without the blank line, the constraint is treated as a
// regular comment and ignored.
package buildtags

import "runtime"

// GetPlatformInfo returns a string describing the current platform.
// This function is always available regardless of build tags, because
// it uses runtime.GOOS and runtime.GOARCH which are set at compile time.
func GetPlatformInfo() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}
