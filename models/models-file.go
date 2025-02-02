package models

type FileCard struct {
	Object            string       `json:"object"`
	Id                string       `json:"id"`
	OracleId          string       `json:"oracle_id"`
	MultiverseIds     []int        `json:"multiverse_ids"`
	MtgoId            int          `json:"mtgo_id"`
	ArenaId           int          `json:"arena_id"`
	TcgplayerId       int          `json:""`
	Name              string       `json:"name"`
	Lang              string       `json:"lang"`
	ReleasedAt        string       `json:"released_at"`
	Uri               string       `json:"uri"`
	ScryfallUri       string       `json:"scryfall_uri"`
	Layout            string       `json:"layout"`
	HighresImage      bool         `json:"highres_image"`
	ImageStatus       string       `json:"image_status"`
	ImageUris         ImageUris    `json:"image_uris"`
	ManaCost          string       `json:"mana_cost"`
	Cmc               float64      `json:"cmc"`
	TypeLine          string       `json:"type_line"`
	OracleText        string       `json:"oracle_text"`
	Power             string       `json:"power"`
	Toughness         string       `json:"toughness"`
	Colors            []string     `json:"colors"`
	ColorIdentity     []string     `json:"color_identity"`
	Keywords          []string     `json:"keywords"`
	ProducedMana      []string     `json:"produced_mana"`
	AllParts          []Related    `json:"all_parts"`
	Legalities        Legalities   `json:"legalities"`
	Games             []string     `json:"games"`
	Reserved          bool         `json:"reserved"`
	Foil              bool         `json:"foil"`
	NonFoil           bool         `json:"nonfoil"`
	Finishes          []string     `json:"finishes"`
	Oversized         bool         `json:"oversized"`
	Promo             bool         `json:"promo"`
	Reprint           bool         `json:"reprint"`
	Variation         bool         `json:"variation"`
	VariationOf       string       `json:"variation_of"`
	SetId             string       `json:"set_id"`
	Set               string       `json:"set"`
	SetName           string       `json:"set_name"`
	SetType           string       `json:"set_type"`
	SetUri            string       `json:"set_uri"`
	SetSearchUri      string       `json:"set_search_uri"`
	ScryfallSetUri    string       `json:"scryfall_set_uri"`
	RulingsUri        string       `json:"rulings_uri"`
	PrintsSearchUri   string       `json:"prints_search_uri"`
	CollectorNumber   string       `json:"collector_number"`
	Digital           bool         `json:"digital"`
	Rarity            string       `json:"rarity"`
	CardBackId        string       `json:"card_back_id"`
	Artist            string       `json:"artist"`
	ArtistIds         []string     `json:"artist_ids"`
	IllustrationId    string       `json:"illustration_id"`
	BorderColor       string       `json:"border_color"`
	Frame             string       `json:"frame"`
	FullArt           bool         `json:"full_art"`
	Textless          bool         `json:"textless"`
	Booster           bool         `json:"booster"`
	StorySpotlight    bool         `json:"story_spotlight"`
	PromoTypes        []string     `json:"promo_types"`
	Prices            Prices       `json:"prices"`
	RelatedUris       RelatedUris  `json:"related_uris"`
	PurchaseUris      PurchaseUris `json:"purchase_uris"`
	PrintedName       string       `json:"printed_name"`
	PrintedText       string       `json:"printed_text"`
	PrintedTypeLine   string       `json:"printed_type_line"`
	FlavorText        string       `json:"flavor_text"`
	CardFaces         []CardFaces  `json:"card_faces"`
	TcgplayerEtchedId int          `json:"tcgplayer_etched_id"`
	MtgoFoilId        int          `json:"mtgo_foil_id"`
	CardmarketId      int          `json:"cardmarket_id"`
	ColorIndicator    []string     `json:"color_indicator"`
	Defense           string       `json:"defense"`
	Loyalty           string       `json:"loyalty"`
	EdhrecRank        int          `json:"edhrec_rank"`
	HandModifier      string       `json:"hand_modifier"`
	LifeModifier      string       `json:"life_modifier"`
	PennyRank         int          `json:"penny_rank"`
	AttractionLights  []int        `json:"attraction_lights"`
	ContentWarning    bool         `json:"content_warning"`
	FlavorName        string       `json:"flavor_name"`
	FrameEffects      []string     `json:"frame_effects"`
	SecurityStamp     string       `json:"security_stamp"`
	Previewed_at      string       `json:"preview.previewed_at"`
	PreviewUri        string       `json:"preview.source_uri"`
	PreviewSource     string       `json:"preview.source"`
}

type ImageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Legalities struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Timeless        string `json:"timeless"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Explorer        string `json:"explorer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Oathbreaker     string `json:"oathbreaker"`
	StandardBrawl   string `json:"standardbrawl"`
	Brawl           string `json:"brawl"`
	Alchemy         string `json:"alchemy"`
	PauperCommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	Oldschool       string `json:"oldschool"`
	Premodern       string `json:"premodern"`
	Predh           string `json:"predh"`
}

type Prices struct {
	Usd       string `json:"usd"`
	UsdFoil   string `json:"usd_foil"`
	UsdEtched string `json:"usd_etched"`
	Eur       string `json:"eur"`
	EurFoil   string `json:"eur_foil"`
	Tix       string `json:"tix"`
}

type RelatedUris struct {
	TcgInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TcgInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	Edhrec              string `json:"edhrec"`
}

type PurchaseUris struct {
	TcgPlayer   string `json:"tcgplayer"`
	CardMarket  string `json:"cardmarket"`
	CardHoarder string `json:"cardhoarder"`
}
