package steamclient

import "strconv"

type VanityURLType uint64

const (
	DefaultVanityURLType           = 1
	GroupVanityURLType             = 2
	OfficialGameGroupVanityURLType = 2
)

func (v VanityURLType) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

//type GetGlobalStatsForGameAPIResponse struct {} // TODO: Probably not worth it
//type GetAppListWithFiltersAPIResponse struct{} // TODO:

////////////////////////////////////////////////////////////////////
// GetBadgesAPIResponse
////////////////////////////////////////////////////////////////////

type GetBadgesAPIResponse struct {
	Response BadgeResponse `json:"response"`
}
type BadgeResponse struct {
	Badges                     []Badge `json:"badges"`
	PlayerXp                   uint64  `json:"player_xp"`
	PlayerLevel                uint64  `json:"player_level"`
	PlayerXpNeededToLevelUp    uint64  `json:"player_xp_needed_to_level_up"`
	PlayerXpNeededCurrentLevel uint64  `json:"player_xp_needed_current_level"`
}

type Badge struct {
	ID             int64 `json:"badgeid"`
	Level          int64 `json:"level"`
	CompletionTime int64 `json:"completion_time"`
	Xp             int64 `json:"xp"`
	Scarcity       int64 `json:"scarcity"`
}

////////////////////////////////////////////////////////////////////
// GetCommunityBadgeProgressAPIResponse
////////////////////////////////////////////////////////////////////

type GetCommunityBadgeProgressAPIResponse struct {
	Response CommunityBadges `json:"response"`
}
type CommunityBadges struct {
	Quests []Quest `json:"quests"`
}
type Quest struct {
	ID        int64 `json:"questid"`
	Completed bool  `json:"completed"`
}

////////////////////////////////////////////////////////////////////
// GetSteamLevelAPIResponse
////////////////////////////////////////////////////////////////////

type GetSteamLevelAPIResponse struct {
	Response SteamLevel `json:"response"`
}

type SteamLevel struct {
	PlayerLevel int `json:"player_level"`
}

////////////////////////////////////////////////////////////////////
// GetPlayerBansAPIResponse
////////////////////////////////////////////////////////////////////

type GetPlayerBansAPIResponse struct {
	Players []PlayerBan `json:"players"`
}

type PlayerBan struct {
	ID               string `json:"SteamId"`
	CommunityBanned  bool   `json:"CommunityBanned"`
	VACBanned        bool   `json:"VACBanned"`
	NumberOfVACBans  int64  `json:"NumberOfVACBans"`
	DaysSinceLastBan int64  `json:"DaysSinceLastBan"`
	NumberOfGameBans int64  `json:"NumberOfGameBans"`
	EconomyBan       string `json:"EconomyBan"`
}

////////////////////////////////////////////////////////////////////
// GetUserGroupListAPIResponse
////////////////////////////////////////////////////////////////////

type GetUserGroupListAPIResponse struct {
	Response UserGroupListResponse `json:"response"`
}

type UserGroupListResponse struct {
	Success bool    `json:"success"`
	Groups  []Group `json:"groups"`
}

type Group struct {
	ID string `json:"gid"`
}

////////////////////////////////////////////////////////////////////
// ResolveVanityURLAPIResponse
////////////////////////////////////////////////////////////////////

type ResolveVanityURLAPIResponse struct {
	Response ResolveVanityURL `json:"response"`
}

type ResolveVanityURL struct {
	ID      string `json:"steamid"`
	Success int    `json:"success"`
	Message string `json:"message"`
}

////////////////////////////////////////////////////////////////////
// GetGlobalAchievementPercentagesForAppAPIResponse
////////////////////////////////////////////////////////////////////

type GetGlobalAchievementPercentagesForAppAPIResponse struct {
	Achievementpercentages Achievementpercentages `json:"achievementpercentages"`
}

type Achievementpercentages struct {
	Achievements []GlobalAchievementPercentage `json:"achievements"`
}

type GlobalAchievementPercentage struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
}

////////////////////////////////////////////////////////////////////
// GetNumberOfCurrentPlayersAPIResponse
////////////////////////////////////////////////////////////////////

type GetNumberOfCurrentPlayersAPIResponse struct {
	Response NumberOfCurrentPlayers `json:"response"`
}

type NumberOfCurrentPlayers struct {
	PlayerCount int `json:"player_count"`
	Result      int `json:"result"`
}

////////////////////////////////////////////////////////////////////
// GetSchemaForGameAPIResponse
////////////////////////////////////////////////////////////////////

type GetSchemaForGameAPIResponse struct {
	Game SchemaForGame `json:"game"`
}

type SchemaForGame struct {
	GameName           string             `json:"gameName"`
	GameVersion        string             `json:"gameVersion"`
	AvailableGameStats AvailableGameStats `json:"availableGameStats"`
}

type AvailableGameStats struct {
	Achievements []GameStatsAchievements `json:"achievements"`
	Stats        []GameStats             `json:"stats"`
}
type GameStatsAchievements struct {
	Name         string `json:"name"`
	Defaultvalue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	Icongray     string `json:"icongray"`
}

type GameStats struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
}

////////////////////////////////////////////////////////////////////
// GetWishlistAPIResponse
////////////////////////////////////////////////////////////////////

type GetWishlistAPIResponse struct {
	Response Wishlist `json:"response"`
}

type Wishlist struct {
	Items []WishListItem `json:"items"`
}

type WishListItem struct {
	ID        int `json:"appid"`
	Priority  int `json:"priority"`
	DateAdded int `json:"date_added"`
}

////////////////////////////////////////////////////////////////////
// GetWishlistItemCountAPIResponse
////////////////////////////////////////////////////////////////////

type GetWishlistItemCountAPIResponse struct {
	Response WishlistItemCount `json:"response"`
}

type WishlistItemCount struct {
	Count uint64 `json:"count"`
}
