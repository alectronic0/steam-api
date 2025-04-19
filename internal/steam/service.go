package steam

type UserInfo struct {
}

type UserInfoWithGameInfo struct {
}

type IService interface {
	GetUserInfo(steamID string) (UserInfo, error)
}

type Service struct {
	client Client
}

func (s Service) GetUserInfo(steamID string) (*UserInfo, error) {
	var (
		err                 error
		playerSummaries     *GetPlayerSummariesAPIResponse
		friendList          *GetFriendListAPIResponse
		ownedGames          *GetOwnedGamesAPIResponse
		recentlyPlayedGames *GetRecentlyPlayedGamesAPIResponse
	)

	playerSummaries, err = s.client.GetPlayerSummaries(steamID)
	if err != nil {
		return nil, err
	}

	friendList, err = s.client.GetFriendList(steamID, GetFriendListFilterAll)
	if err != nil {
		return nil, err
	}

	ownedGames, err = s.client.GetOwnedGames(steamID, true, true)
	if err != nil {
		return nil, err
	}

	recentlyPlayedGames, err = s.client.GetRecentlyPlayedGames(steamID, nil)
	if err != nil {
		return nil, err
	}

	return &UserInfo{}, nil
}

func (s Service) GetFriendList(steamID string, appID string) (*UserInfoWithGameInfo, error) {
	var (
		err                 error
		playerSummaries     *GetPlayerSummariesAPIResponse
		friendList          *GetFriendListAPIResponse
		ownedGames          *GetOwnedGamesAPIResponse
		recentlyPlayedGames *GetRecentlyPlayedGamesAPIResponse
		playerAchievements  *GetPlayerAchievementsAPIResponse
		userStatsForGame    *GetUserStatsForGameAPIResponse
	)

	playerSummaries, err = s.client.GetPlayerSummaries(steamID)
	if err != nil {
		return nil, err
	}

	friendList, err = s.client.GetFriendList(steamID, All)
	if err != nil {
		return nil, err
	}

	ownedGames, err = s.client.GetOwnedGames(steamID, true, true)
	if err != nil {
		return nil, err
	}

	recentlyPlayedGames, err = s.client.GetRecentlyPlayedGames(steamID, nil)
	if err != nil {
		return nil, err
	}

	playerAchievements, err = s.client.GetPlayerAchievements(steamID, appID)
	if err != nil {
		return nil, err
	}

	userStatsForGame, err = s.client.GetUserStatsForGame(steamID, appID)
	if err != nil {
		return nil, err
	}

	return &UserInfoWithGameInfo{}, nil
}
