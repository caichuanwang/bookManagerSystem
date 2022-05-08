package modal

type BookBaseInfo struct {
	Isbn        string  `json:"isbn" form:"isbn" valid:"type(string),required"`
	BookName    string  `json:"bookName" form:"bookName" valid:"type(string),required"`
	Author      string  `json:"author" form:"author"`
	Publisher   string  `json:"publisher" form:"publisher"`
	PublishTime string  `json:"publishTime" form:"publishTime"`
	BookStock   uint    `json:"bookStock" form:"bookStock"`
	Price       float32 `json:"price" form:"price"`
	TypeId      uint    `json:"typeId" form:"typeId"`
	Context     string  `json:"context" form:"context"`
	PageNum     string  `json:"pageNum" form:"pageNum"`
	Translator  string  `json:"translator" form:"translator"`
}

type BookInfo struct {
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
