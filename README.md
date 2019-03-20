[![Build Status](https://travis-ci.org/i3/go-i3.svg?branch=master)](https://travis-ci.org/i3/go-i3)
[![Go Report Card](https://goreportcard.com/badge/go.i3wm.org/i3)](https://goreportcard.com/report/go.i3wm.org/i3)
[![GoDoc](https://godoc.org/go.i3wm.org/i3?status.svg)](https://godoc.org/go.i3wm.org/sway)

Package sway provides a convenient interface to the i3 window manager via [its IPC
interface](https://i3wm.org/docs/ipc.html).

See [its documentation](https://godoc.org/go.i3wm.org/i3) for more details.

## Advantages over other sway IPC packages

Here comes a grab bag of features to which we paid attention. At the time of
writing, most other sway IPC packages lack at least a good number of these
features:

* Retries are transparently handled: programs using this package will recover
  automatically from in-place sway restarts. Additionally, programs can be started
  from xsession or user sessions before sway is even running.

* Version checks are transparently handled: if your program uses features which
  are not supported by the running sway version, helpful error messages will be
  returned at run time.

* Comprehensive: the entire documented IPC interface of the latest stable sway
  version is covered by this package. Tagged releases match sway’s major and minor
  version.

* Consistent and familiar: once familiar with the sway IPC protocol’s features,
  you should have no trouble matching the documentation to API and vice-versa.

* Good test coverage (hard to display in a badge, as our multi-process setup
  breaks `go test`’s `-coverprofile` flag).

* Implemented in pure Go, without resorting to the `unsafe` package.

* Works on little and big endian architectures.

## Scope

sway’s entire documented IPC interface is available in this package.

In addition, helper functions which are useful for a broad range of programs
(and only those!) are provided, e.g. Node’s FindChild and FindFocused.

Packages which introduce higher-level abstractions should feel free to use this
package as a building block.

## Assumptions

* The `sway(1)` binary must be in `$PATH` so that the IPC socket path can be retrieved.
* For transparent version checks to work, the running sway version must be ≥ 4.3 (released 2012-09-19).

## Testing

Be sure to include the target sway version (the most recent stable release) in
`$PATH` and use `go test` as usual:

```shell
PATH=~/i3/build/i3:$PATH go test -v go.i3wm.org/i3
```
