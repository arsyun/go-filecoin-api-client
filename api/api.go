package api

import (
	iface "go-filecoin-api-client/interface-go-filecoin"
	"net/http"
	gohttp "net/http"
)

type HttpApi struct {
	url     string
	httpcli gohttp.Client
	Headers http.Header
}

func NewAPI(url string) *HttpApi {
	if url == "" {
		url = "127.0.0.1:3453"
	}
	return &HttpApi{
		url: url,
	}
}

func (api *HttpApi) Request(command string, args ...string) RequestBuilder {
	headers := make(map[string]string)
	if api.Headers != nil {
		for k := range api.Headers {
			headers[k] = api.Headers.Get(k)
		}
	}
	return &requestBuilder{
		command: command,
		args:    args,
		shell:   api,
		headers: headers,
	}
}

func (api *HttpApi) Actor() iface.ActorAPI {
	return (*ActorAPI)(api)
}

func (api *HttpApi) Address() iface.AddressAPI {
	return (*AddressAPI)(api)
}

func (api *HttpApi) Basic() iface.BasicAPI {
	return (*BasicAPI)(api)
}

func (api *HttpApi) BitSwap() iface.BitSwapAPI {
	return (*BitSwapAPI)(api)
}

func (api *HttpApi) BootStrap() iface.BootStrapAPI {
	return (*BootStrapAPI)(api)
}

func (api *HttpApi) Chain() iface.ChainAPI {
	return (*ChainAPI)(api)
}

func (api *HttpApi) Client() iface.ClientAPI {
	return (*ClientAPI)(api)
}

func (api *HttpApi) Config() iface.ConfigAPI {
	return (*ConfigAPI)(api)
}

func (api *HttpApi) Deals() iface.DealsAPI {
	return (*DealsAPI)(api)
}

func (api *HttpApi) Dag() iface.DagAPI {
	return (*DagAPI)(api)
}

func (api *HttpApi) DHT() iface.DhtAPI {
	return (*DhtAPI)(api)
}

func (api *HttpApi) Inspect() iface.InspectAPI {
	return (*InspectAPI)(api)
}

func (api *HttpApi) Log() iface.LogAPI {
	return (*LogAPI)(api)
}

func (api *HttpApi) Message() iface.MessageAPI {
	return (*MessageAPI)(api)
}

func (api *HttpApi) Paych() iface.PaychAPI {
	return (*PaychAPI)(api)
}

func (api *HttpApi) RetrievalClient() iface.RetrievalClientAPI {
	return (*RetrievalClientAPI)(api)
}

func (api *HttpApi) Show() iface.ShowAPI {
	return (*ShowAPI)(api)
}

func (api *HttpApi) Stats() iface.StatsAPI {
	return (*StatsAPI)(api)
}

func (api *HttpApi) Swarm() iface.SwarmAPI {
	return (*SwarmAPI)(api)
}

func (api *HttpApi) Miner() iface.MinerAPI {
	return (*MinerAPI)(api)
}

func (api *HttpApi) Mining() iface.MiningAPI {
	return (*MiningAPI)(api)
}

func (api *HttpApi) Mpool() iface.MpoolAPI {
	return (*MpoolAPI)(api)
}

func (api *HttpApi) Wallet() iface.WalletAPI {
	return (*WalletAPI)(api)
}
