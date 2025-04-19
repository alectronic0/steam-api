package steam

type GetPlayerSummariesAPIResponse struct {
	Response struct {
		Players []Player `json:"players"`
	} `json:"response"`
}

type Player struct {
	SteamID                  string `json:"steamid"`
	PersonaName              string `json:"personaname"`
	ProfileURL               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	LastLogoff               int64  `json:"lastlogoff"`
	CommentPermission        int    `json:"commentpermission"`
}
