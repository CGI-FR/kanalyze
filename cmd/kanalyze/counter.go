// Copyright (C) 2026 CGI France
//
// This file is part of kanalyze.
//
// kanalyze is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// kanalyze is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with kanalyze.  If not, see <http://www.gnu.org/licenses/>.

package main

import "hash/maphash"

type Counter struct {
	counters map[uint64]uint64
	hash     maphash.Hash
}

func NewCounter() *Counter {
	return &Counter{
		counters: make(map[uint64]uint64),
		hash:     maphash.Hash{},
	}
}

func (d *Counter) Ingest(data []byte) uint64 {
	d.hash.Reset()
	_, _ = d.hash.Write(data) // never fails
	d.counters[d.hash.Sum64()]++

	return d.counters[d.hash.Sum64()]
}
