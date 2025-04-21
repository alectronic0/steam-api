package steamservice

import (
	"fmt"
	"steam-api/internal/steamclient"
	"steam-api/pkg/utils"
	"strconv"
	"time"
)

type OwnedGames struct {
	GameCount uint64               `json:"game_count"`
	Games     map[string]OwnedGame `json:"games"`
}

type OwnedGame struct {
	ID                       string         `json:"id"`
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

	// Optional Fields
	StoreData *steamclient.StoreData `json:"store_data,omitempty"`
}

func OwnedGamesFromAPI(m *steamclient.GetOwnedGamesAPIResponse) OwnedGames {
	if m == nil {
		return OwnedGames{
			GameCount: 0,
			Games:     map[string]OwnedGame{},
		}
	}

	response := m.Response
	return OwnedGames{
		GameCount: response.GameCount,
		Games:     ownedGameFromAPI(response.Games),
	}
}

const SteamImgAssetHeader = "http://media.steampowered.com/steamcommunity/public/images/apps/%s/%s.jpg"

func ownedGameFromAPI(m []steamclient.OwnedGame) map[string]OwnedGame {
	ownedGames := map[string]OwnedGame{}
	for _, g := range m {
		stringID := strconv.FormatUint(g.ID, 10)
		ownedGames[stringID] = OwnedGame{
			ID:                       stringID,
			Name:                     g.Name,
			ImgIconUrl:               fmt.Sprintf(SteamImgAssetHeader, stringID, g.ImgIconUrl),
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
