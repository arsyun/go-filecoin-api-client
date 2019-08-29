package options

type MinerCreateSettings struct {
	Sectorsize int64
	From       string
	PeerId     string
	GasPrice   string
	GasLimit   string
}

type MinerSetPriceSettings struct {
	MinerAddr string
	GasPrice  string
	GasLimit  string
}

type MinerUpdatePeerIdSettings struct {
	GasPrice string
	GasLimit string
}

type MinerCreateOption func(*MinerCreateSettings) error
type MinerSetPriceOption func(*MinerSetPriceSettings) error
type MinerUpdatePeerIdOption func(*MinerUpdatePeerIdSettings) error

func MinerCreateOptions(opts ...MinerCreateOption) (*MinerCreateSettings, error) {
	options := &MinerCreateSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func MinerSetPriceOptions(opts ...MinerSetPriceOption) (*MinerSetPriceSettings, error) {
	options := &MinerSetPriceSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func MinerUpdatePeerIdOptions(opts ...MinerUpdatePeerIdOption) (*MinerUpdatePeerIdSettings, error) {
	options := &MinerUpdatePeerIdSettings{}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type minerOpts struct{}

var Miner minerOpts

func (minerOpts) Sectorsize(sectorsize int64) MinerCreateOption {
	return func(settings *MinerCreateSettings) error {
		settings.Sectorsize = sectorsize
		return nil
	}
}

func (minerOpts) From(addr string) MinerCreateOption {
	return func(settings *MinerCreateSettings) error {
		settings.From = addr
		return nil
	}
}

func (minerOpts) PeerId(peerid string) MinerCreateOption {
	return func(settings *MinerCreateSettings) error {
		settings.PeerId = peerid
		return nil
	}
}

func (minerOpts) CreateMinerGasLimit(gaslimit string) MinerCreateOption {
	return func(settings *MinerCreateSettings) error {
		settings.GasLimit = gaslimit
		return nil
	}
}

func (minerOpts) CreateMinerGasPrice(gasprice string) MinerCreateOption {
	return func(settings *MinerCreateSettings) error {
		settings.GasPrice = gasprice
		return nil
	}
}

func (minerOpts) MinerAddr(addr string) MinerSetPriceOption {
	return func(settings *MinerSetPriceSettings) error {
		settings.MinerAddr = addr
		return nil
	}
}

func (minerOpts) SetPriceGasPrice(gasprice string) MinerSetPriceOption {
	return func(settings *MinerSetPriceSettings) error {
		settings.GasPrice = gasprice
		return nil
	}
}

func (minerOpts) SetPriceGasLimit(gaslimit string) MinerSetPriceOption {
	return func(settings *MinerSetPriceSettings) error {
		settings.GasLimit = gaslimit
		return nil
	}
}

func (minerOpts) UpPeerIdGasPrice(upprice string) MinerUpdatePeerIdOption {
	return func(settings *MinerUpdatePeerIdSettings) error {
		settings.GasPrice = upprice
		return nil
	}
}

func (minerOpts) UpPeerIdGasLimit(uplimit string) MinerUpdatePeerIdOption {
	return func(settings *MinerUpdatePeerIdSettings) error {
		settings.GasLimit = uplimit
		return nil
	}
}
