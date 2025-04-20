package steamservice

import "steam-api/internal/steamclient"

type UserStatsForGame struct {
	ID           string           `json:"id"`
	GameName     string           `json:"game_name"`
	Achievements map[string]int64 `json:"achievements"`
	Stats        map[string]int64 `json:"stats"`
}

func UserStatsForGameFromAPI(m *steamclient.GetUserStatsForGameAPIResponse) *UserStatsForGame {
	if m == nil {
		return nil
	}

	playerstats := m.Playerstats
	return &UserStatsForGame{
		ID:           playerstats.ID,
		GameName:     playerstats.GameName,
		Achievements: achievementFromAPI(playerstats.Achievements),
		Stats:        statFromApi(playerstats.Stats),
	}
}

func achievementFromAPI(m []steamclient.Achievement) map[string]int64 {
	achievement := make(map[string]int64)
	for _, a := range m {
		achievement[a.Name] = a.Achieved
	}

	return achievement
}

func statFromApi(m []steamclient.Stat) map[string]int64 {
	achievement := make(map[string]int64)
	for _, s := range m {
		achievement[s.Name] = s.Value
	}

	return achievement
}
