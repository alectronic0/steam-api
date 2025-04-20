package steamservice

import (
	"steam-api/internal/steamclient"
	"steam-api/pkg/utils"
	"time"
)

type RecentlyPlayedGames struct {
	TotalCount int                  `json:"total_count"`
	Games      []RecentlyPlayedGame `json:"games"`
}

type RecentlyPlayedGame struct {
	ID                     int            `json:"id"`
	Name                   string         `json:"name"`
	ImgIconUrl             string         `json:"img_icon_url"`
	Playtime2Weeks         utils.Duration `json:"playtime_2_weeks"`
	PlaytimeForever        utils.Duration `json:"playtime_forever"`
	PlaytimeWindowsForever utils.Duration `json:"playtime_windows_forever"`
	PlaytimeMacForever     utils.Duration `json:"playtime_mac_forever"`
	PlaytimeLinuxForever   utils.Duration `json:"playtime_linux_forever"`
	PlaytimeDeckForever    utils.Duration `json:"playtime_deck_forever"`
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
			Playtime2Weeks:         utils.NewDuration(time.Minute * time.Duration(v.Playtime2Weeks)),
			PlaytimeForever:        utils.NewDuration(time.Minute * time.Duration(v.PlaytimeForever)),
			PlaytimeWindowsForever: utils.NewDuration(time.Minute * time.Duration(v.PlaytimeWindowsForever)),
			PlaytimeMacForever:     utils.NewDuration(time.Minute * time.Duration(v.PlaytimeMacForever)),
			PlaytimeLinuxForever:   utils.NewDuration(time.Minute * time.Duration(v.PlaytimeLinuxForever)),
			PlaytimeDeckForever:    utils.NewDuration(time.Minute * time.Duration(v.PlaytimeDeckForever)),
		}
	}

	return recentlyPlayedGames
}
