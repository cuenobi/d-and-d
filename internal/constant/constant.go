package constant

type (
	Role             uint
	Visibility       bool
	Race             string
	Class            string
	DifficultyLevels uint
	Params           string
	ImageStoragePath string
)

const (
	User  Role = 1
	Admin Role = 2
)

const (
	Private Visibility = true
	Public  Visibility = false
)

const (
	ImageFileName Params = "filename"
)

const (
	QuestPath     ImageStoragePath = "quests"
	CharacterPath ImageStoragePath = "characters"
)
