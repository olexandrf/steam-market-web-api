package main

import (
	"context"
	"fmt"
	"log"
)

func logOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	ctx := context.Background()
	steamClient := NewSteamWebAPIClient(nil, steamURL)
	steamNews, err := steamClient.GetNews(ctx, cs2appid, 1)
	logOnError(err, "GetNews")
	fmt.Printf("%+v\n", steamNews)
	SteamMarketClient := NewSteamMarketClient(nil, steamMarketURL)
	steamMarketResponse, err := SteamMarketClient.GetMarketList(ctx, cs2appid)
	logOnError(err, "SearchMarket")
	fmt.Printf("%+v\n", steamMarketResponse)
	steamPriceOverview, err := SteamMarketClient.GetPrices(ctx, cs2appid, "1", "AK-47 | Redline (Field-Tested)")
	logOnError(err, "GetPrices")
	fmt.Printf("%+v\n", steamPriceOverview)
}
