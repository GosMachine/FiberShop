package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Clients    ClientsConfig `yaml:"clients"`
	AppSecret  string        `yaml:"appSecret" env-required:"true" env:"SECRET"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

type Client struct {
	Address      string `yaml:"address"`
	Timeout      int    `yaml:"timeout"`
	RetriesCount int    `yaml:"retries_count"`
}

type ClientsConfig struct {
	Auth Client `yaml:"auth"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist" + path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
