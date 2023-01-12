package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	PostgresMaxConn  int32
}

func Load() Config {

	cfg := Config{}

	cfg.HTTPPort = ":8800"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "ravshanbek"
	cfg.PostgresDatabase = "catalog"
	cfg.PostgresPassword = "8990"
	cfg.PostgresPort = "5432"
	cfg.PostgresMaxConn = 30

	return cfg
}
