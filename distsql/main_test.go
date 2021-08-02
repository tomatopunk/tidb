package distsql

import (
	"github.com/pingcap/tidb/util/testbridge"
	"go.uber.org/goleak"
	"testing"
)

func TestMain(m *testing.M) {
	//testbridge.WorkaroundGoCheckFlags()
	testbridge.WorkaroundGoCheckFlags()
	opts := []goleak.Option{
		goleak.IgnoreTopFunction("go.etcd.io/etcd/pkg/logutil.(*MergeLogger).outputLoop"),
	}
	goleak.VerifyTestMain(m, opts...)
}
