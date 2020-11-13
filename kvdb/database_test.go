package kvdb

import "testing"

type testdriver struct {
	Nop
}

//Features return supported features
func (t testdriver) Features() Feature {
	return FeatureStore | FeatureTTLStore
}
func TestDatabase(t *testing.T) {
	db := New()
	db.Driver = testdriver{}
	if db.ShouldSupport(FeatureStore) != nil {
		t.Fatal()
	}
	if db.ShouldSupport(FeatureNext) != ErrFeatureNotSupported {
		t.Fatal()
	}
	if db.ShouldNotSupport(FeatureNext) != nil {
		t.Fatal()
	}
	if db.ShouldNotSupport(FeatureStore) != ErrFeatureSupported {
		t.Fatal()
	}
}
