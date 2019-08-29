package options

type DealsRedeemSettings struct {
	GasLimit string
	GasPrice string
}

type DealsOption func(*DealsRedeemSettings) error

func DealsOptions(opts ...DealsOption) (*DealsRedeemSettings, error) {
	options := &DealsRedeemSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type dealsOpts struct{}

var Deals dealsOpts

func (dealsOpts) GasLimit(gaslimit string) DealsOption {
	return func(settings *DealsRedeemSettings) error {
		settings.GasLimit = gaslimit
		return nil
	}
}

func (dealsOpts) GasPrice(gasprice string) DealsOption {
	return func(settings *DealsRedeemSettings) error {
		settings.GasPrice = gasprice
		return nil
	}
}
