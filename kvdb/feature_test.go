package kvdb

import "testing"

func TestFeature(t *testing.T) {
	if !(FeatureStore | FeatureTTLStore).SupportAll(FeatureTTLStore) {
		t.Fatal()
	}
	if !(FeatureStore | FeatureTTLStore).SupportAll(FeatureTTLStore | FeatureTTLStore) {
		t.Fatal()
	}
	if (FeatureStore | FeatureTTLStore).SupportAll(FeatureTTLStore | FeatureTTLStore | FeatureNext) {
		t.Fatal()
	}
	if (FeatureStore | FeatureTTLStore).SupportAny(FeatureNext) {
		t.Fatal()
	}
	if !(FeatureStore | FeatureTTLStore).SupportAny(FeatureStore | FeatureTTLStore) {
		t.Fatal()
	}
	if !(FeatureStore | FeatureTTLStore).SupportAny(FeatureTTLStore | FeatureTTLStore | FeatureNext) {
		t.Fatal()
	}
}
