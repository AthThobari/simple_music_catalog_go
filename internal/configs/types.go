package configs

type (
	Config struct {
		Service Service `mapstructure:"service"`
		Database DatabaseConfig`mapstructure:"database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
		SecretKey string `mapstructure:"secretKey"`
	}

	DatabaseConfig struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)