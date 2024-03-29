package cmd

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/leebenson/conform"
	"github.com/mitchellh/go-homedir"
	"github.com/sanity-io/litter"
	"github.com/spf13/viper"
	"os"
)

// GlobalConfig the global config object
var GlobalConfig *config

func readGlobalConfig() {
	// Priority of configuration options
	// 1: CLI Parameters
	// 2: environment
	// 3: config.yaml
	// 4: defaults
	config, err := readConfig()
	if err != nil {
		panic(err.Error())
	}

	// Set config object for main package
	GlobalConfig = config
}

var defaultConfig = &config{
	Debug:            false,
	BackendAddr:      "http://localhost:8000",
	AuthEnabled:      false,
	ReverseProxyAddr: "http://localhost:9000",
}

// configInit must be called from the packages' init() func
func configInit() error {
	cliFlags()
	return bindFlagsAndEnv()
}

// Create private data struct to hold config options.
// `mapstructure` => viper tags
// `struct` => fatih structs tag
// `env` => environment variable name
type config struct {
	ReverseProxyAddr string `mapstructure:"proxy" structs:"proxy" env:"MOAR_S3_PROXY_URL"`
	BackendAddr      string `mapstructure:"addr" structs:"addr" env:"MOAR_BACKEND_ADDR"`
	Debug            bool   `mapstructure:"debug" structs:"debug" env:"MOAR_DEBUG"`
	AuthEnabled      bool   `mapstructure:"auth_enabled" structs:"auth_enabled" env:"MOAR_AUTH_ENABLED"`

	// sensitive
	AccessToken string `mapstructure:"access_token" structs:"access_token" env:"MOAR_ACCESS_TOKEN" conform:"redact"`
}

// cliFlags defines cli parameters for all config options
func cliFlags() {
	// Keep cli parameters in sync with the config struct
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.moar.yaml)")
	rootCmd.PersistentFlags().String("addr", defaultConfig.BackendAddr, "The backend service address")
	rootCmd.PersistentFlags().StringP("proxy", "p", defaultConfig.ReverseProxyAddr, "The reverse proxy which is directed at the module storage.")
	rootCmd.PersistentFlags().BoolP("debug", "d", defaultConfig.Debug, "Toggles whether debug logs are enabled")
	rootCmd.PersistentFlags().Bool("auth_enabled", defaultConfig.AuthEnabled, "Whether the backend service should use authentication")
	rootCmd.PersistentFlags().String("access_token", defaultConfig.AccessToken, "Access token")
}

// bindFlagsAndEnv will assign the environment variables to the cli parameters
func bindFlagsAndEnv() (err error) {
	for _, field := range structs.Fields(&config{}) {
		// Get the struct tag values
		key := field.Tag("structs")
		env := field.Tag("env")

		// Bind cobra flags to viper
		err = viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))
		if err != nil {
			return err
		}
		err = viper.BindEnv(key, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// Print the config object
// but remove sensitive data
func (c *config) Print() {
	cp := *c
	_ = conform.Strings(&cp)
	litter.Dump(cp)
}

// String the config object
// but remove sensitive data
func (c *config) String() string {
	cp := *c
	_ = conform.Strings(&cp)
	return litter.Sdump(cp)
}

// readConfig a helper to read default from a default config object.
func readConfig() (*config, error) {
	// Create a map of the default config
	defaultsAsMap := structs.Map(defaultConfig)

	// Set defaults
	for key, value := range defaultsAsMap {
		viper.SetDefault(key, value)
	}

	// Read config from file
	viper.SetConfigName(".moar")
	viper.AddConfigPath(".")
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".moar" (without extension).
	viper.AddConfigPath(home)
	if err = viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Unmarshal config into struct
	c := &config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
