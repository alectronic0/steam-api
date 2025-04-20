package steamclient

type GetPlayerSummariesAPIResponse struct {
	Response PlayerSummaries `json:"response"`
}

type PlayerSummaries struct {
	Players []Player `json:"players"`
}

type Player struct {
	ID                       string `json:"steamid"`
	Communityvisibilitystate int64  `json:"communityvisibilitystate"`
	Profilestate             int64  `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Lastlogoff               int64  `json:"lastlogoff"`
	Personastate             int64  `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int64  `json:"timecreated"`
	Personastateflags        int64  `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
}
