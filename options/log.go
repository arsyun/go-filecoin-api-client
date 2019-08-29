package options

type LogLevelSettings struct {
	Subsystem string
}

type LogLevelOption func(*LogLevelSettings) error

func LogLevelOptions(opts ...LogLevelOption) (*LogLevelSettings, error) {
	options := &LogLevelSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type logOpts struct{}

var Log logOpts

func (logOpts) Subsystem(subsystem string) LogLevelOption {
	return func(settings *LogLevelSettings) error {
		settings.Subsystem = subsystem
		return nil
	}
}
