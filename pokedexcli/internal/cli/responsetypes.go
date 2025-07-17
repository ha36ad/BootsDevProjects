package cli

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type exploreAreaResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type pokemonResponse struct {
	ID                     int                    `json:"id"`
	Name                   string                 `json:"name"`
	BaseExperience         int                    `json:"base_experience"`
	Height                 int                    `json:"height"`
	Weight                 int                    `json:"weight"`
	Types                  []pokemonTypeSlot      `json:"types"`
	Stats                  []pokemonStat          `json:"stats"`
	Abilities              []pokemonAbility       `json:"abilities"`
	Forms                  []namedAPIResource     `json:"forms"`
	GameIndices            []gameIndex            `json:"game_indices"`
	HeldItems              []heldItem             `json:"held_items"`
	LocationAreaEncounters string                 `json:"location_area_encounters"`
	Moves                  []pokemonMove          `json:"moves"`
	Species                namedAPIResource       `json:"species"`
	Sprites                map[string]interface{} `json:"sprites"`
	Cries                  map[string]string      `json:"cries"`
	PastTypes              []pastType             `json:"past_types"`
	PastAbilities          []pastAbility          `json:"past_abilities"`
}

type pokemonTypeSlot struct {
	Slot int              `json:"slot"`
	Type namedAPIResource `json:"type"`
}

type pokemonStat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     namedAPIResource `json:"stat"`
}

type pokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  namedAPIResource `json:"ability"`
}

type gameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   namedAPIResource `json:"version"`
}

type heldItem struct {
	Item           namedAPIResource  `json:"item"`
	VersionDetails []heldItemVersion `json:"version_details"`
}

type heldItemVersion struct {
	Rarity  int              `json:"rarity"`
	Version namedAPIResource `json:"version"`
}

type pokemonMove struct {
	Move                namedAPIResource    `json:"move"`
	VersionGroupDetails []moveVersionDetail `json:"version_group_details"`
}

type moveVersionDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    namedAPIResource `json:"version_group"`
	MoveLearnMethod namedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type pastType struct {
	Generation namedAPIResource  `json:"generation"`
	Types      []pokemonTypeSlot `json:"types"`
}

type pastAbility struct {
	Generation namedAPIResource    `json:"generation"`
	Abilities  []pastAbilityDetail `json:"abilities"`
}

type pastAbilityDetail struct {
	Ability  *namedAPIResource `json:"ability"`
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}

type namedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
