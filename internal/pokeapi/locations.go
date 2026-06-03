package pokeapi

type RespAreas struct {
	Count	   			 int	      		   `json:"count"`
	Next	   			 *string    		   `json:"next"`
	Previous   			 *string    		   `json:"previous"`
	Results	   			 []BasicData 		   `json:"results"`
}

type LocationArea struct {
	Id		   			 int		  		   `json:"id"`
	Name	   			 string	  			   `json:"name"`
	GameIndex  			 int		  		   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location			 BasicData			   `json:"location"`
	Names				 []Name				   `json:"names"`
	PokemonEncounters	 []PokemonEncounter	   `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod		 BasicData			   `json:"encounter_method"`
	VersionDetails		 []EncounterVersion	   `json:"version_details"`
}

type EncounterVersion struct {
	Rate 				 int				   `json:"rate"`
	Version 			 BasicData			   `json:"version"`
}

type Name struct {
	Name				 string				   `json:"name"`
	Language			 BasicData			   `json:"language"`
}

type PokemonEncounter struct {
	Pokemon				 BasicData			   `json:"pokemon"`
	VersionDetails		 []PokemonVersion	   `json:"version_details"`
}

type PokemonVersion struct {
	Version				 BasicData			   `json:"version"`
	MaxChance			 int				   `json:"max_chance"`
	EncounterDetails	 []EncounterDetail	   `json:"encounter_details"`
}

type EncounterDetail struct {
	MinLvl				 int				   `json:"min_level"`
	MaxLvl				 int				   `json:"max_level"`
	ConditionVals		 []BasicData		   `json:"condition_values"`
	Chance 				 int				   `json:"chance"`
	Method				 BasicData			   `json:"method"`
}