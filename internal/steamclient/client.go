package steamclient

import (
	"net/url"
	"steam-api/pkg/httpclient"
	"steam-api/pkg/utils"
	"strconv"
	"strings"
)

const (
	BaseApiURL = "https://api.steampowered.com"
)

// IClient - https://developer.valvesoftware.com/wiki/Steam_Web_API
type IClient interface {
	GetSupportedAPIList() (*GetSupportedAPIListAPIResponse, error)
	GetAppList() (*GetAppListAPIResponse, error)

	GetPlayerSummaries(steamIDs ...uint64) (*GetPlayerSummariesAPIResponse, error)
	GetFriendList(steamID uint64, filter GetFriendListFilter) (*GetFriendListAPIResponse, error)
	GetPlayerAchievements(steamID, appID uint64) (*GetPlayerAchievementsAPIResponse, error)
	GetUserStatsForGame(steamID, appID uint64) (*GetUserStatsForGameAPIResponse, error)
	GetOwnedGames(steamID uint64, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error)
	GetRecentlyPlayedGames(steamID uint64, count *uint64) (*GetRecentlyPlayedGamesAPIResponse, error)

	// To implement
	GetBadges(steamID uint64) (*GetBadgesAPIResponse, error)
	GetCommunityBadgeProgress(steamID uint64, badgeID uint64) (*GetCommunityBadgeProgressAPIResponse, error)
	GetSteamLevel(steamID uint64) (*GetSteamLevelAPIResponse, error)
	GetPlayerBans(steamIDs ...uint64) (*GetPlayerBansAPIResponse, error)
	GetUserGroupList(steamID uint64) (*GetUserGroupListAPIResponse, error)
	ResolveVanityURL(vanityUrl string, vanityURLType VanityURLType) (*ResolveVanityURLAPIResponse, error)
	GetGlobalAchievementPercentagesForApp(appId uint64) (*GetGlobalAchievementPercentagesForAppAPIResponse, error)
	// GetGlobalStatsForGame() (*GetGlobalStatsForGameAPIResponse, error) // TODO: Probably not worth it
	GetNumberOfCurrentPlayers(appId uint64) (*GetNumberOfCurrentPlayersAPIResponse, error)
	GetSchemaForGame(appId uint64) (*GetSchemaForGameAPIResponse, error)
	// GetAppListWithFilters() (*GetAppListWithFiltersAPIResponse, error) // TODO:
	GetWishlist(steamID uint64) (*GetWishlistAPIResponse, error)
	GetWishlistItemCount(steamID uint64) (*GetWishlistItemCountAPIResponse, error)
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
	return httpclient.Get[GetSupportedAPIListAPIResponse](reqURL)
}

func (c Client) GetAppList() (*GetAppListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamApps/GetAppList/v0002"
	return httpclient.Get[GetAppListAPIResponse](reqURL)
}

func (c Client) GetPlayerSummaries(steamIDs ...uint64) (*GetPlayerSummariesAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetPlayerSummaries/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamids=" + url.QueryEscape(strings.Join(utils.ListUint64ToString(steamIDs...), ",")) +
		"&format=json"

	return httpclient.Get[GetPlayerSummariesAPIResponse](reqURL)
}

func (c Client) GetFriendList(steamID uint64, filter GetFriendListFilter) (*GetFriendListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetFriendList/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&relationship=" + url.QueryEscape(filter.String()) +
		"&format=json"

	return httpclient.Get[GetFriendListAPIResponse](reqURL)
}

func (c Client) GetPlayerAchievements(steamID uint64, appid uint64) (*GetPlayerAchievementsAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetPlayerAchievements/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&appid=" + url.QueryEscape(strconv.FormatUint(appid, 10)) +
		"&format=json"

	return httpclient.Get[GetPlayerAchievementsAPIResponse](reqURL)
}

func (c Client) GetUserStatsForGame(steamID, appID uint64) (*GetUserStatsForGameAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetUserStatsForGame/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&appid=" + url.QueryEscape(strconv.FormatUint(appID, 10)) +
		"&format=json"

	return httpclient.Get[GetUserStatsForGameAPIResponse](reqURL)
}

func (c Client) GetOwnedGames(steamID uint64, includeAppInfo bool, includePlayedFreeGames bool) (*GetOwnedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetOwnedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&include_appinfo=" + strconv.FormatBool(includeAppInfo) +
		"&include_played_free_games=" + strconv.FormatBool(includePlayedFreeGames) +
		"&format=json"

	return httpclient.Get[GetOwnedGamesAPIResponse](reqURL)
}

func (c Client) GetRecentlyPlayedGames(steamID uint64, count *uint64) (*GetRecentlyPlayedGamesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetRecentlyPlayedGames/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	if count != nil {
		reqURL += "&count=" + strconv.FormatUint(*count, 10)
	}

	return httpclient.Get[GetRecentlyPlayedGamesAPIResponse](reqURL)
}

func (c Client) GetBadges(steamID uint64) (*GetBadgesAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetBadges/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	return httpclient.Get[GetBadgesAPIResponse](reqURL)
}

func (c Client) GetCommunityBadgeProgress(steamID uint64, badgeID uint64) (*GetCommunityBadgeProgressAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetCommunityBadgeProgress/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&badgeid" + url.QueryEscape(strconv.FormatUint(badgeID, 10)) +
		"&format=json"

	return httpclient.Get[GetCommunityBadgeProgressAPIResponse](reqURL)
}

func (c Client) GetSteamLevel(steamID uint64) (*GetSteamLevelAPIResponse, error) {
	reqURL := BaseApiURL + "/IPlayerService/GetSteamLevel/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	return httpclient.Get[GetSteamLevelAPIResponse](reqURL)
}

func (c Client) GetPlayerBans(steamIDs ...uint64) (*GetPlayerBansAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetPlayerBans/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamids=" + url.QueryEscape(strings.Join(utils.ListUint64ToString(steamIDs...), ",")) +
		"&format=json"

	return httpclient.Get[GetPlayerBansAPIResponse](reqURL)
}

func (c Client) GetUserGroupList(steamID uint64) (*GetUserGroupListAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/GetUserGroupList/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	return httpclient.Get[GetUserGroupListAPIResponse](reqURL)
}

func (c Client) ResolveVanityURL(vanityUrl string, vanityURLType VanityURLType) (*ResolveVanityURLAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUser/ResolveVanityURL/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&vanityurl=" + url.QueryEscape(vanityUrl) +
		"&url_type=" + url.QueryEscape(vanityURLType.String()) +
		"&format=json"

	return httpclient.Get[ResolveVanityURLAPIResponse](reqURL)
}

func (c Client) GetGlobalAchievementPercentagesForApp(appID uint64) (*GetGlobalAchievementPercentagesForAppAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(strconv.FormatUint(appID, 10)) +
		"&format=json"

	return httpclient.Get[GetGlobalAchievementPercentagesForAppAPIResponse](reqURL)
}

func (c Client) GetNumberOfCurrentPlayers(appID uint64) (*GetNumberOfCurrentPlayersAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetNumberOfCurrentPlayers/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(strconv.FormatUint(appID, 10)) +
		"&format=json"

	return httpclient.Get[GetNumberOfCurrentPlayersAPIResponse](reqURL)
}

func (c Client) GetSchemaForGame(appID uint64) (*GetSchemaForGameAPIResponse, error) {
	reqURL := BaseApiURL + "/ISteamUserStats/GetSchemaForGame/v0002/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&appid=" + url.QueryEscape(strconv.FormatUint(appID, 10)) +
		"&format=json"

	return httpclient.Get[GetSchemaForGameAPIResponse](reqURL)
}

func (c Client) GetWishlist(steamID uint64) (*GetWishlistAPIResponse, error) {
	reqURL := BaseApiURL + "/IWishlistService/GetWishlist/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	return httpclient.Get[GetWishlistAPIResponse](reqURL)
}

func (c Client) GetWishlistItemCount(steamID uint64) (*GetWishlistItemCountAPIResponse, error) {
	reqURL := BaseApiURL + "/IWishlistService/GetWishlistItemCount/v0001/" +
		"?key=" + url.QueryEscape(c.apiKey) +
		"&steamid=" + url.QueryEscape(strconv.FormatUint(steamID, 10)) +
		"&format=json"

	return httpclient.Get[GetWishlistItemCountAPIResponse](reqURL)
}
