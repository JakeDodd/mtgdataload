package models

import (
	"reflect"
	"sort"
)

func ternary_f(s string) bool {
	if s == "legal" {
		return true
	} else {
		return false
	}
}

func (fc FileCard) FileCardToCard() Cards {
	card := Cards{
		OracleId:         fc.OracleId,
		Object:           fc.Object,
		CardName:         fc.Name,
		Layout:           fc.Layout,
		ManaCost:         fc.ManaCost,
		Cmc:              fc.Cmc,
		TypeLine:         fc.TypeLine,
		Power:            fc.Power,
		Toughness:        fc.Toughness,
		Colors:           fc.Colors,
		ColorIdentity:    fc.ColorIdentity,
		Keywords:         fc.Keywords,
		ProducedMana:     fc.ProducedMana,
		Reserved:         fc.Reserved,
		StandardF:        ternary_f(fc.Legalities.Standard),
		FutureF:          ternary_f(fc.Legalities.Future),
		HistoricF:        ternary_f(fc.Legalities.Historic),
		TimelessF:        ternary_f(fc.Legalities.Timeless),
		GladiatorF:       ternary_f(fc.Legalities.Gladiator),
		PioneerF:         ternary_f(fc.Legalities.Pioneer),
		ExplorerF:        ternary_f(fc.Legalities.Explorer),
		ModernF:          ternary_f(fc.Legalities.Modern),
		LegacyF:          ternary_f(fc.Legalities.Legacy),
		PauperF:          ternary_f(fc.Legalities.Pauper),
		VintageF:         ternary_f(fc.Legalities.Vintage),
		PennyF:           ternary_f(fc.Legalities.Penny),
		CommanderF:       ternary_f(fc.Legalities.Commander),
		OathbreakerF:     ternary_f(fc.Legalities.Oathbreaker),
		StandardbrawlF:   ternary_f(fc.Legalities.StandardBrawl),
		BrawlF:           ternary_f(fc.Legalities.Brawl),
		AlchemyF:         ternary_f(fc.Legalities.Alchemy),
		PaupercommanderF: ternary_f(fc.Legalities.PauperCommander),
		DuelF:            ternary_f(fc.Legalities.Duel),
		PremodernF:       ternary_f(fc.Legalities.Premodern),
		PredhF:           ternary_f(fc.Legalities.Predh),
		Defense:          fc.Defense,
		Loyalty:          fc.Loyalty,
		EdhrecRank:       fc.EdhrecRank,
		HandModifier:     fc.HandModifier,
		LifeModifier:     fc.LifeModifier,
		PennyRank:        fc.PennyRank,
	}
	return card
}

func (c1 Cards) CompareCards(c2 Cards) bool {
	// First, check if they are the same type
	if reflect.TypeOf(c1) != reflect.TypeOf(c2) {
		return false
	}

	// Compare the fields one by one
	return c1.Object == c2.Object &&
		c1.OracleId == c2.OracleId &&
		c1.CardName == c2.CardName &&
		c1.Layout == c2.Layout &&
		c1.ManaCost == c2.ManaCost &&
		c1.Cmc == c2.Cmc &&
		c1.TypeLine == c2.TypeLine &&
		c1.Power == c2.Power &&
		c1.Toughness == c2.Toughness &&
		compareStringSlices(c1.Colors, c2.Colors) &&
		compareStringSlices(c1.ColorIdentity, c2.ColorIdentity) &&
		compareStringSlices(c1.Keywords, c2.Keywords) &&
		compareStringSlices(c1.ProducedMana, c2.ProducedMana) &&
		c1.Reserved == c2.Reserved &&
		c1.StandardF == c2.StandardF &&
		c1.FutureF == c2.FutureF &&
		c1.HistoricF == c2.HistoricF &&
		c1.TimelessF == c2.TimelessF &&
		c1.GladiatorF == c2.GladiatorF &&
		c1.PioneerF == c2.PioneerF &&
		c1.ExplorerF == c2.ExplorerF &&
		c1.ModernF == c2.ModernF &&
		c1.LegacyF == c2.LegacyF &&
		c1.PauperF == c2.PauperF &&
		c1.VintageF == c2.VintageF &&
		c1.PennyF == c2.PennyF &&
		c1.CommanderF == c2.CommanderF &&
		c1.OathbreakerF == c2.OathbreakerF &&
		c1.StandardbrawlF == c2.StandardbrawlF &&
		c1.BrawlF == c2.BrawlF &&
		c1.AlchemyF == c2.AlchemyF &&
		c1.PaupercommanderF == c2.PaupercommanderF &&
		c1.DuelF == c2.DuelF &&
		c1.PremodernF == c2.PremodernF &&
		c1.PredhF == c2.PredhF &&
		compareStringSlices(c1.ColorIndicator, c2.ColorIndicator) &&
		c1.Defense == c2.Defense &&
		c1.Loyalty == c2.Loyalty &&
		c1.EdhrecRank == c2.EdhrecRank &&
		c1.HandModifier == c2.HandModifier &&
		c1.LifeModifier == c2.LifeModifier
}

func (fc FileCard) FileCardToSet() MtgSet {
	set := MtgSet{
		SetId:          fc.SetId,
		SetCode:        fc.Set,
		SetName:        fc.SetName,
		SetType:        fc.SetType,
		SetUri:         fc.SetUri,
		SetSearchUri:   fc.SetSearchUri,
		ScryfallSetUri: fc.ScryfallSetUri,
	}
	return set
}

func (fc FileCard) FileCardToPrint() Prints {
	print := Prints{
		MtgoId:            fc.MtgoId,
		MtgoFoilId:        fc.MtgoId,
		ArenaId:           fc.ArenaId,
		TcgplayerId:       fc.TcgplayerId,
		TcgplayerEtchedId: fc.TcgplayerEtchedId,
		ReleasedAt:        fc.ReleasedAt,
		Games:             fc.Games,
		AttractionLights:  fc.AttractionLights,
		Oversized:         fc.Oversized,
		SetId:             fc.SetId,
		OracleText:        fc.OracleText,
		CollectorNumber:   fc.CollectorNumber,
		Digital:           fc.Digital,
		Rarity:            fc.Rarity,
		OldschoolF:        ternary_f(fc.Legalities.Oldschool),
		CardBackId:        fc.CardBackId,
		Artist:            fc.Artist,
		IllustrationId:    fc.IllustrationId,
		BorderColor:       fc.BorderColor,
		Frame:             fc.Frame,
		FullArt:           fc.FullArt,
		Textless:          fc.Textless,
		Booster:           fc.Booster,
		StorySpotlight:    fc.StorySpotlight,
		TcgArticlesUri:    fc.RelatedUris.TcgInfiniteArticles,
		TcgDecksUri:       fc.RelatedUris.TcgInfiniteDecks,
		EdhrecUri:         fc.RelatedUris.Edhrec,
		TcgBuyUri:         fc.PurchaseUris.TcgPlayer,
		CardmarketBuyUri:  fc.PurchaseUris.CardMarket,
		CardhoarderBuyUri: fc.PurchaseUris.CardHoarder,
		OracleId:          fc.OracleId,
		CardName:          fc.Name,
		PrintsSearchUri:   fc.PrintsSearchUri,
		Related:           fc.AllParts,
		FlavorName:        fc.FlavorName,
		FrameEffects:      fc.FrameEffects,
		SecurityStamp:     fc.SecurityStamp,
		PreviewedAt:       fc.Previewed_at,
		PreviewUri:        fc.PreviewUri,
		PreviewSource:     fc.PreviewSource,
		ContentWarning:    fc.ContentWarning,
		BorderEffects:     fc.FrameEffects,
		Lang:              fc.Lang,
		MultiverseIds:     fc.MultiverseIds,
		GathererUri:       fc.RelatedUris.Gatherer,
		ScryfallUri:       fc.ScryfallUri,
		RulingsUri:        fc.RulingsUri,
		HighresImage:      fc.HighresImage,
		ImageStatus:       fc.ImageStatus,
		Foil:              fc.Foil,
		NotFoil:           fc.NonFoil,
		Finishes:          fc.Finishes,
		Promo:             fc.Promo,
		PromoTypes:        fc.PromoTypes,
		Reprint:           fc.Reprint,
		Variation:         fc.Variation,
		VariationOf:       fc.VariationOf,
		PriceUsd:          fc.Prices.Usd,
		PriceUsdFoil:      fc.Prices.UsdFoil,
		PriceUsdEtched:    fc.Prices.UsdEtched,
		PriceEur:          fc.Prices.Eur,
		PriceEurFoil:      fc.Prices.EurFoil,
		PriceTix:          fc.Prices.Tix,
		PrintedName:       fc.PrintedName,
		PrintedText:       fc.PrintedText,
		PrintedTypeLine:   fc.PrintedTypeLine,
		FlavorText:        fc.FlavorText,
		CardmarketId:      fc.CardmarketId,
		Uri:               fc.Uri,
		Id:                fc.Id,
		CardFaces:         fc.CardFaces,
		PngUri:            fc.ImageUris.Png,
		BoarderCropUri:    fc.ImageUris.BorderCrop,
		ArtCropUri:        fc.ImageUris.ArtCrop,
		LargeUri:          fc.ImageUris.Large,
		NormalUri:         fc.ImageUris.Normal,
		SmallUri:          fc.ImageUris.Small,
	}
	return print
}

func (s1 MtgSet) CompareSets(s2 MtgSet) bool {
	// First, check if they are the same type
	if reflect.TypeOf(s1) != reflect.TypeOf(s2) {
		return false
	}

	// Compare the fields one by one
	return s1.SetId == s2.SetId &&
		s1.SetCode == s2.SetCode &&
		s1.SetName == s2.SetName &&
		s1.SetType == s2.SetType &&
		s1.SetUri == s2.SetUri &&
		s1.SetSearchUri == s2.SetSearchUri &&
		s1.ScryfallSetUri == s2.ScryfallSetUri
}

func (p1 Prints) ComparePrints(p2 Prints) bool {
	if reflect.TypeOf(p1) != reflect.TypeOf(p2) {
		return false
	}

	return p1.MtgoId == p2.MtgoId &&
		p1.MtgoFoilId == p2.MtgoFoilId &&
		p1.ArenaId == p2.ArenaId &&
		p1.TcgplayerId == p2.TcgplayerId &&
		p1.TcgplayerEtchedId == p2.TcgplayerEtchedId &&
		p1.ReleasedAt == p2.ReleasedAt &&
		compareStringSlices(p1.Games, p2.Games) &&
		compareIntSlices(p1.AttractionLights, p2.AttractionLights) &&
		p1.Oversized == p2.Oversized &&
		p1.SetId == p2.SetId &&
		p1.OracleText == p2.OracleText &&
		p1.CollectorNumber == p2.CollectorNumber &&
		p1.Digital == p2.Digital &&
		p1.Rarity == p2.Rarity &&
		p1.OldschoolF == p2.OldschoolF &&
		p1.CardBackId == p2.CardBackId &&
		p1.Artist == p2.Artist &&
		p1.IllustrationId == p2.IllustrationId &&
		p1.BorderColor == p2.BorderColor &&
		p1.Frame == p2.Frame &&
		p1.FullArt == p2.FullArt &&
		p1.Textless == p2.Textless &&
		p1.Booster == p2.Booster &&
		p1.StorySpotlight == p1.StorySpotlight &&
		p1.GathererUri == p2.GathererUri &&
		p1.TcgArticlesUri == p2.TcgArticlesUri &&
		p1.TcgDecksUri == p2.TcgDecksUri &&
		p1.EdhrecUri == p2.EdhrecUri &&
		p1.TcgBuyUri == p2.TcgBuyUri &&
		p1.CardmarketBuyUri == p2.CardmarketBuyUri &&
		p1.CardhoarderBuyUri == p2.CardhoarderBuyUri &&
		p1.OracleId == p2.OracleId &&
		p1.CardName == p2.CardName &&
		p1.PrintsSearchUri == p2.PrintsSearchUri &&
		compareRelatedSlices(p1.Related, p2.Related) &&
		p1.FlavorName == p2.FlavorName &&
		compareStringSlices(p1.FrameEffects, p2.FrameEffects) &&
		p1.SecurityStamp == p2.SecurityStamp &&
		p1.PreviewedAt == p2.PreviewedAt &&
		p1.PreviewUri == p2.PreviewUri &&
		p1.PreviewSource == p2.PreviewSource &&
		p1.ContentWarning == p2.ContentWarning &&
		compareStringSlices(p1.BorderEffects, p2.BorderEffects) &&
		p1.Lang == p2.Lang &&
		compareIntSlices(p1.MultiverseIds, p2.MultiverseIds) &&
		p1.GathererUri == p2.GathererUri &&
		p1.ScryfallUri == p2.ScryfallUri &&
		p1.RulingsUri == p2.RulingsUri &&
		p1.HighresImage == p2.HighresImage &&
		p1.ImageStatus == p2.ImageStatus &&
		p1.Foil == p2.Foil &&
		p1.NotFoil == p2.NotFoil &&
		compareStringSlices(p1.Finishes, p2.Finishes) &&
		p1.Promo == p2.Promo &&
		compareStringSlices(p1.PromoTypes, p2.PromoTypes) &&
		p1.Reprint == p2.Reprint &&
		p1.Variation == p2.Variation &&
		p1.VariationOf == p2.VariationOf &&
		p1.PriceUsd == p2.PriceUsd &&
		p1.PriceUsdFoil == p2.PriceUsdFoil &&
		p1.PriceUsdEtched == p2.PriceUsdEtched &&
		p1.PriceEur == p2.PriceEur &&
		p1.PriceEurFoil == p2.PriceEurFoil &&
		p1.PriceTix == p2.PriceTix &&
		p1.PrintedName == p2.PrintedName &&
		p1.PrintedText == p2.PrintedText &&
		p1.PrintedTypeLine == p2.PrintedTypeLine &&
		p1.FlavorText == p2.FlavorText &&
		p1.CardmarketId == p2.CardmarketId &&
		p1.Uri == p2.Uri &&
		p1.Id == p2.Id &&
		compareCardFacesSlices(p1.CardFaces, p2.CardFaces) &&
		p1.PngUri == p2.PngUri &&
		p1.BoarderCropUri == p2.BoarderCropUri &&
		p1.ArtCropUri == p2.ArtCropUri &&
		p1.LargeUri == p2.LargeUri &&
		p1.NormalUri == p2.NormalUri &&
		p1.SmallUri == p2.SmallUri
}

func compareStringSlices(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Strings(s1)
	sort.Strings(s2)
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func compareIntSlices(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Ints(s1)
	sort.Ints(s2)
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func compareCardFacesSlices(s1 []CardFaces, s2 []CardFaces) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Slice(s1, func(i, j int) bool {
		return s1[i].Name < s1[j].Name
	})

	sort.Slice(s2, func(i, j int) bool {
		return s2[i].Name < s2[j].Name
	})
	for i := range s1 {
		if compareCardFaces(s1[i], s2[i]) {
			return false
		}
	}
	return true
}
func compareCardFaces(c1, c2 CardFaces) bool {
	return c1.Artist == c2.Artist &&
		c1.ArtistId == c2.ArtistId &&
		c1.Cmc == c2.Cmc &&
		compareStringSlices(c1.ColorIndicator, c2.ColorIndicator) &&
		compareStringSlices(c1.Colors, c2.Colors) &&
		c1.Defense == c2.Defense &&
		c1.FlavorText == c2.FlavorText &&
		c1.IllustrationId == c2.IllustrationId &&
		c1.PngUri == c2.PngUri &&
		c1.BoarderCropUri == c2.BoarderCropUri &&
		c1.ArtCropUri == c2.ArtCropUri &&
		c1.LargeUri == c2.LargeUri &&
		c1.NormalUri == c2.NormalUri &&
		c1.SmallUri == c2.SmallUri &&
		c1.Layout == c2.Layout &&
		c1.Loyalty == c2.Loyalty &&
		c1.ManaCost == c2.ManaCost &&
		c1.Name == c2.Name &&
		c1.Object == c2.Object &&
		c1.OracleId == c2.OracleId &&
		c1.OracleText == c2.OracleText &&
		c1.Power == c2.Power &&
		c1.PrintedName == c2.PrintedName &&
		c1.PrintedText == c2.PrintedText &&
		c1.PrintedTypeLine == c2.PrintedTypeLine &&
		c1.Toughness == c2.Toughness &&
		c1.TypeLine == c2.TypeLine &&
		c1.Watermark == c2.Watermark
}

func compareRelatedSlices(r1 []Related, r2 []Related) bool {
	if len(r1) != len(r2) {
		return false
	}
	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Id < r1[j].Id
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i].Id < r2[j].Id
	})
	for i := range r1 {
		if !compareRelated(r1[i], r2[i]) {
			return false
		}
	}
	return true
}

func compareRelated(r1, r2 Related) bool {
	return r1.Object == r2.Object &&
		r1.Id == r2.Id &&
		r1.Component == r2.Component &&
		r1.Name == r2.Name &&
		r1.TypeLine == r2.TypeLine &&
		r1.Uri == r2.Uri
}
