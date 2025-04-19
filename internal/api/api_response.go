package api

type Response struct {
	User1ID     User   `json:"user1"`
	User2ID     User   `json:"user2"`
	SharedGames []Game `json:"shared_games"`
}

type User struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Games []Game `json:"games"`
}

type Game struct {
	Name      string `json:"name"`
	AppID     int    `json:"appid"`
	ImgHeader string `json:"img_header"`
}
