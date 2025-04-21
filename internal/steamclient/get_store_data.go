package steamclient

type GetStoreDataAPIResponse map[string]StoreDataResponse

type StoreDataResponse struct {
	Success bool      `json:"success"`
	Data    StoreData `json:"data"`
}

type StoreData struct {
	Type                string             `json:"type"`
	Name                string             `json:"name"`
	SteamAppID          int                `json:"steam_appid"`
	RequiredAge         int                `json:"required_age"`
	IsFree              bool               `json:"is_free"`
	ControllerSupport   string             `json:"controller_support"`
	Dlc                 []int              `json:"dlc"`
	DetailedDescription string             `json:"detailed_description"`
	AboutTheGame        string             `json:"about_the_game"`
	ShortDescription    string             `json:"short_description"`
	SupportedLanguages  string             `json:"supported_languages"`
	Reviews             string             `json:"reviews"`
	HeaderImage         string             `json:"header_image"`
	CapsuleImage        string             `json:"capsule_image"`
	CapsuleImagev5      string             `json:"capsule_imagev5"`
	Website             string             `json:"website"`
	PcRequirements      Requirements       `json:"pc_requirements"`
	MacRequirements     Requirements       `json:"mac_requirements"`
	LinuxRequirements   Requirements       `json:"linux_requirements"`
	LegalNotice         string             `json:"legal_notice"`
	Developers          []string           `json:"developers"`
	Publishers          []string           `json:"publishers"`
	PriceOverview       PriceOverview      `json:"price_overview"`
	Packages            []int              `json:"packages"`
	PackageGroups       []PackageGroup     `json:"package_groups"`
	Platforms           Platforms          `json:"platforms"`
	Metacritic          Metacritic         `json:"metacritic"`
	Categories          []Tag              `json:"categories"`
	Genres              []Tag              `json:"genres"`
	Screenshots         []Screenshot       `json:"screenshots"`
	Movies              []Movie            `json:"movies"`
	Recommendations     Recommendations    `json:"recommendations"`
	Achievements        Achievements       `json:"achievements"`
	ReleaseDate         ReleaseDate        `json:"release_date"`
	SupportInfo         SupportInfo        `json:"support_info"`
	Background          string             `json:"background"`
	BackgroundRaw       string             `json:"background_raw"`
	ContentDescriptors  ContentDescriptors `json:"content_descriptors"`
	Ratings             Ratings            `json:"ratings"`
}

type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type PriceOverview struct {
	Currency         string `json:"currency"`
	Initial          int    `json:"initial"`
	Final            int    `json:"final"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}

type PackageGroup struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	SelectionText           string `json:"selection_text"`
	SaveText                string `json:"save_text"`
	DisplayType             int    `json:"display_type"`
	IsRecurringSubscription string `json:"is_recurring_subscription"`
	Subs                    []Sub  `json:"subs"`
}

type Sub struct {
	PackageID                int    `json:"packageid"`
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}
type Metacritic struct {
	Score int    `json:"score"`
	Url   string `json:"url"`
}

type Tag struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Screenshot struct {
	ID            string `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type Movie struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Webm      Video  `json:"webm"`
	Mp4       Video  `json:"mp4"`
	Highlight bool   `json:"highlight"`
}

type Video struct {
	Field1 string `json:"480"`
	Max    string `json:"max"`
}

type Recommendations struct {
	Total int `json:"total"`
}

type Achievements struct {
	Total       int         `json:"total"`
	Highlighted []Highlight `json:"highlighted"`
}

type Highlight struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type SupportInfo struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type ContentDescriptors struct {
	IDs   []uint64 `json:"ids"`
	Notes string   `json:"notes"`
}

type Ratings struct {
	Esrb         SimpleRatings   `json:"esrb"`
	Dejus        DetailedRatings `json:"dejus"`
	SteamGermany DetailedRatings `json:"steam_germany"`
}
type SimpleRatings struct {
	Rating      string `json:"rating"`
	Descriptors string `json:"descriptors"`
}

type DetailedRatings struct {
	RatingGenerated string `json:"rating_generated"`
	Rating          string `json:"rating"`
	RequiredAge     string `json:"required_age"`
	Banned          string `json:"banned"`
	UseAgeGate      string `json:"use_age_gate"`
	Descriptors     string `json:"descriptors"`
}
