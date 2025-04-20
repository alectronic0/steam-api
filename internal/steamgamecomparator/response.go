package steamgamecomparator

import "steam-api/internal/steamservice"

type Response struct {
	User1 *steamservice.UserInfoWithGameInfo `json:"user1"`
	User2 *steamservice.UserInfoWithGameInfo `json:"user2"`

	SharedGames []steamservice.OwnedGame `json:"shared_games"`
}
