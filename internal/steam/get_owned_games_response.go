package steam

type GetOwnedGamesAPIResponse struct {
	Response struct {
		GameCount int         `json:"game_count"`
		Games     []OwnedGame `json:"games"`
	} `json:"response"`
}

type GetOwnedGamesResponse struct {
	GameCount int         `json:"game_count"`
	Games     []OwnedGame `json:"games"`
}

//	{
//	   "appid": 4000,
//	   "playtime_forever": 3,
//	   "playtime_windows_forever": 0,
//	   "playtime_mac_forever": 0,
//	   "playtime_linux_forever": 0,
//	   "playtime_deck_forever": 0,
//	    "rtime_last_played": 1452718395,
//	    "playtime_disconnected": 0
//	}
type OwnedGame struct {
	AppID                    int    `json:"appid"`
	Name                     string `json:"name,omitempty"`
	PlaytimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
}
