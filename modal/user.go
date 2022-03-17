package modal

type User struct {
	Id                int64  `json:"id,omitempty" form:"id"`
	User_name         string `json:"user_name,omitempty" form:"user_name"`
	User_password     string `json:"user_password,omitempty" form:"user_password"`
	Sex               int8   `json:"sex,omitempty" form:"sex"`
	Birthday          string `json:"birthday" form:"birthday"`
	Borrow_book_count int    `json:"borrow_book_count,omitempty" form:"borrow_book_count"`
	Phone             string `json:"phone,omitempty" form:"phone"`
	Remake            string `json:"remake,omitempty" form:"remake"`
	Email             string `json:"email,omitempty" form:"email"`
	Role              int    `json:"role" form:"role"`
}

func NewUser() *User {
	return &User{}
}

func NewUserWithAllKeys() *User {
	return &User{
		Id:                0,
		User_name:         "",
		User_password:     "",
		Sex:               0,
		Birthday:          "",
		Borrow_book_count: 0,
		Phone:             "",
		Remake:            "",
		Email:             "",
		Role:              0,
	}
}
