package steamclient

import (
	"net/url"
	"steam-api/pkg/httpclient"
	"strconv"
	"strings"
)

const (
	BaseApiURL   = "https://api.steampowered.com"
	BaseStoreURL = "https://store.steampowered.com"
)

// IClient - https://developer.valvesoftware.com/wiki/Steam_Web_API
type IClient interface {
	// SteamIDs

	GetPlayerSummaries(steamIDs ...string) (*GetPlayerSummariesAPIResponse, error)
	GetPlayerBans(steamIDs ...string) (*GetPlayerBansAPIResponse, error)

	// SteamID

	GetBadges(steamID string) (*GetBadgesAPIResponse, error)
	GetCommunityBadgeProgress(steamID, badgeID string) (*GetCommunityBadgeProgressAPIResponse, error)
	GetFriendList(steamID string, filter GetFriendListFilter) (*GetFriendListAPIResponse, error)
	GetOwnedGames(steamID string, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error)
	GetRecentlyPlayedGames(steamID string, count *uint64) (*GetRecentlyPlayedGamesAPIResponse, error)
	GetSteamLevel(steamID string) (*GetSteamLevelAPIResponse, error)
	GetUserGroupList(steamID string) (*GetUserGroupListAPIResponse, error)
	GetWishlist(steamID string) (*GetWishlistAPIResponse, error)
	GetWishlistItemCount(steamID string) (*GetWishlistItemCountAPIResponse, error)

	// AppIDs

	GetStoreData(appIDs ...string) (GetStoreDataAPIResponse, error)

	// AppID Only

	GetGlobalAchievementPercentagesForApp(appID string) (*GetGlobalAchievementPercentagesForAppAPIResponse, error)
	GetNumberOfCurrentPlayers(appID string) (*GetNumberOfCurrentPlayersAPIResponse, error)
	GetSchemaForGame(appID string) (*GetSchemaForGameAPIResponse, error)

	// SteamID & AppID

	GetPlayerAchievements(steamID, appID string) (*GetPlayerAchievementsAPIResponse, error)
	GetUserStatsForGame(steamID, appID string) (*GetUserStatsForGameAPIResponse, error)

	// Other API

	GetSupportedAPIList() (*GetSupportedAPIListAPIResponse, error)
	GetAppList() (*GetAppListAPIResponse, error)
	ResolveVanityURL(vanityUrl string, vanityURLType VanityURLType) (*ResolveVanityURLAPIResponse, error)

	// To implement
	// GetAppListWithFilters() (*GetAppListWithFiltersAPIResponse, error) // TODO:
	// GetGlobalStatsForGame() (*GetGlobalStatsForGameAPIResponse, error) // TODO: Probably not worth it
}

type Client struct {
	apiKey string
}

func New(apiKey string) IClient {
	return &Client{
		apiKey: apiKey,
	}
}

func (c Client) GetSupportedAPIList() (*GetSupportedAPIListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamWebAPIUtil/GetSupportedAPIList/v0001/?key=" + url.QueryEscape(c.apiKey)
	return httpclient.NillableGet[GetSupportedAPIListAPIResponse](reqURL)
}

func (c Client) GetAppList() (*GetAppListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamApps/GetAppList/v0002"
	return httpclient.NillableGet[GetAppListAPIResponse](reqURL)
}

func (c Client) GetPlayerSummaries(steamIDs ...string) (*GetPlayerSummariesAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetPlayerSummaries/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamids=" + url.QueryEscape(strings.Join(steamIDs, ",")) +
		"&format=json"

	return httpclient.NillableGet[GetPlayerSummariesAPIResponse](reqURL)
}

func (c Client) GetFriendList(steamID string, filter GetFriendListFilter) (*GetFriendListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetFriendList/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&relationship=" + url.QueryEscape(filter.String()) +
		"&format=json"

	return httpclient.NillableGet[GetFriendListAPIResponse](reqURL)
}

func (c Client) GetPlayerAchievements(steamID, appID string) (*GetPlayerAchievementsAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetPlayerAchievements/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.NillableGet[GetPlayerAchievementsAPIResponse](reqURL)
}

func (c Client) GetUserStatsForGame(steamID, appID string) (*GetUserStatsForGameAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetUserStatsForGame/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.NillableGet[GetUserStatsForGameAPIResponse](reqURL)
}

func (c Client) GetOwnedGames(steamID string, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetOwnedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&include_appinfo=" + strconv.FormatBool(includeAppInfo) +
		"&include_played_free_games=" + strconv.FormatBool(includePlayedFreeGames) +
		"&format=json"

	return httpclient.NillableGet[GetOwnedGamesAPIResponse](reqURL)
}

func (c Client) GetRecentlyPlayedGames(steamID string, count *uint64) (*GetRecentlyPlayedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetRecentlyPlayedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	if count != nil {
		reqURL += "&count=" + strconv.FormatUint(*count, 10)
	}

	return httpclient.NillableGet[GetRecentlyPlayedGamesAPIResponse](reqURL)
}

func (c Client) GetBadges(steamID string) (*GetBadgesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetBadges/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	return httpclient.NillableGet[GetBadgesAPIResponse](reqURL)
}

func (c Client) GetCommunityBadgeProgress(steamID, badgeID string) (*GetCommunityBadgeProgressAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetCommunityBadgeProgress/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&badgeid" + url.QueryEscape(badgeID) +
		"&format=json"

	return httpclient.NillableGet[GetCommunityBadgeProgressAPIResponse](reqURL)
}

func (c Client) GetSteamLevel(steamID string) (*GetSteamLevelAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetSteamLevel/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	return httpclient.NillableGet[GetSteamLevelAPIResponse](reqURL)
}

func (c Client) GetPlayerBans(steamIDs ...string) (*GetPlayerBansAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetPlayerBans/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamids=" + url.QueryEscape(strings.Join(steamIDs, ",")) +
		"&format=json"

	return httpclient.NillableGet[GetPlayerBansAPIResponse](reqURL)
}

func (c Client) GetUserGroupList(steamID string) (*GetUserGroupListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetUserGroupList/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	return httpclient.NillableGet[GetUserGroupListAPIResponse](reqURL)
}

func (c Client) ResolveVanityURL(vanityUrl string, vanityURLType VanityURLType) (*ResolveVanityURLAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/ResolveVanityURL/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&vanityurl=" + url.QueryEscape(vanityUrl) +
		"&url_type=" + url.QueryEscape(vanityURLType.String()) +
		"&format=json"

	return httpclient.NillableGet[ResolveVanityURLAPIResponse](reqURL)
}

func (c Client) GetGlobalAchievementPercentagesForApp(appID string) (*GetGlobalAchievementPercentagesForAppAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.NillableGet[GetGlobalAchievementPercentagesForAppAPIResponse](reqURL)
}

func (c Client) GetNumberOfCurrentPlayers(appID string) (*GetNumberOfCurrentPlayersAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetNumberOfCurrentPlayers/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.NillableGet[GetNumberOfCurrentPlayersAPIResponse](reqURL)
}

func (c Client) GetSchemaForGame(appID string) (*GetSchemaForGameAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetSchemaForGame/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(appID) +
		"&format=json"

	return httpclient.NillableGet[GetSchemaForGameAPIResponse](reqURL)
}

func (c Client) GetWishlist(steamID string) (*GetWishlistAPIResponse, error) {
	reqURL := BaseApiURL + "/IWishlistService/GetWishlist/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	return httpclient.NillableGet[GetWishlistAPIResponse](reqURL)
}

func (c Client) GetWishlistItemCount(steamID string) (*GetWishlistItemCountAPIResponse, error) {
	reqURL := BaseApiURL + "/IWishlistService/GetWishlistItemCount/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(steamID) +
		"&format=json"

	return httpclient.NillableGet[GetWishlistItemCountAPIResponse](reqURL)
}

func (c Client) GetStoreData(appIDs ...string) (GetStoreDataAPIResponse, error) {
	resp := GetStoreDataAPIResponse{}

	count := 0

	for _, appID := range appIDs {
		if count < 5 {
			reqURL := BaseStoreURL + "/api/appdetails?appids=" + url.QueryEscape(appID)
			r, err := httpclient.Get[GetStoreDataAPIResponse](reqURL)
			if err != nil {
				return GetStoreDataAPIResponse{}, err
			}
			resp[appID] = r[appID]
		}
		count++
	}

	println(len(resp))

	return resp, nil
}
