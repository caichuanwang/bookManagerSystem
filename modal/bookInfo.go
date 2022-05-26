package modal

type BookBaseInfo struct {
	Isbn        string     `json:"isbn" form:"isbn" valid:"type(string),required" gorm:"primary_key;comment:图书唯一编号"`
	BookName    string     `json:"bookName" form:"bookName" valid:"type(string),required" gorm:"not null;column:bookName"`
	Author      string     `json:"author" form:"author"`
	Publisher   string     `json:"publisher" form:"publisher"`
	PublishTime string     `json:"publishTime" form:"publishTime" gorm:"column:publishTime"`
	BookStock   uint       `json:"bookStock" form:"bookStock" gorm:"column:bookStock"`
	Price       float32    `json:"price" form:"price" gorm:"type:decimal(8,2)"`
	TypeId      uint       `json:"typeId" form:"typeId" gorm:"type:integer;column:typeId"`
	Context     string     `json:"context" form:"context" gorm:"type:text"`
	PageNum     string     `json:"pageNum" form:"pageNum" gorm:"column:pageNum"`
	Translator  string     `json:"translator" form:"translator"`
	BookList    []BookList `json:"bookList" gorm:"many2many:book_list_map"`
}

type BookInfoReturn struct {
	BookBaseInfo
	TypeName string `json:"typeName"`
	Photo    string `json:"photo" form:"photo"`
}

type CreateBookInfoParams struct {
	BookBaseInfo
	Photo []map[string]any `json:"photo" form:"photo"`
}

type QueryBookInfoParams struct {
	BookName string `json:"bookName" form:"bookName" valid:"type(string),required"`
	Page
	OrderBy
}
type BookBorrowTopRes struct {
	BookInfoReturn
	Score uint `json:"score"`
}

type BookInfo struct {
	BookBaseInfo
	Photo string `json:"photo" form:"photo"`
}
