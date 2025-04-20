package steamservice

import (
	"steam-api/internal/steamclient"
	"time"
)

type FriendsList []Friend

type Friend struct {
	ID           string    `json:"id"`
	Relationship string    `json:"relationship"`
	FriendSince  time.Time `json:"friend_since"`
}

func FriendsListFromAPI(m *steamclient.GetFriendListAPIResponse) FriendsList {
	if m == nil {
		return FriendsList{}
	}

	friends := m.FriendsList.Friends
	friendsList := make(FriendsList, len(friends))
	for i, f := range friends {
		friendsList[i] = Friend{
			ID:           f.ID,
			Relationship: f.Relationship,
			FriendSince:  time.Unix(f.FriendSince, 0),
		}
	}

	return friendsList
}
