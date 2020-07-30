package htmlselector

// OptionConfig ...
type OptionConfig struct {
	skipEmptyTags       bool
	skipEmptyAttributes bool
}

// Option ...
type Option func(config *OptionConfig)

// SkipEmptyTags ...
func SkipEmptyTags() Option {
	return func(config *OptionConfig) {
		config.skipEmptyTags = true
	}
}

// SkipEmptyAttributes ...
func SkipEmptyAttributes() Option {
	return func(config *OptionConfig) {
		config.skipEmptyAttributes = true
	}
}

func newOptionConfig(options []Option) OptionConfig {
	var config OptionConfig
	for _, option := range options {
		option(&config)
	}

	return config
}
