// Copyright 2021 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

// +build invariants,!race tracing,!race

package invariants

import "runtime"

// SetFinalizer is a wrapper around runtime.SetFinalizer that is a no-op under
// race builds or if neither the invariants or tracing build tags are
// specified.
func SetFinalizer(obj, finalizer interface{}) {
	runtime.SetFinalizer(obj, finalizer)
}
