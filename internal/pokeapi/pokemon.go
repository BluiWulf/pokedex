package pokeapi

type PokemonInfo struct {
	Id					 int				   `json:"id"`
	Name				 string				   `json:"name"`
	BaseXP				 int				   `json:"base_experience"`
	Height				 int				   `json:"height"`
	IsDefault			 bool				   `json:"is_default"`
	Order				 int				   `json:"order"`
	Weight				 int				   `json:"weight"`
	Abilities			 []AbilityInfo		   `json:"abilities"`
	Forms				 []BasicData		   `json:"forms"`
	GameIndices			 []GameIndexInfo	   `json:"game_indices"`
	HeldItems			 []HeldItem			   `json:"held_items"`
	LocationEncounters	 string				   `json:"location_area_encounters"`
	Moves				 []MoveInfo			   `json:"moves"`
	Species				 BasicData			   `json:"species"`
	Sprites				 SpriteInfo			   `json:"sprites"`
	Cries				 CryInfo			   `json:"cries"`
	Stats				 []StatInfo			   `json:"stats"`
	Types				 []TypeInfo			   `json:"types"`
	PastTypes			 []PastType			   `json:"past_types"`
	PastAbilities		 []PastAbility		   `json:"past_abilities"`
}

type PastAbility struct {
	Generation			 BasicData			   `json:"generation"`
	Abilities			 []AbilityInfo		   `json:"abilities"`
}

type AbilityInfo struct {
	IsHidden			 bool				   `json:"is_hidden"`
	Slot				 int				   `json:"slot"`
	Ability				 BasicData			   `json:"ability"`
}

type GameIndexInfo struct {
	GameIndex			 int				   `json:"game_index"`
	Version				 BasicData			   `json:"version"`
}

type HeldItem struct {
	Item				 BasicData			   `json:"item"`
	VersionDetails		 []ItemVersion		   `json:"version_details"`
}

type ItemVersion struct {
	Rarity				 int				   `json:"rarity"`
	Version				 BasicData			   `json:"version"`
}

type MoveInfo struct {
	Move				 BasicData			   `json:"move"`
	VersionGroupDetails	 []VersionGroupDetail  `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LvlLearnedAt		 int				   `json:"level_learned_at"`
	VersionGroup		 BasicData			   `json:"version_group"`
	MoveLearnMethod		 BasicData			   `json:"move_learn_method"`
	Order				 int				   `json:"order"`
}

type CryInfo struct {
	Latest				 string				   `json:"latest"`
	Legacy				 string				   `json:"legacy"`
}

type StatInfo struct {
	BaseStat			 int				   `json:"base_stat"`
	Effort				 int				   `json:"effort"`
	Stat				 BasicData			   `json:"stat"`
}

type PastType struct {
	Generation			 BasicData			   `json:"generation"`
	Types				 []TypeInfo			   `json:"types"`
}

type TypeInfo struct {
	Slot				 int				   `json:"slot"`
	Type				 BasicData			   `json:"type"`
}
