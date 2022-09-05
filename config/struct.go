package config

type Config struct {
	ServiceName      string
	DomainConfig     DomainConfig
	RepositoryConfig RepositoryConfig
	UsecaseConfig    UsecaseConfig
	DeliveryConfig   DeliveryConfig
}

type DomainConfig struct {
	Path    string
	Package string
}

type RepositoryConfig struct {
	Path    string
	Package string
}
type UsecaseConfig struct {
	Path    string
	Package string
}
type DeliveryConfig struct {
	Path            string
	Package         string
	ConverterConfig struct {
		Path    string
		Package string
	}
}
