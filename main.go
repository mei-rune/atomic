package atomic

import (
	uber_atomic "go.uber.org/atomic"
)

type Int32 = uber_atomic.Int32
type Uint32 = uber_atomic.Uint32

type Int64 = uber_atomic.Int64
type Uint64 = uber_atomic.Uint64

type Duration = uber_atomic.Duration

type Float64 = uber_atomic.Float64
type Float32 = uber_atomic.Float32

type String = uber_atomic.String
type Error = uber_atomic.Error

type Bool32 = uber_atomic.Bool
