package steamservice

import (
	"steam-api/internal/steamclient"
	"time"
)

type PlayerSummaries map[uint64]PlayerSummary

type PlayerSummary struct {
	ID                       uint64    `json:"id"`
	CommunityVisibilityState int64     `json:"community_visibility_state"`
	ProfileState             int64     `json:"profile_state"`
	PersonaName              string    `json:"persona_name"`
	ProfileUrl               string    `json:"profile_url"`
	Avatar                   string    `json:"avatar"`
	AvatarMedium             string    `json:"avatar_medium"`
	AvatarFull               string    `json:"avatar_full"`
	AvatarHash               string    `json:"avatar_hash"`
	LastLogoff               time.Time `json:"last_logoff"`
	PersonaState             int64     `json:"persona_state"`
	RealName                 string    `json:"real_name"`
	PrimaryClanID            string    `json:"primary_clan_id"`
	TimeCreated              time.Time `json:"time_created"`
	PersonaStateFlags        int64     `json:"persona_state_flags"`
	CountryCode              string    `json:"country_code"`
}

func PlayerSummariesFromAPI(m *steamclient.GetPlayerSummariesAPIResponse) PlayerSummaries {
	if m == nil {
		return nil
	}

	playerSummaries := PlayerSummaries{}
	for _, v := range m.Response.Players {
		playerSummaries[v.ID] = PlayerSummary{
			ID:                       v.ID,
			CommunityVisibilityState: v.Communityvisibilitystate,
			ProfileState:             v.Profilestate,
			PersonaName:              v.Personaname,
			ProfileUrl:               v.Profileurl,
			Avatar:                   v.Avatar,
			AvatarMedium:             v.Avatarmedium,
			AvatarFull:               v.Avatarfull,
			AvatarHash:               v.Avatarhash,
			LastLogoff:               time.Unix(v.Lastlogoff, 0),
			PersonaState:             v.Personastate,
			RealName:                 v.Realname,
			PrimaryClanID:            v.Primaryclanid,
			TimeCreated:              time.Unix(v.Timecreated, 0),
			PersonaStateFlags:        v.Personastateflags,
			CountryCode:              v.Loccountrycode,
		}
	}

	return playerSummaries
}
