// Author: slowpoke <mail plus git at slowpoke dot io>
// Repository: https://github.com/proxypoke/i3ipc
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License.
// It comes without any warranty, to the extent permitted by
// applicable law. For a copy of the license, see COPYING or
// head to http://sam.zoy.org/wtfpl/COPYING.

package i3ipc

import (
	"testing"
)

func TestGetTree(t *testing.T) {
	ipc, _ := GetIPCSocket()

	//root, err := GetTree(ipc)
	_, err := ipc.GetTree()
	if err != nil {
		t.Errorf("Getting tree failed: %v", err)
	}
}
