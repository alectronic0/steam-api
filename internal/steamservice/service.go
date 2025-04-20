package steamservice

import "steam-api/internal/steamclient"

type IService interface {
	GetUserInfo(steamID string) (*UserInfoWithGameInfo, error)
	GetUserInfoWithGameInfo(steamID string, appID string) (*UserInfoWithGameInfo, error)
}

type Service struct {
	client steamclient.IClient
}

func New(client steamclient.IClient) IService {
	return &Service{
		client: client,
	}
}

func (s Service) GetUserInfo(steamID string) (*UserInfoWithGameInfo, error) {
	var (
		err                         error
		playerSummariesResponse     *steamclient.GetPlayerSummariesAPIResponse
		friendListResponse          *steamclient.GetFriendListAPIResponse
		ownedGamesResponse          *steamclient.GetOwnedGamesAPIResponse
		recentlyPlayedGamesResponse *steamclient.GetRecentlyPlayedGamesAPIResponse
	)

	playerSummariesResponse, err = s.client.GetPlayerSummaries(steamID)
	if err != nil {
		return nil, err
	}

	friendListResponse, err = s.client.GetFriendList(steamID, steamclient.GetFriendListFilterAll)
	if err != nil {
		return nil, err
	}

	ownedGamesResponse, err = s.client.GetOwnedGames(steamID, true, true)
	if err != nil {
		return nil, err
	}

	recentlyPlayedGamesResponse, err = s.client.GetRecentlyPlayedGames(steamID, nil)
	if err != nil {
		return nil, err
	}

	return &UserInfoWithGameInfo{
		UserInfo: UserInfo{
			PlayerSummary:       PlayerSummariesFromAPI(playerSummariesResponse)[steamID],
			FriendsList:         FriendsListFromAPI(friendListResponse),
			OwnedGames:          OwnedGamesFromAPI(ownedGamesResponse),
			RecentlyPlayedGames: RecentlyPlayedGamesFromAPI(recentlyPlayedGamesResponse),
		},
	}, nil
}

func (s Service) GetUserInfoWithGameInfo(steamID string, appID string) (*UserInfoWithGameInfo, error) {
	var (
		err                         error
		playerSummariesResponse     *steamclient.GetPlayerSummariesAPIResponse
		friendListResponse          *steamclient.GetFriendListAPIResponse
		ownedGamesResponse          *steamclient.GetOwnedGamesAPIResponse
		recentlyPlayedGamesResponse *steamclient.GetRecentlyPlayedGamesAPIResponse
		playerAchievementsResponse  *steamclient.GetPlayerAchievementsAPIResponse
		userStatsForGameResponse    *steamclient.GetUserStatsForGameAPIResponse
	)

	playerSummariesResponse, err = s.client.GetPlayerSummaries(steamID)
	if err != nil {
		return nil, err
	}

	friendListResponse, err = s.client.GetFriendList(steamID, steamclient.GetFriendListFilterAll)
	if err != nil {
		return nil, err
	}

	ownedGamesResponse, err = s.client.GetOwnedGames(steamID, true, true)
	if err != nil {
		return nil, err
	}

	recentlyPlayedGamesResponse, err = s.client.GetRecentlyPlayedGames(steamID, nil)
	if err != nil {
		return nil, err
	}

	playerAchievementsResponse, err = s.client.GetPlayerAchievements(steamID, appID)
	if err != nil {
		return nil, err
	}

	userStatsForGameResponse, err = s.client.GetUserStatsForGame(steamID, appID)
	if err != nil {
		return nil, err
	}

	return &UserInfoWithGameInfo{
		UserInfo: UserInfo{
			PlayerSummary:       PlayerSummariesFromAPI(playerSummariesResponse)[steamID],
			FriendsList:         FriendsListFromAPI(friendListResponse),
			OwnedGames:          OwnedGamesFromAPI(ownedGamesResponse),
			RecentlyPlayedGames: RecentlyPlayedGamesFromAPI(recentlyPlayedGamesResponse),
		},
		UserGameInfo: &UserGameInfo{
			PlayerAchievements: PlayerAchievementsFromAPI(playerAchievementsResponse),
			UserStatsForGame:   UserStatsForGameFromAPI(userStatsForGameResponse),
		},
	}, nil
}
