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
