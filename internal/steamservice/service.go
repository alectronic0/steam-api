package steamservice

import "steam-api/internal/steamclient"

type IService interface {
	GetSupportApiList(apiKey string, steamID, appID uint64) (SupportedAPIList, error)

	GetUserInfo(steamID uint64) (*UserInfoWithGameInfo, error)
	GetUserInfoWithGameInfo(steamID, appID uint64) (*UserInfoWithGameInfo, error)
	GetAppList() (map[uint64]string, error)
}

type Service struct {
	client steamclient.IClient
}

func (s Service) GetAppList() (map[uint64]string, error) {
	appList, err := s.client.GetAppList()
	if err != nil {
		return map[uint64]string{}, err
	}

	mapOfApps := map[uint64]string{}
	for _, v := range appList.AppList.Apps {
		mapOfApps[v.AppID] = v.Name
	}

	return mapOfApps, nil
}

func (s Service) GetSupportApiList(apiKey string, steamID, appID uint64) (SupportedAPIList, error) {
	supportedAPIList, err := s.client.GetSupportedAPIList()
	if err != nil {
		return SupportedAPIList{}, err
	}

	return SupportedAPIListFromAPI(supportedAPIList, apiKey, steamID, appID), nil
}

func New(client steamclient.IClient) IService {
	return &Service{
		client: client,
	}
}

func (s Service) GetUserInfo(steamID uint64) (*UserInfoWithGameInfo, error) {
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

func (s Service) GetUserInfoWithGameInfo(steamID, appID uint64) (*UserInfoWithGameInfo, error) {
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
