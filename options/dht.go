package options

type DhtFindProvsSettings struct {
	NumProviders int
}

type DhtFindProvsOption func(*DhtFindProvsSettings) error

func DhtProvsOptions(opts ...DhtFindProvsOption) (*DhtFindProvsSettings, error) {
	options := &DhtFindProvsSettings{
		NumProviders: 1,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type dhtOpts struct{}

var DHT dhtOpts

func (dhtOpts) NumProviders(num int) DhtFindProvsOption {
	return func(settings *DhtFindProvsSettings) error {
		settings.NumProviders = num
		return nil
	}
}
