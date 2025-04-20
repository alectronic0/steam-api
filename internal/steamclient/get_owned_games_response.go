package steamclient

type GetOwnedGamesAPIResponse struct {
	Response OwnedGames `json:"response"`
}

type OwnedGames struct {
	GameCount int64       `json:"game_count"`
	Games     []OwnedGame `json:"games"`
}

type OwnedGame struct {
	ID                       int64  `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int64  `json:"playtime_forever"`
	ImgIconUrl               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeWindowsForever   int64  `json:"playtime_windows_forever"`
	PlaytimeMacForever       int64  `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int64  `json:"playtime_linux_forever"`
	PlaytimeDeckForever      int64  `json:"playtime_deck_forever"`
	RtimeLastPlayed          int64  `json:"rtime_last_played"`
	PlaytimeDisconnected     int64  `json:"playtime_disconnected"`
}
