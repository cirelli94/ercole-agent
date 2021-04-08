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

package marshal

import (
	"testing"

	"github.com/ercole-io/ercole/v2/model"
	"github.com/stretchr/testify/assert"
)

var testClusterMembershipStatusOutput string = `OracleClusterware: Y
VeritasClusterServer: N
VeritasClusterHostnames:
SunCluster: Y`

var testClusterMembershipStatusVeritas string = `OracleClusterware: N
VeritasClusterServer: Y
VeritasClusterHostnames: 0 sdlsts101;1 sdlsts102;2 sdlsts103 ;
SunCluster: N`

func TestClusterMembershipStatusOutput(t *testing.T) {
	testCases := []struct {
		output   string
		expected model.ClusterMembershipStatus
	}{
		{
			output: (testClusterMembershipStatusOutput),
			expected: model.ClusterMembershipStatus{
				OracleClusterware:    true,
				VeritasClusterServer: false,
				SunCluster:           true,
				HACMP:                false,
			},
		},
		{
			output: testClusterMembershipStatusVeritas,
			expected: model.ClusterMembershipStatus{
				OracleClusterware:       false,
				SunCluster:              false,
				HACMP:                   false,
				VeritasClusterServer:    true,
				VeritasClusterHostnames: []string{"sdlsts101", "sdlsts102", "sdlsts103"},
			},
		},
	}

	for _, tc := range testCases {
		actual := ClusterMembershipStatus([]byte(tc.output))
		assert.Equal(t, tc.expected, actual)
	}
}

func TestClusterMembershipStatusOutputShouldCrash(t *testing.T) {
	cmdOutput := []byte("pippo")

	assert.Panics(t, func() { ClusterMembershipStatus(cmdOutput) })
}
