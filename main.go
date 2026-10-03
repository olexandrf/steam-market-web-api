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
	steamClient := NewSteamWebAPIClient(nil, SteamURL)
	steamNews, err := steamClient.GetNews(ctx, CS2ID, 1, DefaultFormat)
	logOnError(err, "GetNews")
	fmt.Printf("%+v\n", steamNews)
	steamAchievements, err := steamClient.GetGlobalAchievementsPercentages(ctx, CS2ID, DefaultFormat)
	logOnError(err, "GetGlobalAchievementsPercentages")
	fmt.Printf("%+v\n", steamAchievements)
	SteamMarketClient := NewSteamMarketClient(nil, SteamMarketURL)
	steamMarketResponse, err := SteamMarketClient.GetMarketList(ctx, CS2ID)
	logOnError(err, "SearchMarket")
	fmt.Printf("%+v\n", steamMarketResponse)
	steamPriceOverview, err := SteamMarketClient.GetPrices(ctx, CS2ID, "1", "AK-47 | Redline (Field-Tested)")
	logOnError(err, "GetPrices")
	fmt.Printf("%+v\n", steamPriceOverview)
}
