package config

import (
	"golang-api/apiError"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, zerr := cfg.Build()
	Logger = logger

	apiError.HandleError(zerr)
	zap.ReplaceGlobals(logger)
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		// if no config was found just ignore the error, maybe the env is used
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			apiError.HandleError(err)
		}
	}
	viper.SetEnvPrefix("cloud_prefix")

	viper.AutomaticEnv()
}
