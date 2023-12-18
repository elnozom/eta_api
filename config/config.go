// package config

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// // Config func to get env value from key ---
// func Config(key string) string {
// 	// load .env file
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Print("Error loading .env file")
// 	}
// 	return os.Getenv(key)

// }
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type StateConfig struct {
	State string `mapstructure:"STATE"`
}

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	State             string `mapstructure:"STATE"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	ClientID          string `mapstructure:"CLIENT_ID"`
	ClientSecret      string `mapstructure:"CLIENT_SECRET"`
	POSClientID       string `mapstructure:"POS_CLIENT_ID"`
	POSClientSecret   string `mapstructure:"POS_CLIENT_SECRET"`
	SignerUrl         string `mapstructure:"SIGNER_URL"`
	AuthApiUrl        string `mapstructure:"AUTH_API_URL"`
	SubmitApiUrl      string `mapstructure:"SUBMIT_API_URL"`
	PosSerial         string `mapstructure:"POS_SERIAL"`
	PosOsVersion      string `mapstructure:"POS_OS_VERSION"`
	Port              string `mapstructure:"PORT"`

	PreSharedKey string `mapstructure:"PRESHARED_KEY"`
}

func loadAncScan(config *Config) (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig reads configuration from file or environment variables.
func LoadState(path string) (config StateConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("state")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, state string) (config Config, err error) {
	stateEnvFilePath := fmt.Sprintf("%s.env", state)
	viper.SetConfigName(stateEnvFilePath)
	err = loadAncScan(&config)
	if err != nil {
		return
	}
	viper.SetConfigName("shared.env")
	err = loadAncScan(&config)
	if err != nil {
		return
	}
	return
}
