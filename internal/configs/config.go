package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFolders []string
	configFile string
	configType string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultConfigFolder(),
		configFile: getDefaultConfigFile(),
		configType: getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, configFolder := range opt.configFolders {
		viper.AddConfigPath(configFolder)
	}
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

type Option func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}