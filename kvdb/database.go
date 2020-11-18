package kvdb

//Database key-value database struct
type Database struct {
	Driver
}

//ShouldSupport return ErrFeatureNotSupported if any given feature not be supported by driver
func (d *Database) ShouldSupport(features Feature) error {
	if !d.Driver.Features().SupportAll(features) {
		return ErrFeatureNotSupported
	}
	return nil
}

//ShouldNotSupport return ErrFeatureSupported if any given feature  be supported by driver
func (d *Database) ShouldNotSupport(features Feature) error {
	if d.Driver.Features().SupportAny(features) {
		return ErrFeatureSupported
	}
	return nil
}

//New create new database
func New() *Database {
	return &Database{}
}
