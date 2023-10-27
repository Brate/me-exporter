package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	MeUrl          string
	Hash           string
	ExporterConfig *Config

	authPath       = flag.String("auth.file", "./auth.yaml", "Path to auth.yaml file.")
	ListenAddress  = flag.String("web.listen-address", ":10005", "Address on which to expose metrics and web interface.")
	RequestTimeout = flag.Int64("request.timeout", 5, "Timeout for HTTP requests to ME.")
	Workers        = flag.Int64("workers", 4, "Number of collector workers to start.")
)

type Config struct {
	Instances []AuthEntry `yaml:"auth"`
}

func init() {
	flag.Parse()
	LoadConfig()
}

func LoadConfig() {
	_ = LoadYAML(*authPath)
}

func LoadYAML(path string) (err error) {
	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}

	newConfig := &Config{}
	if err = yaml.Unmarshal(file, newConfig); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
		return
	}

	ExporterConfig = newConfig
	return
}
func (c *Config) FindAuthByInstance(instance string) *AuthEntry {
	for _, inst := range ExporterConfig.Instances {
		if inst.Instance == instance {
			return &inst
		}
	}
	return nil
}
