package books

type book struct {
	Name string
}

func New() book {
	return book{Name: "Harry"}
}
