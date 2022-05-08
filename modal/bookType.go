package modal

type BookType struct {
	Id       uint   `json:"id,omitempty" form:"id"`
	TypeName string `json:"typeName" form:"typeName" valid:"type(string),required"`
	Level    string `json:"level" form:"level"`
	PId      uint   `json:"pId" form:"pId"`
	Remake   string `json:"remake" form:"remake"`
}

type BookTypeResult struct {
	Id       uint   `json:"id" `
	TypeName string `json:"typeName"`
	Level    string `json:"level"`
	PName    string `json:"pName"`
	PId      string `json:"pId"`
	Remake   string `json:"remake"`
}

type ItemBookType struct {
	Id       uint   `json:"id"`
	TypeName string `json:"name"`
	PId      uint   `json:"pId"`
	Remake   string `json:"remake"`
	Level    string `json:"level"`
}
type TreeBookType struct {
	ItemBookType
	Children []*TreeBookType `json:"children"`
}

type BookTypeReqParams struct {
	TypeName string `json:"typeName" form:"typeName" valid:"type(string),required"`
	Level    string `json:"level" form:"level"`
	PName    string `json:"pName"`
	Page
	OrderBy
	KeyWords
}
