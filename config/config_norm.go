//+build !rhel5

// Copyright (c) 2020 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetBaseDir return executable base directory, os independant
func GetBaseDir() string {
	var s string
	if runtime.GOOS == "windows" {
		s, _ = os.Executable()
		s = filepath.Dir(s)
	} else {
		s, _ = os.Readlink("/proc/self/exe")
		s = filepath.Dir(s)
	}

	return s
}
