package modal

type User struct {
	Id                int64      `json:"id" form:"id"  gorm:"primary_key;autoIncrement"`
	User_name         string     `json:"user_name" form:"user_name" valid:"type(string),required"`
	User_password     string     `json:"user_password" form:"user_password" valid:"type(string),required"`
	Sex               int8       `json:"sex,omitempty" form:"sex" valid:"type(int)"`
	Birthday          string     `json:"birthday" form:"birthday"`
	Borrow_book_count int        `json:"borrow_book_count,omitempty" form:"borrow_book_count" valid:"type(int),range(0|5)" gorm:"type:integer" `
	Phone             string     `json:"phone,omitempty" form:"phone"`
	Remake            string     `json:"remake,omitempty" form:"remake"`
	Email             string     `json:"email,omitempty" form:"email" valid:"email"`
	Role              string     `json:"role" form:"role" gorm:"type:integer"`
	BookList          []BookList `json:"bookList" gorm:"many2many:book_user_map" `
}

func NewUser() *User {
	return &User{}
}

type UserListResult struct {
	Id                int64  `json:"id,omitempty" form:"id"`
	User_name         string `json:"user_name,omitempty" form:"user_name" valid:"string"`
	Sex               int8   `json:"sex,omitempty" form:"sex"`
	Birthday          string `json:"birthday" form:"birthday"`
	Borrow_book_count int    `json:"borrow_book_count,omitempty" form:"borrow_book_count"`
	Phone             string `json:"phone,omitempty" form:"phone"`
	Remake            string `json:"remake,omitempty" form:"remake"`
	Email             string `json:"email,omitempty" form:"email"`
	Role              string `json:"role" form:"role"`
	RoleName          string `json:"roleName" from:"roleName"`
}
type QueryUserParams struct {
	User_name         string `json:"user_name" form:"user_name"`
	Borrow_book_count string `json:"borrow_book_count" form:"borrow_book_count"`
	Role              string `json:"role" form:"role"`
	Page
	OrderBy
}
