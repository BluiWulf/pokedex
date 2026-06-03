package pokeapi

type SpriteInfo struct {
	Gen4Base
	Other					OtherSet	`json:"other"`
	Versions				GenVersions	`json:"versions"`
}

type OtherSet struct {
	DreamWorld				Gen7Base	`json:"dream_world"`
	Home					Gen6Base	`json:"home"`
	OfficialArtwork			Gen3Base	`json:"official_artwork"`
	Showdown				Gen4Base	`json:"showdown"`
}

type GenVersions struct {
	GenerationI				Gen1		`json:"generation-i"`
	GenerationII			Gen2		`json:"generation-ii"`
	GenerationIII			Gen3		`json:"generation-iii"`
	GenerationIV			Gen4		`json:"generation-iv"`
	GenerationV				Gen5		`json:"generation-v"`
	GenerationVI			Gen6		`json:"generation-vi"`
	GenerationVII			Gen7		`json:"generation-vii"`
	GenerationVIII			Gen8		`json:"generation-viii"`
}

type Gen1 struct {
	RedBlue					Gen1Base	`json:"red-blue"`
	Yellow					Gen1Base	`json:"yellow"`
}

type Gen2 struct {
	Crystal					Gen2Base	`json:"crystal"`
	Gold					Gen2Base	`json:"gold"`
	Silver					Gen2Base	`json:"silver"`
}

type Gen3 struct {
	Emerald					Gen3Base	`json:"emerald"`
	FireRedLeafGreen		Gen2Base	`json:"firered-leafgreen"`
	RubySapphire			Gen2Base	`json:"ruby-sapphire"`
}

type Gen4 struct {
	DiamondPearl			Gen4Base	`json:"diamond-pearl"`
	HeartGoldSoulSilver		Gen4Base	`json:"heartgold-soulsilver"`
	Platinum				Gen4Base	`json:"platinum"`
}

type Gen5 struct {
	BlackWhite				Gen5Base	`json:"black-white"`
}

type Gen6 struct {
	OmegaRubyAlphaSapphire	Gen6Base	`json:"omegaruby-alphasapphire"`
	XY						Gen6Base	`json:"x-y"`
}

type Gen7 struct {
	Icons					Gen7Base	`json:"icons"`
	UltraSunUltraMoon		Gen6Base	`json:"ultra-sun-ultra-moon"`
}

type Gen8 struct {
	Icons					Gen7Base	`json:"icons"`
}

type Gen1Base struct {
	BackDefault				*string		`json:"back_default"`
	BackGray				*string		`json:"back_gray"`
	FrontDefault			*string		`json:"front_default"`
	FrontGray				*string		`json:"front_gray"`
}

type Gen2Base struct {
	BackDefault				*string		`json:"back_default"`
	BackShiny				*string		`json:"back_shiny"`
	FrontDefault			*string		`json:"front_default"`
	FrontShiny				*string		`json:"front_shiny"`
}

type Gen3Base struct {
	FrontDefault			*string		`json:"front_default"`
	FrontShiny				*string		`json:"front_shiny"`
}

type Gen4Base struct {
	BackDefault				*string		`json:"back_default"`
	BackFemale				*string		`json:"back_female"`
	BackShiny				*string		`json:"back_shiny"`
	BackShinyFemale			*string		`json:"back_shiny_female"`
	FrontDefault			*string		`json:"front_default"`
	FrontFemale				*string		`json:"front_female"`
	FrontShiny				*string		`json:"front_shiny"`
	FrontShinyFemale		*string		`json:"front_shiny_female"`
}

type Gen5Base struct {
	Animated				Gen4Base	`json:"animated"`
	Gen4Base
}

type Gen6Base struct {
	FrontDefault			*string		`json:"front_default"`
	FrontFemale				*string		`json:"front_female"`
	FrontShiny				*string		`json:"front_shiny"`
	FrontShinyFemale		*string		`json:"front_shiny_female"`
}

type Gen7Base struct {
	FrontDefault			*string		`json:"front_default"`
	FrontFemale				*string		`json:"front_female"`
}
