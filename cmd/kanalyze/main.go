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

// main package
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Provisioned by ldflags.
var (
	name      string //nolint: gochecknoglobals
	version   string //nolint: gochecknoglobals
	commit    string //nolint: gochecknoglobals
	buildDate string //nolint: gochecknoglobals
	builtBy   string //nolint: gochecknoglobals
)

const refreshInterval = 500 * time.Millisecond

func main() {
	//nolint: exhaustruct
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("%v %v (commit=%v date=%v by=%v)", name, version, commit, buildDate, builtBy)

	driver := NewCounter()
	histogram := NewHistogram()

	ticker := time.NewTicker(refreshInterval)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			histogram.Render()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)    //nolint:mnd
	scanner.Buffer(buf, 1024*1024*100) //nolint:mnd // increase buffer up to 100 MB

	for scanner.Scan() {
		line := scanner.Bytes()
		histogram.Update(driver.Ingest(line))
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("error reading standard input") //nolint:gocritic
	}

	histogram.Render()
	fmt.Println()
}
