package models

type MtgSet struct {
	SetId          string `json:"set_id"`
	SetCode        string `json:"set_code"`
	SetName        string `json:"Set_name"`
	SetType        string `json:"set_type"`
	SetUri         string `json:"set_uri"`
	SetSearchUri   string `json:"set_search_uri"`
	ScryfallSetUri string `json:"scryfall_set_uri"`
}

type Related struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Component string `json:"component"`
	Name      string `json:"name"`
	TypeLine  string `json:"type_line"`
	Uri       string `json:"uri"`
}

type Cards struct {
	Object           string `json:"object"`
	OracleId         string `json:"oracle_id"`
	CardName         string
	ScryfallUri      string
	Layout           string   `json:"layout"`
	ManaCost         string   `json:"mana_cost"`
	Cmc              float64  `json:"cmc"`
	TypeLine         string   `json:"type_line"`
	OracleText       string   `json:"oracle_text"`
	Power            string   `json:"power"`
	Toughness        string   `json:"toughness"`
	Colors           []string `json:"colors"`
	ColorIdentity    []string `json:"color_identity"`
	Keywords         []string `json:"keywords"`
	ProducedMana     []string `json:"produced_mana"`
	AttractionLights []int
	Reserved         bool   `json:"reserved"`
	RulingsUri       string `json:"rulings_uri"`
	StandardF        bool
	FutureF          bool
	HistoricF        bool
	TimelessF        bool
	GladiatorF       bool
	PioneerF         bool
	ExplorerF        bool
	ModernF          bool
	LegacyF          bool
	PauperF          bool
	VintageF         bool
	PennyF           bool
	CommanderF       bool
	OathbreakerF     bool
	StandardbrawlF   bool
	BrawlF           bool
	AlchemyF         bool
	PaupercommanderF bool
	DuelF            bool
	OldschoolF       bool
	PremodernF       bool
	PredhF           bool
	CardFaces        []CardFaces `json:"card_faces"`
	ColorIndicator   []string    `json:"color_indicator"`
	Defense          string      `json:"defense"`
	Loyalty          string      `json:"loyalty"`
	EdhrecRank       int         `json:"edhrec_rank"`
	HandModifier     string      `json:"hand_modifier"`
	LifeModifier     string      `json:"life_modifier"`
	PennyRank        int         `json:"penny_rank"`
	ContentWarning   bool        `json:"content_warning"`
}

type CardFaces struct {
	Artist          string   `json:"artist"`
	ArtistId        string   `json:"artist_id"`
	Cmc             float64  `json:"cmc"`
	ColorIndicator  []string `json:"color_indicator"`
	Colors          []string `json:"colors"`
	Defense         string   `json:"defense"`
	FlavorText      string   `json:"flavor_text"`
	IllustrationId  string   `json:"illustration_id"`
	PngUri          string
	BoarderCropUri  string
	ArtCropUri      string
	LargeUri        string
	NormalUri       string
	SmallUri        string
	Layout          string `json:"layout"`
	Loyalty         string `json:"loyalty"`
	ManaCost        string `json:"mana_cost"`
	Name            string `json:"name"`
	Object          string `json:"object"`
	OracleId        string `json:"oracle_id"`
	OracleText      string `json:"oracle_text"`
	Power           string `json:"power"`
	PrintedName     string `json:"printed_name"`
	PrintedText     string `json:"printed_text"`
	PrintedTypeLine string `json:"printed_type_line"`
	Toughness       string `json:"toughness"`
	TypeLine        string `json:"type_line"`
	Watermark       string `json:"watermark"`
}

type Prints struct {
	PrintId           int      `json:"print_id"`
	MultiverseIds     []int    `json:"multiverse_ids"`
	MtgoId            int      `json:"mtgo_id"`
	MtgoFoilId        int      `json:"mtgo_foil_id"`
	ArenaId           int      `json:"arena_id"`
	TcgplayerId       int      `json:"tcgplayer_id"`
	TcgplayerEtchedId int      `json:"tcgplayer_etched_id"`
	ReleasedAt        string   `json:"released_at"`
	Games             []string `json:"games"`
	Oversized         bool     `json:"oversized"`
	SetId             string   `json:"set_id"`
	CollectorNumber   string   `json:"collector_number"`
	Digital           bool     `json:"digital"`
	Rarity            string   `json:"rarity"`
	Card_backId       string   `json:"card_back_id"`
	Artist            string   `json:"artist"`
	IllustrationId    string   `json:"illustration_id"`
	BorderColor       string   `json:"border_color"`
	Frame             string   `json:"frame"`
	FullArt           bool     `json:"full_art"`
	Textless          bool     `json:"textless"`
	Booster           bool     `json:"booster"`
	StorySpotlight    bool     `json:"story_spotlight"`
	GathererUri       string
	TcgArticlesUri    string
	TcgDecksUri       string
	EdhrecUri         string
	TcgBuyUri         string
	CardmarketBuyUri  string
	CardhoarderBuyUri string
	CardName          string    `json:"name"`
	PrintsSearchUri   string    `json:"prints_search_uri"`
	AllParts          []Related `json:"all_parts"`
	FlavorName        string    `json:"flavor_name"`
	FrameEffects      []string  `json:"frame_effects"`
	SecurityStamp     string    `json:"security_stamp"`
	Previewed_at      string    `json:"preview.previewed_at"`
	PreviewUri        string    `json:"preview.source_uri"`
	PreviewSource     string    `json:"preview.source"`
}

type PrintLangs struct {
	Lang            string   `json:"lang"`
	ScryfallUriJson string   `json:"scryfall_uri"`
	HighresImage    bool     `json:"highres_image"`
	ImageStatus     string   `json:"image_status"`
	Foil            bool     `json:"foil"`
	NotFoil         bool     `json:"nonfoil"`
	Finishes        []string `json:"finishes"`
	Promo           bool     `json:"promo"`
	PromoTypes      []string `json:"promo_types"`
	Reprint         bool     `json:"reprint"`
	Variation       bool     `json:"variation"`
	VariationOf     string   `json:"variation_of"`
	PriceUsd        string
	PriceUsdFoil    string
	PriceUsdEtched  string
	PriceEur        string
	PriceEurFoil    string
	PriceTix        string
	PrintId         int    `json:"print_id"`
	PrintedName     string `json:"printed_name"`
	PrintedText     string `json:"printed_text"`
	PrintedTypeLine string `json:"printed_type_line"`
	FlavorText      string `json:"flavor_text"`
	CardmarketId    int    `json:"cardmarket_id"`
	Uri             string `json:"uri"`
	Id              string `json:"id"`
}
