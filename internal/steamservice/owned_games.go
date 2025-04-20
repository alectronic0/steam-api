package steamservice

import (
	"steam-api/internal/steamclient"
	"steam-api/pkg/utils"
	"time"
)

type OwnedGames struct {
	GameCount int64               `json:"game_count"`
	Games     map[int64]OwnedGame `json:"games"`
}

type OwnedGame struct {
	ID                       int64          `json:"id"`
	Name                     string         `json:"name"`
	ImgIconUrl               string         `json:"img_icon_url"`
	HasCommunityVisibleStats bool           `json:"has_community_visible_stats"`
	PlaytimeForever          utils.Duration `json:"playtime_forever"`
	PlaytimeWindowsForever   utils.Duration `json:"playtime_windows_forever"`
	PlaytimeMacForever       utils.Duration `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     utils.Duration `json:"playtime_linux_forever"`
	PlaytimeDeckForever      utils.Duration `json:"playtime_deck_forever"`
	PlaytimeDisconnected     utils.Duration `json:"playtime_disconnected"`
	LastPlayedTime           time.Time      `json:"last_played_time"`
}

func OwnedGamesFromAPI(m *steamclient.GetOwnedGamesAPIResponse) OwnedGames {
	if m == nil {
		return OwnedGames{
			GameCount: int64(0),
			Games:     map[int64]OwnedGame{},
		}
	}

	response := m.Response
	return OwnedGames{
		GameCount: response.GameCount,
		Games:     ownedGameFromAPI(response.Games),
	}
}

func ownedGameFromAPI(m []steamclient.OwnedGame) map[int64]OwnedGame {
	ownedGames := map[int64]OwnedGame{}
	for _, g := range m {
		ownedGames[g.ID] = OwnedGame{
			ID:                       g.ID,
			Name:                     g.Name,
			ImgIconUrl:               g.ImgIconUrl,
			HasCommunityVisibleStats: g.HasCommunityVisibleStats,
			PlaytimeForever:          utils.NewDuration(time.Minute * time.Duration(g.PlaytimeForever)),
			PlaytimeWindowsForever:   utils.NewDuration(time.Minute * time.Duration(g.PlaytimeWindowsForever)),
			PlaytimeMacForever:       utils.NewDuration(time.Minute * time.Duration(g.PlaytimeMacForever)),
			PlaytimeLinuxForever:     utils.NewDuration(time.Minute * time.Duration(g.PlaytimeLinuxForever)),
			PlaytimeDeckForever:      utils.NewDuration(time.Minute * time.Duration(g.PlaytimeDeckForever)),
			LastPlayedTime:           time.Unix(g.RtimeLastPlayed, 0),
			PlaytimeDisconnected:     utils.NewDuration(time.Minute * time.Duration(g.PlaytimeDisconnected)),
		}
	}

	return ownedGames
}
