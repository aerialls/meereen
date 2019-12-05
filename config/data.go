package config

var (
	// DefaultConfig is the default top-level configuration.
	DefaultConfig = Config{
		Global:  DefaultGlobalConfig,
		Metrics: DefaultMetricsConfig,
	}

	// DefaultGlobalConfig is the default global configuration.
	DefaultGlobalConfig = GlobalConfig{
		Retries: 3,
		Delta:   60,
	}

	// DefaultMetricsConfig is the default metrics configuration.
	DefaultMetricsConfig = MetricsConfig{
		Enable:  false,
		Address: "127.0.0.1:8337",
	}
)

// Config struct for the main configuration file
type Config struct {
	Global       GlobalConfig     `yaml:"global"`
	Metrics      MetricsConfig    `yaml:"metrics"`
	ChecksFolder string           `yaml:"checks_folder"`
	Notifiers    []NotifierConfig `yaml:"notifiers"`
}

// GlobalConfig struct for global parameters
type GlobalConfig struct {
	Delta   uint `yaml:"delta,omitempty"`
	Retries uint `yaml:"retries,omitempty"`
}

// MetricsConfig struct for Prometheus metrics parameters
type MetricsConfig struct {
	Enable  bool   `yaml:"enable,omitempty"`
	Address string `yaml:"address,omitempty"`
}

// ChecksConfig struct for a checks file
type ChecksConfig struct {
	Checks []CheckConfig `yaml:"checks"`
}

// CheckConfig struct for a check parameters
type CheckConfig struct {
	Title     string          `yaml:"title"`
	Notifier  string          `yaml:"notifier"`
	Processor ProcessorConfig `yaml:"processor"`
}

// NotifierConfig struct for a notifier parameters
type NotifierConfig struct {
	Name string            `yaml:"name"`
	Kind string            `yaml:"kind"`
	Data map[string]string `yaml:"data"`
}

// ProcessorConfig struct for a processor parameters
type ProcessorConfig struct {
	Kind string            `yaml:"kind"`
	Data map[string]string `yaml:"data"`
}
