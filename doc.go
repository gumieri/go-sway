// Package sway provides a convenient interface to the sway window manager.
//
// Its function and type names don’t stutter, and all functions and methods are
// safe for concurrent use (except where otherwise noted). The package does not
// import "unsafe" and hence should be widely applicable.
//
// UNIX socket connections to sway are transparently managed by the package. Upon
// any read/write errors on a UNIX socket, the package transparently retries for
// up to 10 seconds, but only as long as the sway process keeps running.
//
// The package is published in versioned releases, where the major and minor
// version are identical to the sway release the package is compatible with
// (e.g. 4.14 implements the entire documented IPC interface of sway 4.14).
//
// This package will only ever receive additions, so versioning should only be
// relevant to you if you are interested in a recently-introduced IPC feature.
//
// Message type functions and event types are annotated with the sway version in
// which they were introduced. Under the covers, they use AtLeast, so they
// return a helpful error message at runtime if the running sway version is too
// old.
package sway
