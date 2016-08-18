package domain

type UrlTupleType struct {
	Link  string `json:"url"`
	Title string `json:"title"`
}

type UrlTupleArrayType struct {
	URLs []UrlTupleType `json:"links"`
}
