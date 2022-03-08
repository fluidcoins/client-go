package fluidcoins

import (
	"errors"

	"github.com/fluidcoins/client-go/util"
	"github.com/spf13/viper"
)

// APIAuth holds the API_KEY environment variable loaded by Viper
type APIAuth struct {
	Apikey string `mapstructure:"API_KEY"`
}

// SetupEnv is used to configure Viper
type SetupEnv struct {
	Path      string
	Filename  string
	Extension string
}

// EnvSettings is a configuration type
type EnvSettings func(*SetupEnv)

// SetPath is an EnvSetting type that allows to set the path to folder where environment variable API_KEY is stored.
// If it's not provided, path defaults to "../" which is the root
func SetPath(path string) EnvSettings {
	return func(se *SetupEnv) {
		se.Path = path
	}
}

// SetFileName is an EnvSetting type that allows to set the filename where environment variable API_KEY is stored.
// If it's not provided, filename defaults to "app"
func SetFileName(name string) EnvSettings {
	return func(se *SetupEnv) {
		se.Filename = name
	}
}

// SetFileName is an EnvSetting type that allows to set the file extension( e.g env, yaml ) where environment variable API_KEY is stored.
// If it's not provided, filename defaults to "env"
func SetExtension(extension string) EnvSettings {
	return func(se *SetupEnv) {
		se.Extension = extension
	}
}

func setDefaultEnvSettings(se *SetupEnv) {
	se.Path = "../"
	se.Filename = "app"
	se.Extension = "env"
}

// LoadAPIKey loads generated API key from the environment variables for
// authenticating API calls.
func LoadAPIKey(opts ...EnvSettings) (string, error) {
	s := new(SetupEnv)

	setDefaultEnvSettings(s)

	for _, o := range opts {
		o(s)
	}

	viper.AddConfigPath(s.Path)
	viper.SetConfigName(s.Filename)
	viper.SetConfigType(s.Extension)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	a := new(APIAuth)
	err = viper.Unmarshal(a)
	if err != nil {
		return "", err
	}

	if util.IsStringEmpty(a.Apikey) {
		return "", errors.New("environment variable API_KEY is empty")
	}

	return a.Apikey, nil
}
