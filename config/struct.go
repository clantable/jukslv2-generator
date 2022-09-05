package config

type Config struct {
	ServiceName      string           `yaml:"ServiceName"`
	DomainConfig     DomainConfig     `yaml:"DomainConfig"`
	RepositoryConfig RepositoryConfig `yaml:"RepositoryConfig"`
	UseCaseConfig    UsecaseConfig    `yaml:"UsecaseConfig"`
	HandlerConfig    HandlerConfig    `yaml:"HandlerConfig"`
}

type DomainConfig struct {
	Path    string `yaml:"Path"`
	Package string `yaml:"Package"`
}

type RepositoryConfig struct {
	Path    string `yaml:"Path"`
	Package string `yaml:"Package"`
}
type UsecaseConfig struct {
	Path    string `yaml:"Path"`
	Package string `yaml:"Package"`
}
type HandlerConfig struct {
	Path    string `yaml:"Path"`
	Package string `yaml:"Package"`
}
