package models

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type BookRequest struct {
	Title    string `json:"title"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type BookUpdateRequest struct {
	Title    string `json:"title,omitempty"`
	Penulis  string `json:"penulis,omitempty"`
	Penerbit string `json:"penerbit,omitempty"`
}
