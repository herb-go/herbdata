package commonkvdb

import (
	"testing"

	"github.com/herb-go/herbdata/kvdb"
	"github.com/herb-go/herbdata/kvdb/featuretestutil"
)

func TestInMemory(t *testing.T) {
	featuretestutil.TestDriver(func() kvdb.Driver { return NewInMemory() }, t.Fatal)
}
