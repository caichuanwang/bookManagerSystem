package feModal

type Role struct {
	Id          uint16 `json:"id,omitempty" form:"id"`
	Role_name   string `json:"role_name,omitempty" form:"role_name"`
	Role_weight string `json:"role_weight,omitempty" form:"role_weight"`
}
