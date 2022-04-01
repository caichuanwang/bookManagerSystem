package feModal

type User struct {
	Id                int64  `json:"id,omitempty" form:"id"`
	User_name         string `json:"user_name,omitempty" form:"user_name"`
	Sex               int8   `json:"sex,omitempty" form:"sex"`
	Birthday          string `json:"birthday" form:"birthday"`
	Borrow_book_count int    `json:"borrow_book_count,omitempty" form:"borrow_book_count"`
	Phone             string `json:"phone,omitempty" form:"phone"`
	Remake            string `json:"remake,omitempty" form:"remake"`
	Email             string `json:"email,omitempty" form:"email"`
	Role              string `json:"role" form:"role"`
}
type QueryUserParams struct {
	User_name         string `json:"user_name" form:"user_name"`
	Borrow_book_count string `json:"borrow_book_count" form:"borrow_book_count"`
	Role              string `json:"role" form:"role"`
	Page
	OrderBy
}

type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type OrderBy struct {
	Order_by   string `json:"order_by"`
	Order_type string `json:"order_type"`
}
