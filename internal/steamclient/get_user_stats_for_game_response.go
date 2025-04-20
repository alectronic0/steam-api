package steamclient

type GetUserStatsForGameAPIResponse struct {
	Playerstats PlayerStats `json:"playerstats"`
}

type PlayerStats struct {
	ID           string        `json:"steamID"`
	GameName     string        `json:"gameName"`
	Achievements []Achievement `json:"achievements"`
	Stats        []Stat        `json:"stats"`
}

type Achievement struct {
	Name     string `json:"name"`
	Achieved int64  `json:"achieved"`
}

type Stat struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}
