package core

// Config struct for the main configuration file
type Config struct {
	ChecksFolder string `yaml:"checks"`
	Notifiers    []ConfigNotifier
}

// ConfigCheck struct for a check parameters
type ConfigCheck struct {
	Title     string
	Notifier  string
	Processor ConfigProcessor
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
