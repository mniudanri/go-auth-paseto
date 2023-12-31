package util

import (
	"time"

	"github.com/mniudanri/go-auth-paseto/docs" // docs is auto generated by Swag CLI.
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// mapping fields from config
type Config struct {
	Host                 string        `mapstructure:"HOST"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// load config from path
func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to assemble config")
	}

	return config
}

func SetupSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "APIs Specification"
	docs.SwaggerInfo.Description = "This is APIs specification of sample golang project"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
