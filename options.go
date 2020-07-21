package htmlselector

// OptionConfig ...
type OptionConfig struct {
}

// Option ...
type Option func(config *OptionConfig)

func newOptionConfig(options []Option) OptionConfig {
	var config OptionConfig
	for _, option := range options {
		option(&config)
	}

	return config
}
