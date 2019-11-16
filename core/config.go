package core

// Config struct for the main configuration file
type Config struct {
	ChecksFolder string `yaml:"checks"`
	Delta        uint64
	Notifiers    []ConfigNotifier
}

// ConfigCheck struct for a check parameters
type ConfigCheck struct {
	Title     string
	Retries   uint
	Notifier  string
	Processor ConfigProcessor
}

// ConfigChecks struct for a checks file
type ConfigChecks struct {
	Checks []ConfigCheck
}

// ConfigNotifier struct for a notifier parameters
type ConfigNotifier struct {
	Name string
	Kind string
	Data map[string]string
}

// ConfigProcessor struct for a processor parameters
type ConfigProcessor struct {
	Kind string
	Data map[string]string
}
