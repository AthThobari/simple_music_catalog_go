package configs

type (
	Config struct {
		Service Service `mapstructure:"service"`
		Database DatabaseConfig`mapstructure:"database"`
		SpotifyConfig SpotifyConfig
	}

	Service struct {
		Port string `mapstructure:"port"`
		SecretKey string `mapstructure:"secretKey"`
	}

	DatabaseConfig struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientID string
		ClientSecret string
	}
)