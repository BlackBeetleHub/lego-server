package support

type Config struct {
	Database   string
	ListenIP   string
	ListenPort string
}

func (c *Config) GetDBStringConnector() string {
	return c.Database
}
