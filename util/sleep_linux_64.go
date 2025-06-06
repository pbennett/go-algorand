// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

//go:build linux && !(arm || 386)
// +build linux,!arm,!386

package util

import (
	"syscall"
	"time"
)

// NanoSleep sleeps for the given d duration.
func NanoSleep(d time.Duration) {
	timeSpec := &syscall.Timespec{
		Nsec: d.Nanoseconds() % time.Second.Nanoseconds(),
		Sec:  d.Nanoseconds() / time.Second.Nanoseconds(),
	}
	syscall.Nanosleep(timeSpec, nil) //nolint:errcheck // ignoring error
}
