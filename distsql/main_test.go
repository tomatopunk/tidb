package distsql

import (
	"github.com/pingcap/tidb/util/testbridge"
	"testing"
)

func TestMain(m *testing.M) {
	testbridge.WorkaroundGoCheckFlags()
}
