package options

type MessageSendSettings struct {
	From     string
	GasLimit string
	GasPrice string
	Value    int64
}

type MessageWaitSettings struct {
	Message bool
	Receipt bool
	Return  bool
	Timeout string
}

type MessageSendOption func(*MessageSendSettings) error
type MessageWaitOption func(*MessageWaitSettings) error

func MessageSendOptions(opts ...MessageSendOption) (*MessageSendSettings, error) {
	options := &MessageSendSettings{
		Value: 0,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func MessageWaitOptions(opts ...MessageWaitOption) (*MessageWaitSettings, error) {
	options := &MessageWaitSettings{
		Message: true,
		Receipt: true,
		Return:  false,
		Timeout: "10m",
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type messageOpts struct{}

var Message messageOpts

func (messageOpts) From(from string) MessageSendOption {
	return func(settings *MessageSendSettings) error {
		settings.From = from
		return nil
	}
}

func (messageOpts) GasLimit(gaslimit string) MessageSendOption {
	return func(settings *MessageSendSettings) error {
		settings.GasLimit = gaslimit
		return nil
	}
}

func (messageOpts) GasPrice(gasprice string) MessageSendOption {
	return func(settings *MessageSendSettings) error {
		settings.GasPrice = gasprice
		return nil
	}
}

func (messageOpts) Value(value int64) MessageSendOption {
	return func(settings *MessageSendSettings) error {
		settings.Value = value
		return nil
	}
}

func (messageOpts) ShowMessage(msg bool) MessageWaitOption {
	return func(settings *MessageWaitSettings) error {
		settings.Message = msg
		return nil
	}
}

func (messageOpts) ShowReceipt(receipt bool) MessageWaitOption {
	return func(settings *MessageWaitSettings) error {
		settings.Receipt = receipt
		return nil
	}
}

func (messageOpts) IsReturn(re bool) MessageWaitOption {
	return func(settings *MessageWaitSettings) error {
		settings.Return = re
		return nil
	}
}

func (messageOpts) TimeOut(timeout string) MessageWaitOption {
	return func(settings *MessageWaitSettings) error {
		settings.Timeout = timeout
		return nil
	}
}
