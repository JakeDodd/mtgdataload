DROP TABLE IF EXISTS print_lang, prints, card_keywords, cards, mtg_set, keywords;
DROP TYPE IF EXISTS card_faces, related_uris, prices, legality, image_uris, all_parts;



Create TYPE all_parts as(
    object_parts text,
    id text,
    component text,
    card_name text,
    type_line text,
    uri text

);

CREATE TYPE legality as(
    standard_f text,
    future_f text,
    historic_f text,
    timeless_f text,
    gladiator_f text,
    pioneer_f text,
    explorer_f text,
    modern_f text,
    legacy_f text,
    pauper_f text,
    vintage_f text,
    penny_f text,
    commander_f text,
    oathbreaker_f text,
    standardbrawl_f text,
    brawl_f text,
    alchemy_f text,
    paupercommander_f text,
    duel_f text,
    oldschool_f text,
    premodern_f text,
    predh_f text
);
CREATE TYPE prices as(
    usd text,
    usd_foil text,
    usd_etched text,
    eur text,
    eur_foil text,
    tix text

);
CREATE TYPE related_uris as(
    gatherer text,
    tcg_articles text,
    tcg_decks text,
    edhrec text,
    tcg_buy text,
    cardmarket_buy text,
    cardhoarder_buy text
);

Create TYPE card_faces as(
	artist          text,
	artist_id        text,
	cmc             integer,
	color_indicator  text[],
	colors          text[],
	defense         text,
	flavor_text      text,
	illustration_id  text,
	image_uris       image_uris,
	layout          text,
	loyalty        text,
	mana_cost       text,
	card_name          text,
	object_type         text,
	oracle_id        text,
	oracle_text     text,
	power           text,
	printed_name    text,
	printed_text    text,
	printed_type_line text,
	toughness      text,
	type_line       text,
	watermark      text
    
);
CREATE TABLE card_faces (
    
)

CREATE TABLE mtg_set (
    set_id text,
    set_code text,
    set_name text,
    set_type text,
    set_uri text,
    set_search_uri text,
    scryfall_set_uri text,
    PRIMARY KEY (set_id)
);
CREATE TABLE cards (
    object text not null,
    oracle_id text not null, 
    card_name text not null,
    scryfall_uri text not null,
    layout text not null,
    mana_cost text,
    cmc integer not null,
    type_line text not null,
    oracle_text text,
    power text,
    toughness text,
    color_identity text[],
    produced_mana text[],
    reserved boolean not null,
    rulings_uri text,
    legalities legality,
    defense text,
    loyalty text,
    color_indicator text[],
    card_faces card_faces[], --hold off for now
    edhrec_rank integer,
    hand_modifier text,
    life_modifier text,
    attraction_lights integer[],
    content_warning boolean,


    

    PRIMARY KEY (card_name)
);

CREATE TABLE card_colors (
    card_name text REFERENCES cards (card_name),
    color text,
    PRIMARY KEY (color, card_name)
)

CREATE TABLE card_keywords (
    keyword text,
    card_name text REFERENCES cards (card_name),
    PRIMARY KEY (keyword, card_name)
);
CREATE TABLE prints (
    print_id integer,
    multiverse_ids integer[],
    mtgo_id integer, 
    mtgo_foil_id integer,
    arena_id integer,
    tcgplayer_id integer,
    tcgplayer_etched_id integer,
    released_at text not null,
    images image_uris,
    games text[] not null,
    oversized boolean not null,
    set_id text REFERENCES mtg_set (set_id),
    collector_number text,
    digital boolean not null,
    rarity text,
    card_back_id text not null, --is the normal mtg card back considered a card in this database
    artist text,
    artist_ids text[],
    illustration_id text,
    border_color text,
    frame text,
    full_art boolean not null,
    textless boolean not null,
    booster boolean not null,
    story_spotlight boolean not null,
    related_articles related_uris,
    card_name text REFERENCES cards (card_name),
    prints_search_uri text,
    all_parts all_parts[], --dont touch for now
    flavor_name text,
    frame_effects text[],
    security_stamp text,
    previewed_at text,
    previewed_source_uri text,
    previewsource text,
    
    png_uri text,
    boarder_crop_uri text,
    art_crop_uri text,
    large_uri text,
    normal_uri text,
    small_uri text,

    PRIMARY KEY (print_id) --card_name, set_id, booster
);
CREATE TABLE print_lang (
    lang text not null,
    scryfall_uri_json text not null,
    highres_image boolean not null,
    image_status text not null,
    foil boolean not null,
    not_foil boolean not null,
    finishes text[],
    promo boolean not null,
    promo_types text[],
    reprint boolean not null,
    variation boolean not null, --NOTE figure out if print can have different values foil-variation
    variation_of text, --in scryfall this says uuid, not sure if this is text
    --(I just learned variation is basically a widespread misprint)
    price prices not null,
    print_id integer REFERENCES prints (print_id),
	printed_name     text,
	printed_next     text,    
	printed_type_line text,
    cardmarket_id integer,
    watermark text,
    purchase_uris text[],

    PRIMARY KEY (print_id, lang)
);