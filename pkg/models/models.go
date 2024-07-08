package models

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Sirname   string `json:"sirname"`
	Biography string `json:"biography"`
	Birthday  string `json:"birthday"`
}

type RelAuthorBook struct {
	Author Author `json:"author"`
	Book   Book   `json:"book"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id,omitempty"`
	Author   string `json:"author,omitempty"`
	Year     int    `json:"year"`
	ISBN     string `json:"isbn"`
}
