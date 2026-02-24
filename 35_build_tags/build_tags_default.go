// This file provides the default implementation that is always compiled.
// In a real project, you might have platform-specific files like:
//
//	//go:build linux
//	// build_tags_linux.go - Linux-specific implementation
//
//	//go:build darwin
//	// build_tags_darwin.go - macOS-specific implementation
//
// Each platform file would provide the same function signatures with
// different implementations. The default file (this one) has no build
// constraint, so it is always included.
//
// If you use build tags incorrectly, you may end up with missing
// function definitions on certain platforms, causing compile errors only
// on those targets.
package buildtags

// DefaultMessage is a constant that is always available because this file
// has no build constraint. Platform-specific files could override this
// behavior by providing a function that returns a different value.
const DefaultMessage = "default implementation"

// GetDefaultMessage returns the default message. In a real application,
// platform-specific files would provide alternative implementations
// guarded by build constraints.
func GetDefaultMessage() string {
	return DefaultMessage
}

// BuildTagExample documents common patterns for build constraints.
//
// Single OS constraint:
//
//	//go:build linux
//
// Multiple OS (OR):
//
//	//go:build linux || darwin
//
// OS and architecture (AND):
//
//	//go:build linux && amd64
//
// Negation:
//
//	//go:build !windows
//
// Custom tag (requires -tags flag):
//
//	//go:build mytag
//
// Complex expression:
//
//	//go:build (linux || darwin) && amd64
func BuildTagExample() string {
	return "see source comments for build tag examples"
}
