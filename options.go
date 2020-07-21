package htmlselector

// OptionConfig ...
type OptionConfig struct {
	skipEmptyTags bool
}

// Option ...
type Option func(config *OptionConfig)

// SkipEmptyTags ...
func SkipEmptyTags() Option {
	return func(config *OptionConfig) {
		config.skipEmptyTags = true
	}
}

func newOptionConfig(options []Option) OptionConfig {
	var config OptionConfig
	for _, option := range options {
		option(&config)
	}

	return config
}
