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
	FeatureTTL
	FeaturePersistent
	FeatureStable
	FeatureTransaction
	FeatureCounter
	FeatureNext
	FeaturePrev
	FeatureInsert
	FeatureUpdate
)
