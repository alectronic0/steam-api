package steamservice

import (
	"steam-api/internal/steamclient"
	"steam-api/pkg/utils"
	"strconv"
	"strings"
)

var prefixes = []string{
	"IAuthenticationService",
	"IBroadcastService",
	"ICSGO",
	"ICheatReportingService",
	"IClientStats_1046930",
	"IContentServerConfigService",
	"IContentServerDirectoryService",
	"IDOTA2",
	"IEcon",
	"IGC",
	"IGameNotificationsService",
	"IGameServersService",
	"IHelpRequestLogsService",
	"IInventoryService",
	"IPlayerService/IsPlayingSharedGame",
	"IPlayerService/RecordOfflinePlaytime",
	"IPortal2Leaderboards_620",
	"IPublishedFileService",
	"ISteamApps/GetAppList/v0001",
	"ISteamApps/GetSDRConfig",
	"ISteamApps/GetServersAtAddress",
	"ISteamApps/UpToDateCheck",
	"ISteamBroadcast",
	"ISteamCDN",
	"ISteamDirectory",
	"ISteamEconomy",
	"ISteamNews",
	"ISteamRemoteStorage",
	"ISteamUser/GetPlayerSummaries/v0001",
	"ISteamUserAuth/AuthenticateUserTicket",
	"ISteamUserOAuth",
	"ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0001",
	"ISteamUserStats/GetSchemaForGame/v0001",
	"ISteamUserStats/GetUserStatsForGame/v0001",
	"ISteamWebAPIUtil/GetServerInfo",
	"IStoreService/GetRecommendedTagsForUser",
	"ITF",
	"IWishlistService/GetWishlistSortedFiltered",
}

type SupportedAPIList map[string]SteamAPI

type SteamAPI struct {
	HttpMethod string               `json:"http_method"`
	Url        string               `json:"url"`
	Parameters map[string]Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

const (
	defaultString = "XXXXXXXXXXXXXXXXXXXXXXX"
	defaultNumber = "0"
	defaultBool   = "true"
)

var defaultText = map[string]string{
	"uint32":    defaultNumber,
	"uint64":    defaultNumber,
	"int32":     defaultNumber,
	"{enum}":    defaultString,
	"{message}": defaultString,
	"string":    defaultString,
	"bool":      defaultBool,
}

func SupportedAPIListFromAPI(m *steamclient.GetSupportedAPIListAPIResponse, apiKey string, steamID, appID uint64) SupportedAPIList {
	if m == nil {
		return SupportedAPIList{}
	}

	apiList := map[string]SteamAPI{}
	for _, i := range m.ApiList.Interfaces {
		interfaceName := i.Name
		for _, meathod := range i.Methods {
			var params []string
			parameters := map[string]Parameter{}
			for _, p := range meathod.Parameters {
				parameters[p.Name] = Parameter{
					Name:        p.Name,
					Type:        p.Type,
					Optional:    p.Optional,
					Description: p.Description,
				}

				switch p.Name {
				case "key":
					params = append(params, p.Name+"="+apiKey)
				case "steamid":
					params = append(params, p.Name+"="+strconv.FormatUint(steamID, 10))
				case "steamids":
					params = append(params, p.Name+"="+strconv.FormatUint(steamID, 10))
				case "appid":
					params = append(params, p.Name+"="+strconv.FormatUint(appID, 10))
				case "gameid":
					params = append(params, p.Name+"="+strconv.FormatUint(appID, 10))
				case "vanityurl":
					params = append(params, p.Name+"="+"alectronic0")
				case "url_type":
					params = append(params, p.Name+"="+"1")
				case "language":
					params = append(params, p.Name+"="+"en")
				case "country_code":
					params = append(params, p.Name+"="+"gb")
				default:
					params = append(params, p.Name+"="+defaultText[p.Type])
				}
			}

			methodName := meathod.Name
			apiName := interfaceName + "/" + methodName + "/v000" + strconv.Itoa(meathod.Version)
			urlString := steamclient.BaseApiURL + "/" + apiName + "/?" + strings.Join(params, "&")

			if !utils.StartsWithAny(apiName, prefixes...) {
				apiList[apiName] = SteamAPI{
					HttpMethod: meathod.HttpMethod,
					Url:        urlString,
					Parameters: parameters,
				}
			}
		}
	}

	return apiList
}
