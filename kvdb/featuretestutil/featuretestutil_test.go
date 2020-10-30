package featuretestutil

import (
	"testing"

	"github.com/herb-go/herbdata/kvdb"
)

func TestTester(t *testing.T) {
	var result interface{}
	tester := &Tester{
		Hanlder: func(v ...interface{}) {
			result = v[0]
		},
	}
	tester.Assert(true, "fatal")
	if result != nil {
		t.Fatal()
	}
	tester.Assert(false, "fatal")
	if result != "fatal" {
		t.Fatal()
	}
}
func TestUtil(t *testing.T) {
	TestDriver(kvdb.NewMapStore(), t.Fatal)
}
