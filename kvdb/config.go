package kvdb

type Config struct {
	Driver string
	Config func(v interface{}) error `config:", lazyload"`
}

func (c *Config) ApplyTo(db *Database) error {
	return Apply(db, c.Driver, c.Config)
}

func Apply(db *Database, driver string, loader func(interface{}) error) error {
	d, err := NewDriver(driver, loader)
	if err != nil {
		return err
	}
	db.Driver = d
	return nil
}
