package kvdb

type Feature int64

const (
	FeatureKeyValue = Feature(1 << iota)
	FeaturePersistent
	FeatureReverseIter
	FeatureIter
	FeatureTTL
	FeatureAtomic
	FeatureTransaction
	FeatureCounter
)
