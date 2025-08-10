package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go/v5/dns"
	"sigs.k8s.io/yaml"
)

type Config struct {
	ApiToken string                  `json:"api_token"`
	ZoneID   string                  `json:"zone_id"`
	TTL      dns.TTL                 `json:"ttl"`
	Comment  string                  `json:"comment"`
	Records  map[string]RecordConfig `json:"records"`
}

type RecordConfig struct {
	ApiToken string  `json:"api_token"`
	ZoneID   string  `json:"zone_id"`
	Proxied  bool    `json:"proxied"`
	TTL      dns.TTL `json:"ttl"`
	Comment  string  `json:"comment"`
}

func ParseConfig() (Config, error) {
	var config Config

	configPath := os.Getenv("CLOUDFLARE_DYNDNS_CONFIG_PATH")

	if configPath == "" {
		return config, errors.New("ConfigError: \"CLOUDFLARE_DYNDNS_CONFIG_PATH\" is not set")
	}

	configFile, err := os.ReadFile(configPath)

	if err != nil {
		if os.IsNotExist(err) {
			return config, fmt.Errorf("ConfigError: File does not exist at \"%s\"", configPath)
		}
		log.Fatal(err)
		return config, errors.New("ConfigError: Unable to read configuration file")
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		log.Fatal(err)
		return config, errors.New("ConfigError: Invalid configuration")
	}

	return config, nil
}
