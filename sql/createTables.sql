DROP TABLE IF EXISTS lang_multiverse_id, lang_finish, lang_promo, lang_purchase_uri, print_lang, print_related, print_border_effect, print_frame_effect, print_game, prints, card_color, card_color_identity, card_produced_mana, card_keyword, card_color_indicator, print_attraction_light, card_faces_color, card_faces_color_indicator, print_lang_card_faces, card_faces, cards, mtg_set, related CASCADE;

CREATE TABLE related (
    object_parts text,
    id text,
    component text,
    card_name text,
    type_line text,
    uri text,
    PRIMARY KEY (id)
);

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
    oracle_id text not null, 
    object text not null,
    card_name text not null,
    layout text not null,
    mana_cost text,
    cmc decimal not null,
    type_line text not null,
    power text,
    toughness text,
    reserved boolean not null,
    standard_f boolean,
    future_f boolean,
    historic_f boolean,
    timeless_f boolean,
    gladiator_f boolean,
    pioneer_f boolean,
    explorer_f boolean,
    modern_f boolean,
    legacy_f boolean,
    pauper_f boolean,
    vintage_f boolean,
    penny_f boolean,
    commander_f boolean,
    oathbreaker_f boolean,
    standardbrawl_f boolean,
    brawl_f boolean,
    alchemy_f boolean,
    paupercommander_f boolean,
    duel_f boolean,
    premodern_f boolean,
    predh_f boolean,
    defense text,
    loyalty text,
    edhrec_rank integer,
    hand_modifier text,
    life_modifier text,

    PRIMARY KEY (card_name, oracle_id)
);

CREATE TABLE card_color (
    card_name text ,
    oracle_id text ,
    color text,
    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (color, card_name, oracle_id)
);

CREATE TABLE card_color_identity (
    card_name text ,
    oracle_id text ,
    color text,
    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (color, card_name, oracle_id)
);

CREATE TABLE card_produced_mana (
    card_name text ,
    oracle_id text ,
    color text,
    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (color, card_name, oracle_id)
);

CREATE TABLE card_color_indicator (
    card_name text ,
    oracle_id text ,
    color text,
    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (color, card_name, oracle_id)
);

CREATE TABLE card_keyword (
    card_name text ,
    keyword text,
    oracle_id text ,
    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (keyword, card_name, oracle_id)
);

CREATE TABLE card_faces (
    card_name          text,
    artist          text,
    artist_id        text,
    cmc             decimal,
    defense         text,
    flavor_text      text,
    illustration_id  text,
    png_uri text,
    boarder_crop_uri text,
    art_crop_uri text,
    large_uri text,
    normal_uri text,
    small_uri text,
    layout          text,
    loyalty        text,
    mana_cost       text,
    object_type         text,
    oracle_id        text,
    oracle_text     text,
    power           text,
    printed_name    text,
    printed_text    text,
    printed_type_line text,
    toughness      text,
    type_line       text,
    watermark      text,

    PRIMARY KEY (card_name)
);

CREATE TABLE card_faces_color (
    card_name text REFERENCES card_faces (card_name),
    color text,

    PRIMARY KEY (card_name, color)
);

CREATE TABLE card_faces_color_indicator (
    card_name text REFERENCES card_faces (card_name),
    color text,

    PRIMARY KEY (card_name, color)
);


CREATE TABLE prints (
    mtgo_id integer, 
    mtgo_foil_id integer,
    arena_id integer,
    tcgplayer_id integer,
    tcgplayer_etched_id integer,
    released_at text not null,
    oversized boolean not null,
    set_id text REFERENCES mtg_set (set_id),
    oracle_text text,
    collector_number text,
    digital boolean not null,
    oldschool_f boolean,
    rarity text,
    card_back_id text not null, --is the normal mtg card back considered a card in this database
    artist text,
    illustration_id text,
    border_color text,
    frame text,
    full_art boolean not null,
    textless boolean not null,
    booster boolean not null,
    story_spotlight boolean not null,
    tcg_articles_uri text,
    tcg_decks_uri text,
    edhrec_uri text,
    tcg_buy_uri text,
    cardmarket_buy_uri text,
    cardhoarder_buy_uri text,
    oracle_id text,
    card_name text,
    prints_search_uri text,
    flavor_name text,
    security_stamp text,
    previewed_at text,
    previewed_source_uri text,
    previewsource text,
    content_warning boolean,

    FOREIGN KEY(card_name, oracle_id) REFERENCES cards(card_name, oracle_id),
    PRIMARY KEY (card_name, set_id) --card_name, set_id, booster
);

CREATE TABLE print_attraction_light (
    card_name text,
    set_id text,
    attraction_light integer,
    FOREIGN KEY (card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (attraction_light, card_name, set_id)
);


CREATE TABLE print_game (
    card_name text, 
    set_id text,
    game text,
    FOREIGN KEY (card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (card_name, set_id, game)
);

CREATE TABLE print_border_effect (
    card_name text, 
    set_id text,
    border_effect text,
    FOREIGN KEY (card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (card_name, set_id, border_effect)
);

CREATE TABLE print_frame_effect (
    card_name text, 
    set_id text,
    frame_effect text,
    FOREIGN KEY (card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (card_name, set_id, frame_effect)
);
    
CREATE TABLE print_related (
    print_card_name text, 
    set_id text,
    related_id text REFERENCES related (id),
    FOREIGN KEY (print_card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (print_card_name, set_id, related_id)
);

CREATE TABLE print_lang (
    lang text not null,
    scryfall_uri_json text not null,
    scryfall_uri text not null,
    rulings_uri text,
    gatherer_uri text,
    highres_image boolean not null,
    image_status text not null,
    foil boolean not null,
    not_foil boolean not null,
    promo boolean not null,
    reprint boolean not null,
    variation boolean not null, --NOTE figure out if print can have different values foil-variation
    variation_of text, --in scryfall this says uuid, not sure if this is text
    --(I just learned variation is basically a widespread misprint)
    price_usd text,
    price_usd_foil text,
    price_usd_etched text,
    price_eur text,
    price_eur_foil text,
    price_tix text,
    printed_name     text,
    printed_next     text,    
    printed_type_line text,
    cardmarket_id integer,
    watermark text,
    png_uri text,
    boarder_crop_uri text,
    art_crop_uri text,
    large_uri text,
    normal_uri text,
    small_uri text,
    card_name text,
    set_id text,
    FOREIGN KEY (card_name, set_id) REFERENCES prints(card_name, set_id),
    PRIMARY KEY (card_name, set_id, lang)
);

CREATE TABLE print_lang_card_faces (
    lang text,
    card_name text,
    set_id text,
    card_faces_card_name text REFERENCES card_faces (card_name),

    FOREIGN KEY (card_name, set_id, lang) REFERENCES print_lang (card_name, set_id, lang),
    PRIMARY KEY (lang, card_name, set_id, card_faces_card_name)
);

CREATE TABLE lang_finish (
    card_name text,
    set_id text,
    lang text,
    finish text,

    FOREIGN KEY (card_name, set_id, lang) REFERENCES print_lang (card_name, set_id, lang),
    PRIMARY KEY (card_name, set_id, lang, finish)
);

CREATE TABLE lang_promo (
    card_name text,
    set_id text,
    lang text,
    promo text,

    FOREIGN KEY (card_name, set_id, lang) REFERENCES print_lang (card_name, set_id, lang),
    PRIMARY KEY (card_name, set_id, lang, promo)
);

CREATE TABLE lang_purchase_uri (
    card_name text,
    set_id text,
    lang text,
    purchase_uri text,

    FOREIGN KEY (card_name, set_id, lang) REFERENCES print_lang (card_name, set_id, lang),
    PRIMARY KEY (card_name, set_id, lang, purchase_uri)
);

CREATE TABLE lang_multiverse_id (
    card_name text, 
    set_id text,
    lang text,
    multiverse_id integer,
    FOREIGN KEY (card_name, set_id, lang) REFERENCES print_lang(card_name, set_id, lang),
    PRIMARY KEY (card_name, set_id, lang, multiverse_id)
);
