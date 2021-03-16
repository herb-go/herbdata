package kvdb

//Feature key-value database feature type
type Feature int64

//SupportAll check if all given features are supported
func (f Feature) SupportAll(dst Feature) bool {
	return f&dst == dst
}

//SupportAny check if any one of given features are supported
func (f Feature) SupportAny(dst Feature) bool {
	return f&dst != 0
}

//FeaturesSetEmpty empty feature set
const FeaturesSetEmpty = Feature(0)

const (
	//FeatureStore key-value database store(Set/Get/Delete) feature
	FeatureStore = Feature(1 << iota)
	//FeatureTTLStore key-value database ttl store(SetWithTTL/Get/Delete) feature
	FeatureTTLStore
	//FeatureExpiredStore key-value database expired store(SetWithExpired/Get/Delete) feature
	FeatureExpiredStore
	//FeatureCounter key-value database counter(SetCounter/IncreaseCounter/GetCounter/DeleteCounter) feature
	FeatureCounter
	//FeatureTTLCounter key-value database counter(SetCounterWithTTL/IncreaseCounterWithTTL/GetCounter/DeleteCounter) feature
	FeatureTTLCounter
	//FeatureNext key-value database next (Next) feature
	FeatureNext
	//FeaturePrev key-value database prev (Prev) feature
	FeaturePrev
	//FeatureInsert key-value database insert (Insert) feature
	FeatureInsert
	//FeatureTTLInsert key-value database ttl insert (InsertWithTTL) feature
	FeatureTTLInsert
	//FeatureUpdate key-value database update (Update) feature
	FeatureUpdate
	//FeatureTTLUpdate key-value database ttl update (UpdateWithTTL) feature
	FeatureTTLUpdate
	//FeatureTransaction key-value database transactio(Begin) feature
	FeatureTransaction
	//FeatureNonpersistent  if data will be drop after application restart
	FeatureNonpersistent
	//FeatureUnstable if data may be droped if needed.
	FeatureUnstable
	//FeatureEmbedded if database is embedded.
	FeatureEmbedded
)
