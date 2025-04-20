package steamclient

type GetPlayerAchievementsAPIResponse struct {
	PlayerStats PlayerAchievements `json:"playerstats"`
}

type PlayerAchievements struct {
	ID           string              `json:"steamID"`
	GameName     string              `json:"gameName"`
	Achievements []PlayerAchievement `json:"achievements"`
	Success      bool                `json:"success"`
}

type PlayerAchievement struct {
	APIName    string `json:"apiname"`
	Achieved   int64  `json:"achieved"`
	UnlockTime int64  `json:"unlocktime"`
}
