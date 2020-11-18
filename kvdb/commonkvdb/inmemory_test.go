package commonkvdb

import (
	"fmt"
	"testing"

	"github.com/herb-go/herbdata/kvdb"
	"github.com/herb-go/herbdata/kvdb/featuretestutil"
)

func TestInMemory(t *testing.T) {
	featuretestutil.TestDriver(func() kvdb.Driver {
		d, err := InMemoryFactory(nil)
		if err != nil {
			panic(err)
		}
		return d
	}, func(args ...interface{}) {
		fmt.Println(args...)
		panic("fatal")
	})
}
