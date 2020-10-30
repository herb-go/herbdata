package kvdb

type Feature int64

func (f Feature) SupportAll(dst Feature) bool {
	return f&dst == f
}
func (f Feature) SupportAny(dst Feature) bool {
	return f&dst != 0
}

const (
	FeatureStore = Feature(1 << iota)
	FeatureTTLStore
	FeaturePersistent
	FeatureStable
	FeatureCounter
	FeatureTTLCounter
	FeatureNext
	FeaturePrev
	FeatureInsert
	FeatureUpdate
	FeatureTransaction
)
