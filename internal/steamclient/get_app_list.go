package steamclient

type GetAppListAPIResponse struct {
	AppList AppList `json:"applist"`
}

type App struct {
	AppID string `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps []App `json:"apps"`
}
