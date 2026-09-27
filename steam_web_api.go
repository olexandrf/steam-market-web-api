package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const timeout = 3 * time.Second
const steamURL = "https://api.steampowered.com"
const steamNews = "/ISteamNews/GetNewsForApp/v0002"
const steamGlobalAchievements = "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002"
const format = "json"
const cs2appid = "730"

type SteamWebAPIClient struct {
	url string
	c   *http.Client
}

func NewSteamWebAPIClient(client *http.Client, url string) *SteamWebAPIClient {
	if client == nil {
		client = &http.Client{Timeout: timeout}
	}

	return &SteamWebAPIClient{c: client, url: url}
}

type NewsItem struct {
	Gid           string   `json:"gid"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
	IsExternalURL bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	FeedLabel     string   `json:"feedlabel"`
	Date          int      `json:"date"`
	FeedName      string   `json:"feedname"`
	FeedType      int      `json:"feed_type"`
	Appid         int      `json:"appid"`
	Tags          []string `json:"tags,omitempty"`
}

type AppNews struct {
	Appid     int        `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
	Count     int        `json:"count"`
}

type SteamNews struct {
	AppNews AppNews `json:"appnews"`
}

type Achievement struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
}

type SteamAchievements struct {
	AchievementPercentages struct {
		Achievements []Achievement `json:"achievements"`
	} `json:"achievementpercentages"`
}

func (s *SteamWebAPIClient) GetNews(ctx context.Context, appid string, count int) (*SteamNews, error) {
	params := url.Values{}
	params.Set("appid", appid)
	params.Set("format", format)
	params.Set("count", strconv.Itoa(count))

	reqURL, err := BuildURL(s.url, steamNews, params)
	if err != nil {
		return nil, fmt.Errorf("get news: %w", err)
	}

	news, err := DoJSON[SteamNews](ctx, s.c, reqURL)
	if err != nil {
		return nil, fmt.Errorf("get news: %w", err)
	}
	return news, nil
}

func (s *SteamWebAPIClient) GetGlobalAchievementsPercentages(ctx context.Context, gameId string) (*SteamAchievements, error) {
	params := url.Values{}
	params.Set("gameid", gameId)
	params.Set("format", format)

	reqURL, err := BuildURL(s.url, steamGlobalAchievements, params)
	if err != nil {
		return nil, fmt.Errorf("get global achievements: %w", err)
	}

	achievements, err := DoJSON[SteamAchievements](ctx, s.c, reqURL)
	if err != nil {
		return nil, fmt.Errorf("get global achievements: %w", err)
	}
	return achievements, nil
}
