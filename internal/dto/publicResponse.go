package dto

type PublicQuestResponse struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Difficulty      string   `json:"difficulty"`
	DiffDescription string   `json:"diff_description"`
	Images          []string `json:"images"`
}

type PublicCharacterResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ClassName   string   `json:"class"`
	RaceName    string   `json:"race_name"`
	Images      []string `json:"images"`
}
