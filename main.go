// Copyright 2019 Martin Holst Swende
// This file is part of the archive_viewer library.
//
// The library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the goevmlab library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/holiman/archive_viewer/model"
	"github.com/holiman/archive_viewer/ui"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "filename")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, `
Reads the given archive, display in CLI user interface`)
	}
}

func main() {

	testTraces := []string{
		"/home/user/tools/archive-gitter/archive/ethereum/go-ethereum/private.json",
	}

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("Expected one argument\n")
		flag.Usage()
		os.Exit(1)
	}

	fName := flag.Arg(0)
	// Some debugging help here
	if n, err := strconv.Atoi(fName); err == nil {
		if n < len(testTraces) {
			fName = testTraces[n]
		}
	}
	room, err := model.RoomFromArchive(fName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	mgr := ui.NewViewManager(room)
	mgr.Run()
}
