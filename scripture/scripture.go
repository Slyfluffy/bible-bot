package scripture

type Scripture struct {
	Reference string `json:"reference"`
	Verses    []struct {
		BookID   string `json:"book_id"`
		BookName string `json:"book_name"`
		Chapter  int    `json:"chapter"`
		Verse    int    `json:"verse"`
		Text     string `json:"text"`
	} `json:"verses"`
	Text            string `json:"text"`
	TranslationID   string `json:"translation_id"`
	TranslationName string `json:"translation_name"`
	TranslationNote string `json:"translation_note"`
}