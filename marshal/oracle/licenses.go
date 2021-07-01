// Copyright (c) 2019 Sorint.lab S.p.A.
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

package oracle

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/ercole-io/ercole-agent/v2/marshal"
	"github.com/ercole-io/ercole/v2/model"
	ercutils "github.com/ercole-io/ercole/v2/utils"
	"github.com/hashicorp/go-multierror"
)

// Licenses returns a list of licenses from the output of the licenses
// fetcher command.
func Licenses(cmdOutput []byte) ([]model.OracleDatabaseLicense, error) {
	var licenses []model.OracleDatabaseLicense
	var merr, err error

	scanner := bufio.NewScanner(bytes.NewReader(cmdOutput))
	for scanner.Scan() {
		license := new(model.OracleDatabaseLicense)
		line := scanner.Text()
		splitted := strings.Split(line, ";")

		if len(splitted) == 3 {
			key := strings.TrimSpace(splitted[0])
			value := strings.TrimSpace(splitted[1])
			value = strings.Replace(value, "\t", "", -1)

			license.Name = key

			if strings.TrimSpace(value) != "" {
				if license.Count, err = marshal.TrimParseFloat64(value); err != nil {
					merr = multierror.Append(merr, ercutils.NewError(err))
				}
			}

			licenses = append(licenses, *license)
		}
	}

	if merr != nil {
		return nil, merr
	}
	return licenses, nil
}
