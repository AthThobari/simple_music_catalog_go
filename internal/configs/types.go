package configs

type (
	Config struct {
		Service Service `mapstructure:"service"`
		Database DatabaseConfig`mapstructure:"database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}

	DatabaseConfig struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)