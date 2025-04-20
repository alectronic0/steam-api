package steamservice

import (
	"steam-api/internal/steamclient"
	"time"
)

type OwnedGames struct {
	GameCount int64               `json:"game_count"`
	Games     map[int64]OwnedGame `json:"games"`
}

type OwnedGame struct {
	ID                       int64         `json:"id"`
	Name                     string        `json:"name"`
	ImgIconUrl               string        `json:"img_icon_url"`
	HasCommunityVisibleStats bool          `json:"has_community_visible_stats"`
	PlaytimeForever          time.Duration `json:"playtime_forever"`
	PlaytimeWindowsForever   time.Duration `json:"playtime_windows_forever"`
	PlaytimeMacForever       time.Duration `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     time.Duration `json:"playtime_linux_forever"`
	PlaytimeDeckForever      time.Duration `json:"playtime_deck_forever"`
	PlaytimeDisconnected     time.Duration `json:"playtime_disconnected"`
	LastPlayedTime           time.Time     `json:"last_played_time"`
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
			PlaytimeForever:          time.Second * time.Duration(g.PlaytimeForever),
			PlaytimeWindowsForever:   time.Second * time.Duration(g.PlaytimeWindowsForever),
			PlaytimeMacForever:       time.Second * time.Duration(g.PlaytimeMacForever),
			PlaytimeLinuxForever:     time.Second * time.Duration(g.PlaytimeLinuxForever),
			PlaytimeDeckForever:      time.Second * time.Duration(g.PlaytimeDeckForever),
			LastPlayedTime:           time.Unix(g.RtimeLastPlayed, 0),
			PlaytimeDisconnected:     0,
		}
	}

	return ownedGames
}
