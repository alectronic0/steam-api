package steamclient

type GetAppListAPIResponse struct {
	AppList AppList `json:"applist"`
}

type App struct {
	AppID uint64 `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps []App `json:"apps"`
}
