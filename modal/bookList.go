package modal

import (
	"time"
)

type BookList struct {
	ID       uint       `json:"id" form:"id" gorm:"primary_key;autoIncrement"`
	Name     string     `json:"name" form:"name" gorm:"not null;unique" valid:"type(string),required"`
	UserId   uint       `json:"userId" form:"userId" gorm:"not null"`
	Photo    string     `json:"photo" form:"photo"`
	Remake   string     `json:"remake" form:"remake"`
	Time     time.Time  `json:"time" form:"remake"`
	BookInfo []BookInfo `json:"bookInfo" gorm:"many2many:book_list_map;References:Isbn"`
	User     []User     `json:"user" gorm:"many2many:book_user_map;References:Id"`
}
type SetBook2BookListParams struct {
	BookLists []uint `json:"bookLists" from:"bookLists"`
	Isbn      string `json:"isbn" form:"isbn"`
}

type QueryBookListParams struct {
	Name string `json:"name"`
	Page
}
