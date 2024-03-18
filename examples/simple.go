package main

import (
	"fmt"

	"github.com/harrys522/opensea-streamer/entity"
	"github.com/harrys522/opensea-streamer/opensea"
	"github.com/harrys522/opensea-streamer/types"
	"github.com/mitchellh/mapstructure"
	"github.com/xiaowang7777/phx"
)

func main() {
	client := opensea.NewStreamClient(types.MAINNET, "apikey", phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err == nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnItemListed("collection-slug", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})

	client.OnItemSold("collection-slug", func(response any) {
		var itemSoldEvent entity.ItemSoldEvent
		err := mapstructure.Decode(response, &itemSoldEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemSoldEvent)
	})

	client.OnItemTransferred("collection-slug", func(response any) {
		var itemTransferredEvent entity.ItemTransferredEvent
		err := mapstructure.Decode(response, &itemTransferredEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemTransferredEvent)
	})

	client.OnItemCancelled("collection-slug", func(response any) {
		var itemCancelledEvent entity.ItemCancelledEvent
		err := mapstructure.Decode(response, &itemCancelledEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemCancelledEvent)
	})

	client.OnItemReceivedBid("collection-slug", func(response any) {
		var itemReceivedBidEvent entity.ItemReceivedBidEvent
		err := mapstructure.Decode(response, &itemReceivedBidEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemReceivedBidEvent)
	})
	client.OnItemReceivedOffer("collection-slug", func(response any) {
		var itemReceivedOfferEvent entity.ItemReceivedOfferEvent
		err := mapstructure.Decode(response, &itemReceivedOfferEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemReceivedOfferEvent)
	})

	client.OnItemMetadataUpdated("collection-slug", func(response any) {
		var itemMetadataUpdateEvent entity.ItemMetadataUpdateEvent
		err := mapstructure.Decode(response, &itemMetadataUpdateEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemMetadataUpdateEvent)
	})

	select {}
}
