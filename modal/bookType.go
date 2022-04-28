package modal

type BookType struct {
	Id       uint   `json:"id,omitempty" form:"id"`
	TypeName string `json:"typeName" form:"typeName" valid:"type(string),required"`
	Path     string `json:"path" form:"path" valid:"type(string),required"`
	Level    string `json:"level" form:"level"`
	Branch   string `json:"branch" form:"branch"`
	Remake   string `json:"remake" form:"remake"`
}

func NewBookType() *BookType {
	return &BookType{}
}

type BookTypeReqParams struct {
	Page
	OrderBy
	KeyWords
}
