package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) NewBook(id int, title, author string, year, size int, rate float32) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

func (b Book) Book() interface{} {
	return b
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b Book) ID() (id int) {
	return b.id
}

func (b Book) Title() (title string) {
	return b.title
}

func (b Book) Author() (author string) {
	return b.author
}

func (b Book) Year() (year int) {
	return b.year
}

func (b Book) Size() (size int) {
	return b.size
}

func (b Book) Rate() (rate float32) {
	return b.rate
}
