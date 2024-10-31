package models

// TODO
// *Create struct for original json card table: FileCard
// *For each created table we need a matching struct: Card, Printcard, Langcard, etc
// *Make new Folder in dataload project called models - models.go (package models at top) - all structs go in here
type ImageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type MtgSet struct {
	SetId          string `json:"set_id"`
	SetCode        string `json:"set_code"`
	SetName        string `json:"Set_name"`
	SetType        string `json:"set_type"`
	SetUri         string `json:"set_uri"`
	SetSearchUri   string `json:"set_search_uri"`
	ScryfallSetUri string `json:"scryfall_set_uri"`
}

type AllParts struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Component string `json:"component"`
	Name      string `json:"name"`
	TypeLine  string `json:"type_line"`
	Uri       string `json:"uri"`
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

type FileCard struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	OracleId      string    `json:"oracle_id"`
	MultiverseIds []int     `json:"multiverse_ids"`
	MtgoId        int       `json:"mtgo_id"`
	ArenaId       int       `json:"arena_id"`
	TcgplayerId   int       `json:""`
	Name          string    `json:"name"`
	Lang          string    `json:"lang"`
	ReleasedAt    string    `json:"released_at"`
	Uri           string    `json:"uri"`
	ScryfallUri   string    `json:"scryfall_uri"`
	Layout        string    `json:"layout"`
	HighresImage  bool      `json:"highres_image"`
	ImageStatus   string    `json:"image_status"`
	ImageUris     ImageUris `json:"image_uris"`
	ManaCost      string    `json:"mana_cost"`
	Cmc           float32   `json:"cmc"`
	TypeLine      string    `json:"type_line"`
	OracleText    string    `json:"oracle_text"`
	//do we need to add more? theres more after oracle text in spirit token
	Power           string       `json:"power"`
	Toughness       string       `json:"toughness"`
	Colors          []string     `json:"colors"`
	ColorIdentity   []string     `json:"color_identity"`
	Keywords        []string     `json:"keywords"`
	AllParts        []AllParts   `json:"all_parts"`
	Legalities      Legalities   `json:"legalities"`
	Games           []string     `json:"games"`
	Reserved        bool         `json:"reserved"`
	Foil            bool         `json:"foil"`
	NonFoil         bool         `json:"nonfoil"`
	Finishes        []string     `json:"finishes"`
	Oversized       bool         `json:"oversized"`
	Promo           bool         `json:"promo"`
	Reprint         bool         `json:"reprint"`
	Variation       bool         `json:"variation"`
	SetId           string       `json:"set_id"`
	Set             string       `json:"set"`
	SetName         string       `json:"set_name"`
	SetType         string       `json:"set_type"`
	SetUri          string       `json:"set_uri"`
	SetSearchUri    string       `json:"set_search_uri"`
	ScryfallSetUri  string       `json:"scryfall_set_uri"`
	RulingsUri      string       `json:"rulings_uri"`
	PrintsSearchUri string       `json:"prints_search_uri"`
	CollectorNumber string       `json:"collector_number"`
	Digital         bool         `json:"digital"`
	Rarity          string       `json:"rarity"`
	CardBackId      string       `json:"card_back_id"`
	Artist          string       `json:"artist"`
	ArtistIds       []string     `json:"artist_ids"`
	IllustrationId  string       `json:"illustration_id"`
	BorderColor     string       `json:"border_color"`
	Frame           string       `json:"frame"`
	FullArt         bool         `json:"full_art"`
	Textless        bool         `json:"textless"`
	Booster         bool         `json:"booster"`
	StorySpotlight  bool         `json:"story_spotlight"`
	PromoTypes      []string     `json:"promo_types"`
	Prices          Prices       `json:"prices"`
	RelatedUris     RelatedUris  `json:"related_uris"`
	PurchaseUris    PurchaseUris `json:"purchase_uris"`
	PrintedName     string       `json:"printed_name"`
	PrintedText     string       `json:"printed_text"`
	PrintedTypeLine string       `json:"printed_type_line"`
	FlavorText      string       `json:"flavor_text"`
	//CardFaces       []string     `json:"card_faces"`
}

type Cards struct {
	OracleId      string     `json:"oracle_id"`
	ScryfallUri   string     `json:"scryfall_uri"`
	Layout        string     `json:"layout"`
	ManaCost      string     `json:"mana_cost"`
	Cmc           float32    `json:"cmc"`
	TypeLine      string     `json:"type_line"`
	OracleText    string     `json:"oracle_text"`
	Power         string     `json:"power"`
	Toughness     string     `json:"toughness"`
	Colors        []string   `json:"colors"`
	ColorIdentity []string   `json:"color_identity"`
	Keywords      []string   `json:"keywords"`
	AllParts      []string   `json:"all_parts"`
	ProducedMana  []string   `json:"produced_mana"`
	Reserved      bool       `json:"reserved"`
	RulingsUri    string     `json:"rulings_uri"`
	Legalities    Legalities `json:"legalities"`
	CardFaces     []string   `json:"card_faces"`
}

type Prints struct {
	PrintId         int         `json:"print_id"`
	MultiverseIds   int         `json:"multiverse_ids"`
	MtgoId          int         `json:"mtgo_id"`
	ArenaId         int         `json:"arena_id"`
	TcgplayerId     int         `json:"tcgplayer_id"`
	ReleasedAt      string      `json:"released_at"`
	Images          ImageUris   `json:"images"`
	Games           []string    `json:"games"`
	Oversized       bool        `json:"oversized"`
	SetId           string      `json:"set_id"`
	CollectorNumber string      `json:"collector_number"`
	Digital         bool        `json:"digital"`
	Rarity          string      `json:"rarity"`
	Card_backId     string      `json:"card_back_id"`
	Artist          string      `json:"artist"`
	IllustrationId  string      `json:"illustration_id"`
	BorderColor     string      `json:"border_color"`
	Frame           string      `json:"frame"`
	FullArt         bool        `json:"full_art"`
	Textless        bool        `json:"textless"`
	Booster         bool        `json:"booster"`
	StorySpotlight  bool        `json:"story_spotlight"`
	RelatedArticles RelatedUris `json:"related_articles"`
	CardName        string      `json:"name"`
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
	Reprint         bool     `json:"reprint"`
	Variation       bool     `json:"variation"`
	Price           Prices   `json:"prices"`
	PrintId         int      `json:"print_id"`
	PrintedName     string   `json:"printed_name"`
	PrintedText     string   `json:"printed_text"`
	PrintedTypeLine string   `json:"printed_type_line"`
	//FlavorText      string   `json:"flavor_text"`
}
