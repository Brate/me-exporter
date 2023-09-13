package config

type AuthEntry struct {
	Instance string `yaml:"instance"`
	Hash     string `yaml:"hash"`
	Url      string `yaml:"url"`
}
