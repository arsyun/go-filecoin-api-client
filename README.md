# go-filecoin-api-client

> An API Client For Filecoin

⚠️ WARNING: Filecoin is under heavy development and breaking changes are highly likely between versions. This library is experimental, It may be broken in part or in entirety.

🧩 Filecoin version: **Filecoin 0.4.3**

## Install

```sh
go get "xxx"
```

## Usage

### Golang

```golang
func main() {
	//default url: 127.0.0.1:3453
	AddressAPI := api.NewAPI("").Address()
	addrs, err := AddressAPI.Ls(context.Background())
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("addrs:", addrs)
	return
}
```

## API

Because of the time, the functions provided by the current version of filecoin will be listed, and the unrealized functions will be marked for further improvement.

* [actor.ls](API.md#actorls)

* [address.default]()

* [address.ls]()

* [address.lookup]()

* [address.new]()

* [bitswap.stats]()

* [bootstrap.ls]()

* [chain.head]()

* [chain.ls]()

* [client.cat]()

* [client.import]()
* [client.list-asks]()
* [client.payments]()
* client.propose-storage-deal
* client.query-storage-deal
* client.verify-storage-deal
* [config.get]()
* [config.set]()
* [dag.get]()
* [deals.list]()
* [deals.redeem]()
* [deals.show]()
* [dht.findpeer]()
* dht.findprovs
* dht.query
* [id]()
* [inspect.all]()
* [inspect.config]()
* [inspect.disk]()
* [inspect.environment]()
* [inspect.memory]()
* [inspect.runtime]()
* [log.level]()
* [log.ls]()
* log.tail
* [message.send]()
* [message.status]()
* [message.wait]()
* [miner.collateral]()
* miner.create
* [miner.owner]()
* [miner.power]()
* [miner.proving-period]()
* [miner.setprice]()
* [miner.update-peerid]()
* [mining.address]()
* mining.once
* [mining.seal-now]()
* [mining.start]()
* [mining.status]()
* [mining.stop]()
* [mpool.ls]()
* [mpool.rm]()
* [mpool.show]()
* [paych.cancel]()
* [paych.close]()
* [paych.create]()
* [paych.ls]()
* [paych.extend]()
* [paych.reclaim]()
* [paych.redeem]()
* [paych.voucher]()
* ping
* retrievalclient.retrieve-piece
* [show.block]()
* [show.header]()
* [show.messages]()
* [show.receipts]()
* [stats.bandwith]()
* [swarm.connect]()
* [swarm.peers]()
* [version]()
* [wallet.balance]()
* [wallet.export]()
* [wallet.import]()