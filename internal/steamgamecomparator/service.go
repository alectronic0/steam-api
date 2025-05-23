package steamgamecomparator

import "steam-api/internal/steamservice"

type IService interface {
	CompareOwnedGames(steamID1, steamIDs string) (*Response, error)
}
type Service struct {
	service steamservice.IService
}

func New(service steamservice.IService) IService {
	return &Service{
		service: service,
	}
}

func (s Service) CompareOwnedGames(steamID1, steamID2 string) (*Response, error) {
	user1, err := s.service.GetUserInfo(steamID1)
	if err != nil {
		return nil, err
	}
	user2, err := s.service.GetUserInfo(steamID2)
	if err != nil {
		return nil, err
	}

	games := compareGames(user1, user2)
	games, err = s.service.HydrateGames(games)
	if err != nil {
		return nil, err
	}

	return &Response{
		User1:       user1,
		User2:       user2,
		SharedGames: games,
	}, nil
}

func compareGames(user1Games, user2Games *steamservice.UserInfoWithGameInfo) []steamservice.OwnedGame {
	games1 := user1Games.UserInfo.OwnedGames.Games
	games2 := user2Games.UserInfo.OwnedGames.Games

	var sharedGames []steamservice.OwnedGame
	for kUg1, vUg1 := range games1 {
		if _, ok := games2[kUg1]; ok {
			sharedGames = append(sharedGames, vUg1)
		}
	}

	return sharedGames
}
