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

// Features marshals -action features output
func Features(cmdOutput []byte) ([]model.MicrosoftSQLServerProductFeature, error) {
	var rawOut struct {
		Data []struct {
			Product    string
			Feature    string
			Language   string
			Edition    string
			Version    string
			Clustered  string
			Configured string
		} `json:"data"`
	}

	if err := json.Unmarshal(cmdOutput, &rawOut); err != nil {
		return nil, ercutils.NewError(err)
	}

	out := make([]model.MicrosoftSQLServerProductFeature, len(rawOut.Data))

	for i, v := range rawOut.Data {
		out[i].Product = v.Product
		out[i].Feature = v.Feature
		out[i].Language = v.Language
		out[i].Edition = v.Edition
		out[i].Version = v.Version
		out[i].Clustered = v.Clustered == "Yes"
		out[i].Configured = v.Clustered == "Yes"
	}

	return out, nil
}
