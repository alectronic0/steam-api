package steamclient

type GetRecentlyPlayedGamesAPIResponse struct {
	Response RecentlyPlayedGames `json:"response"`
}

type RecentlyPlayedGames struct {
	TotalCount int                  `json:"total_count"`
	Games      []RecentlyPlayedGame `json:"games"`
}

type RecentlyPlayedGame struct {
	ID                     int    `json:"appid"`
	Name                   string `json:"name"`
	Playtime2Weeks         int    `json:"playtime_2weeks"`
	PlaytimeForever        int    `json:"playtime_forever"`
	ImgIconUrl             string `json:"img_icon_url"`
	PlaytimeWindowsForever int    `json:"playtime_windows_forever"`
	PlaytimeMacForever     int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever   int    `json:"playtime_linux_forever"`
	PlaytimeDeckForever    int    `json:"playtime_deck_forever"`
}
