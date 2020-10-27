package kvdb

type Feature int64

const (
	FeatureCore = Feature(1 << iota)
	FeatureIter
	FeatureTTL
	FeatureTransaction
)
