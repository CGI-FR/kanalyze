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

import (
	"fmt"
	"strings"
	"sync"
)

const (
	maxBarLen = 100
	maxBins   = 50
	barChar   = "█"
)

type Histogram struct {
	bins      map[uint64]uint64
	max       uint64
	lastLines int
	mu        sync.Mutex
}

func NewHistogram() *Histogram {
	return &Histogram{
		bins:      make(map[uint64]uint64),
		max:       0,
		lastLines: 0,
		mu:        sync.Mutex{},
	}
}

func (h *Histogram) Update(val uint64) {
	if val > maxBins {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.bins[val]++

	if val > 1 {
		h.bins[val-1]--
	}

	if val > h.max {
		h.max = val
	}
}

func (h *Histogram) Render() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.lastLines > 0 {
		fmt.Printf("\033[%dA", h.lastLines)
	}

	fmt.Print("\033[J")

	linesPrinted := 0

	// Trouver le comptage maximum pour calculer les proportions
	var maxCount uint64
	for bin := uint64(1); bin <= h.max; bin++ {
		if h.bins[bin] > maxCount {
			maxCount = h.bins[bin]
		}
	}

	for bin := uint64(1); bin <= h.max; bin++ {
		count := h.bins[bin]

		if count > 0 {
			// Calculer la longueur proportionnelle
			var barLen int
			if maxCount > 0 {
				barLen = int(float64(count) * float64(maxBarLen) / float64(maxCount))
			}

			if barLen < 1 && count > 0 {
				barLen = 1 // Au minimum 1 caractère si count > 0
			}

			bar := strings.Repeat(barChar, barLen)

			fmt.Printf("%3d [%10d] \033[32m%s\033[0m\n", bin, count, bar)

			linesPrinted++
		}
	}

	h.lastLines = linesPrinted
}
