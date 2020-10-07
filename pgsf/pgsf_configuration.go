package pgsf

// PgsfConfiguration is used to configure the Pgsf instance.
type Configuration struct {
	Name       string
	MaxClients int
	Port       int
	ListenUrl  string
}

// Gets a default configuration, used for various purposes.
func GetDefaultConfiguration() Configuration {
	return Configuration{
		Name:       "pgsf-server",
		MaxClients: 100,
		Port:       8083,
		ListenUrl:  "/server/",
	}
}
