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
	"encoding/json"

	"github.com/ercole-io/ercole/v2/model"
	ercutils "github.com/ercole-io/ercole/v2/utils"
)

// DbMounted marshals -action dbmounted output
func DbMounted(cmdOutput []byte, inst *model.MicrosoftSQLServerInstance) error {
	var out struct {
		Data struct {
			ServerName string `json:"servername"`
			StateDesc  string `json:"state_desc"`
			Platform   string `json:"platform"`
			DatabaseID int    `json:"database_id"`

			CollationName string `json:"collation_name"`
		} `json:"data"`
	}

	if err := json.Unmarshal(cmdOutput, &out); err != nil {
		return ercutils.NewError(err)
	}

	inst.ServerName = out.Data.ServerName
	inst.StateDesc = out.Data.StateDesc
	inst.Platform = out.Data.Platform
	inst.CollationName = out.Data.CollationName
	inst.Databases = []model.MicrosoftSQLServerDatabase{}
	inst.DatabaseID = out.Data.DatabaseID

	return nil
}
