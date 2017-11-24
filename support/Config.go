package support

type Config struct {
	Database string
}

func (c *Config) getDBStringConnector() string{
	return c.Database
}

