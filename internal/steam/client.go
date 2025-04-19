package steam

import (
	"net/url"
	"steam-api/pkg/httpclient"
	"strconv"
	"strings"
)

const (
	BaseApiURL = "https://api.steampowered.com"
)

// IClient - https://developer.valvesoftware.com/wiki/Steam_Web_API
type IClient interface {
	GetPlayerSummaries(steamID ...string) (*GetPlayerSummariesAPIResponse, error)
	GetFriendList(steamID string, filter GetFriendListFilter) (*GetFriendListAPIResponse, error)
	GetPlayerAchievements(steamID string, appID string) (*GetPlayerAchievementsAPIResponse, error)
	GetUserStatsForGame(steamID, appID string) (*GetUserStatsForGameAPIResponse, error)
	GetOwnedGames(steamID string, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error)
	GetRecentlyPlayedGames(steamID string, count *int64) (*GetRecentlyPlayedGamesAPIResponse, error)
}

type Client struct {
	apiKey string
}

func NewClient(apiKey string) IClient {
	return &Client{
		apiKey: apiKey,
	}
}

func (c Client) GetPlayerSummaries(steamIDs ...string) (*GetPlayerSummariesAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetPlayerSummaries/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamids=" + url.QueryEscape(strings.Join(steamIDs, ",")) +
		"&format=json"

	return httpclient.Get[GetPlayerSummariesAPIResponse](reqURL)
}

func (c Client) GetFriendList(steamID string, filter GetFriendListFilter) (*GetFriendListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetFriendList/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&relationship=" + url.QueryEscape(filter.String()) +
		"&format=json"

	return httpclient.Get[GetFriendListAPIResponse](reqURL)
}

func (c Client) GetPlayerAchievements(steamID string, appid string) (*GetPlayerAchievementsAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetPlayerAchievements/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&appid=" + url.QueryEscape(appid) +
		"&format=json"

	return httpclient.Get[GetPlayerAchievementsAPIResponse](reqURL)
}

func (c Client) GetUserStatsForGame(steamID, appID string) (*GetUserStatsForGameAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetUserStatsForGame/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.Get[GetUserStatsForGameAPIResponse](reqURL)
}

func (c Client) GetOwnedGames(steamID string, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetOwnedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&include_appinfo=" + strconv.FormatBool(includeAppInfo) +
		"&include_played_free_games=" + strconv.FormatBool(includePlayedFreeGames) +
		"&format=json"

	return httpclient.Get[GetOwnedGamesAPIResponse](reqURL)
}

func (c Client) GetRecentlyPlayedGames(steamID string, count *int64) (*GetRecentlyPlayedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetRecentlyPlayedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	if count != nil {
		reqURL += "&count=" + strconv.FormatInt(*count, 10)
	}

	return httpclient.Get[GetRecentlyPlayedGamesAPIResponse](reqURL)
}
