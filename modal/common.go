package modal

type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Page struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type OrderBy struct {
	Order_by   string `json:"order_by"`
	Order_type string `json:"order_type"`
}
type KeyWords struct {
	KeyWord string `json:"keyWord" form:"keyWord" valid:"type(string)"`
}
