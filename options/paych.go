package options

type PaychSettings struct {
	From     string
	GasLimit string
	GasPrice string
	Preview  bool
}

type PaychVoucherSettings struct {
	From    string
	Validat int64
}

type PaychOption func(*PaychSettings) error
type PaychVoucherOption func(*PaychVoucherSettings) error

func PaychOptions(opts ...PaychOption) (*PaychSettings, error) {
	options := &PaychSettings{
		Preview: false,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func PaychVoucherOptions(opts ...PaychVoucherOption) (*PaychVoucherSettings, error) {
	options := &PaychVoucherSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type paychOpts struct{}

var Paych paychOpts

func (paychOpts) From(addr string) PaychOption {
	return func(settings *PaychSettings) error {
		settings.From = addr
		return nil
	}
}

func (paychOpts) GasLimit(gaslimit string) PaychOption {
	return func(settings *PaychSettings) error {
		settings.GasLimit = gaslimit
		return nil
	}
}

func (paychOpts) GasPrice(gasprice string) PaychOption {
	return func(settings *PaychSettings) error {
		settings.GasPrice = gasprice
		return nil
	}
}

func (paychOpts) Preview(preview bool) PaychOption {
	return func(settings *PaychSettings) error {
		settings.Preview = preview
		return nil
	}
}

func (paychOpts) PaychVoucherFrom(addr string) PaychVoucherOption {
	return func(settings *PaychVoucherSettings) error {
		settings.From = addr
		return nil
	}
}

func (paychOpts) Validat(validat int64) PaychVoucherOption {
	return func(settings *PaychVoucherSettings) error {
		settings.Validat = validat
		return nil
	}
}
