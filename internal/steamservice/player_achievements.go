package steamservice

import (
	"steam-api/internal/steamclient"
	"time"
)

type PlayerAchievements struct {
	ID           string              `json:"id"`
	GameName     string              `json:"game_name"`
	Achievements []PlayerAchievement `json:"achievements"`
	Success      bool                `json:"success"`
}

type PlayerAchievement struct {
	APIName    string    `json:"api_name"`
	Achieved   int64     `json:"achieved"`
	UnlockTime time.Time `json:"unlock_time"`
}

func PlayerAchievementsFromAPI(m *steamclient.GetPlayerAchievementsAPIResponse) *PlayerAchievements {
	if m == nil {
		return nil
	}

	stats := m.PlayerStats
	return &PlayerAchievements{
		ID:           stats.ID,
		GameName:     stats.GameName,
		Achievements: playerAchievementsToAPI(stats.Achievements),
		Success:      stats.Success,
	}
}

func playerAchievementsToAPI(m []steamclient.PlayerAchievement) []PlayerAchievement {
	achievements := make([]PlayerAchievement, len(m))
	for i, v := range m {
		achievements[i] = PlayerAchievement{
			APIName:    v.APIName,
			Achieved:   v.Achieved,
			UnlockTime: time.Unix(v.UnlockTime, 0),
		}
	}

	return achievements
}
