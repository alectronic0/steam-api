package steamservice

import "steam-api/internal/steamclient"

type IService interface {
	GetSupportApiList(apiKey, steamID, appID string) (SupportedAPIList, error)

	GetUserInfo(steamID string) (*UserInfoWithGameInfo, error)
	GetUserInfoWithGameInfo(steamID, appID string) (*UserInfoWithGameInfo, error)
	GetAppList() (map[string]string, error)
	HydrateGames(games []OwnedGame) ([]OwnedGame, error)
}

type Service struct {
	client steamclient.IClient
}

func New(client steamclient.IClient) IService {
	return &Service{
		client: client,
	}
}

func (s Service) GetAppList() (map[string]string, error) {
	appList, err := s.client.GetAppList()
	if err != nil {
		return map[string]string{}, err
	}

	mapOfApps := map[string]string{}
	for _, v := range appList.AppList.Apps {
		mapOfApps[v.AppID] = v.Name
	}

	return mapOfApps, nil
}

func (s Service) GetSupportApiList(apiKey, steamID, appID string) (SupportedAPIList, error) {
	supportedAPIList, err := s.client.GetSupportedAPIList()
	if err != nil {
		return SupportedAPIList{}, err
	}

	return SupportedAPIListFromAPI(supportedAPIList, apiKey, steamID, appID), nil
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

func (s Service) GetUserInfoWithGameInfo(steamID, appID string) (*UserInfoWithGameInfo, error) {
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

func (s Service) HydrateGames(games []OwnedGame) ([]OwnedGame, error) {
	var appIds = make([]string, len(games))
	for i, game := range games {
		appIds[i] = game.ID
	}

	println(appIds)

	storeData, err := s.client.GetStoreData(appIds...)
	if err != nil {
		return games, err
	}

	for _, game := range games {
		if data, ok := storeData[game.ID]; ok {
			game.StoreData = &data.Data
		}
	}

	return games, nil
}
