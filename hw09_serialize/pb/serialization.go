package pb

import "google.golang.org/protobuf/proto"

func SerializeBooks(books []*Book) ([]byte, error) {
	bookList := &BookList{
		Books: books,
	}
	return proto.Marshal(bookList)
}

func DeserializeBooks(data []byte) ([]*Book, error) {
	bookList := &BookList{}
	if err := proto.Unmarshal(data, bookList); err != nil {
		return nil, err
	}
	return bookList.Books, nil
}
