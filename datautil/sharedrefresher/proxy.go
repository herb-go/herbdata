package sharedrefresher

//SharedRefresher shared reffresher interface
type SharedRefresher interface {
	//RefreshShared refresh shared data.
	//New data what different from old should be returned
	RefreshShared(old []byte) ([]byte, error)
}
