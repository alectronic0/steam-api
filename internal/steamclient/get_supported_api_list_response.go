package steamclient

type GetSupportedAPIListAPIResponse struct {
	ApiList ApiList `json:"apiList"`
}

type ApiList struct {
	Interface []Interface `json:"interfaces"`
}

type Interface struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name       string      `json:"name"`
	Version    int         `json:"version"`
	HttpMethod string      `json:"httpMethod"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}
