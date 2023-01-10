package models

type CategoryBookPrimaryKey struct {
	Id string `json:"id"`
}

type BookCategory struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Categorys   []Categories
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Categories struct {
	Id   string `json:"id"`
	Name string `json:"Category"`
}

type CategoryBook struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	BookInfos []Books
}
type Books struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
