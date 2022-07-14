package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for webapp config
type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`
		Grpc struct {
			Port int `yaml:"port"`
		} `yaml:"grpc"`
		Rest struct {
			Port int `yaml:"port"`
		} `yaml:"rest"`
	} `yaml:"server"`
	Postgres struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		Dbname string `yaml:"dbname"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
	} `yaml:"postgres"`
}

func New() (*Config, error) {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	// validate config data
	err = validateConfigData(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func parseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func validateConfigData(config *Config) error {
	// validate config data
	// get filter struct metadata
	if config.Env == "" {
		return errors.New("env is empty")
	}
	if config.Server.Host == "" {
		return errors.New("server.host is empty")
	}
	if config.Server.Grpc.Port == 0 {
		return errors.New("server.grpc.port is empty")
	}
	if config.Server.Rest.Port == 0 {
		return errors.New("server.rest.port is empty")
	}
	if config.Postgres.Host == "" {
		return errors.New("postgres.host is empty")
	}
	if config.Postgres.Port == 0 {
		return errors.New("postgres.port is empty")
	}
	if config.Postgres.Dbname == "" {
		return errors.New("postgres.dbname is empty")
	}
	if config.Postgres.User == "" {
		return errors.New("postgres.user is empty")
	}
	if config.Postgres.Pass == "" {
		return errors.New("postgres.pass is empty")
	}
	return nil
}
