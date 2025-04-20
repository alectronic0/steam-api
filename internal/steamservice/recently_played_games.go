package steamservice

import (
	"steam-api/internal/steamclient"
	"time"
)

type RecentlyPlayedGames struct {
	TotalCount int                  `json:"total_count"`
	Games      []RecentlyPlayedGame `json:"games"`
}

type RecentlyPlayedGame struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	ImgIconUrl             string        `json:"img_icon_url"`
	Playtime2Weeks         time.Duration `json:"playtime_2_weeks"`
	PlaytimeForever        time.Duration `json:"playtime_forever"`
	PlaytimeWindowsForever time.Duration `json:"playtime_windows_forever"`
	PlaytimeMacForever     time.Duration `json:"playtime_mac_forever"`
	PlaytimeLinuxForever   time.Duration `json:"playtime_linux_forever"`
	PlaytimeDeckForever    time.Duration `json:"playtime_deck_forever"`
}

func RecentlyPlayedGamesFromAPI(m *steamclient.GetRecentlyPlayedGamesAPIResponse) RecentlyPlayedGames {
	if m == nil {
		return RecentlyPlayedGames{
			TotalCount: 0,
			Games:      []RecentlyPlayedGame{},
		}
	}

	response := m.Response
	return RecentlyPlayedGames{
		TotalCount: response.TotalCount,
		Games:      recentlyPlayedGameFromAPI(response.Games),
	}
}

func recentlyPlayedGameFromAPI(m []steamclient.RecentlyPlayedGame) []RecentlyPlayedGame {
	recentlyPlayedGames := make([]RecentlyPlayedGame, len(m))
	for i, v := range m {
		recentlyPlayedGames[i] = RecentlyPlayedGame{
			ID:                     v.ID,
			Name:                   v.Name,
			ImgIconUrl:             v.ImgIconUrl,
			Playtime2Weeks:         time.Second * time.Duration(v.Playtime2Weeks),
			PlaytimeForever:        time.Second * time.Duration(v.PlaytimeForever),
			PlaytimeWindowsForever: time.Second * time.Duration(v.PlaytimeWindowsForever),
			PlaytimeMacForever:     time.Second * time.Duration(v.PlaytimeMacForever),
			PlaytimeLinuxForever:   time.Second * time.Duration(v.PlaytimeLinuxForever),
			PlaytimeDeckForever:    time.Second * time.Duration(v.PlaytimeDeckForever),
		}
	}

	return recentlyPlayedGames
}
