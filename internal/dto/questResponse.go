package dto

type QuestResponse struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Difficulty      string   `json:"difficulty"`
	DiffDescription string   `json:"diff_description"`
	Images          []string `json:"images"`
}

type CharacterResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ClassName   string   `json:"class"`
	RaceName    string   `json:"race_name"`
	Images      []string `json:"images"`
}
