No longer developed due to the following tool that does exactly what we need this one to do, this SDK is not very useful outside of a very specific use-case.

Much better SDK that we have gone ahead:
> https://github.com/721tools/stream-api-go

# OpenSea Stream API - golang SDK

A Golang SDK for receiving updates from the OpenSea Stream API - pushed over websockets. We currently support the following event types on a per-collection basis:

- item listed
- item sold
- item transferred
- item metadata updates
- item cancelled
- item received offer
- item received bid


Documentation: https://docs.opensea.io/reference/stream-api-overview

# Installation
This module requires Go 1.18 or later.
Standard `go get`:

```
go get github.com/harrys522/opensea-streamer
```

# Getting Started

## Authentication

In order to make onboarding easy, we've integrated the OpenSea Stream API with our existing API key system. The API keys you have been using for the REST API should work here as well. If you don't already have one, request an API key from us [here](https://docs.opensea.io/reference/request-an-api-key).

## Simple example



```golang
import (
    "fmt"
    openseaStream "github.com/harrys522/opensea-streamer"
)

func main() {
    client := openseaStream.NewStreamClient("wss://stream.openseabeta.com/socket", "api-key", phx.LogInfo, func(err error) {
        fmt.Println("NewStreamClient err:", err)
    })
    client.Connect()
    interesting_topics := []string{"item_listed", "item_transferred"}
    channels := client.Subscribe(interesting_topics)
    go func() {
		for {
			select {
			case msg := <-msgChannel:
				// Print the messages from the channel.
				fmt.Println(msg)
				break
			case <-ctx.Done():
				break
			}
		}
	}()
}
```

You can also optionally pass in:

- a `network` if you would like to access testnet networks.
    - The default value is `Network.MAINNET`, which represents the following blockchains: Ethereum, Polygon mainnet, Klaytn mainnet, and Solana mainnet
    - Can also select `Network.TESTNET`, which represents the following blockchains: Rinkeby, Polygon testnet (Mumbai), and Klaytn testnet (Baobab).
- `apiUrl` if you would like to access another OpenSea Stream API endpoint. Not needed if you provide a network or use the default values.
- an `onError` callback to handle errors. The default behavior is to `console.error` the error.
- a `logLevel` to set the log level. The default is `LogLevel.INFO`.

## Available Networks

The OpenSea Stream API is available on the following networks:

### Mainnet

`wss://stream.openseabeta.com/socket`

Mainnet supports events from the following blockchains: Ethereum, Polygon mainnet, Klaytn mainnet, and Solana mainnet.

### Testnet

`wss://testnets-stream.openseabeta.com/socket`

Testnet supports events from the following blockchains: Rinkeby, Polygon testnet (Mumbai), and Klaytn testnet (Baobab).

To create testnet instance of the client, you can create it with the following arguments:
