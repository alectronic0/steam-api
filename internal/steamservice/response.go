package steamservice

type UserInfo struct {
	PlayerSummary       PlayerSummary       `json:"player_summary"`
	FriendsList         FriendsList         `json:"friends_list"`
	OwnedGames          OwnedGames          `json:"owned_games"`
	RecentlyPlayedGames RecentlyPlayedGames `json:"recently_played_games"`
}

type UserInfoWithGameInfo struct {
	UserInfo     UserInfo      `json:"user_info"`
	UserGameInfo *UserGameInfo `json:"user_game_info"`
}

type UserGameInfo struct {
	//GameInfo          GameInfo
	PlayerAchievements *PlayerAchievements `json:"player_achievements"`
	UserStatsForGame   *UserStatsForGame   `json:"user_stats_for_game"`
}
