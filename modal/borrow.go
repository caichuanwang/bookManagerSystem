package modal

type Borrow struct {
	Id                 uint64 `json:"id" form:"id"`
	Borrow_reader_id   uint   `json:"borrow_reader_id" form:"borrow_reader_id"`
	Borrow_book_isbn   string `json:"borrow_book_isbn" form:"borrow_book_isbn"`
	Borrow_time        string `json:"borrow_time" form:"borrow_time"`
	Should_return_time string `json:"should_return_time" form:"should_return_time"`
	Really_return_time string `json:"really_return_time" form:"really_return_time"`
	Agree_borrow_time  string `json:"agree_borrow_time" from:"agree_borrow_time"`
	Is_borrow          uint   `json:"is_borrow" form:"is_borrow"`
	Is_return          uint   `json:"is_return" form:"is_return"`
}

type QueryBorrowParams struct {
	BorrowReaderAndBookName
	Page
	OrderBy
}
type BorrowWithName struct {
	BorrowReaderAndBookName
	Borrow
}

type BorrowReaderAndBookName struct {
	Borrow_reader_name string `json:"borrow_reader_name" form:"borrow_reader_name"`
	Borrow_book_name   string `json:"borrow_book_name" form:"borrow_book_name"`
}

const (
	NO_RETURN = iota + 1
	RETURN
)

const (
	NO_BORROW = iota + 1
	BORROW
)

const BOOK_BORROW_TOP_KEY_REDIS = "book_borrow_top"
