package modal

type Role struct {
	Id          uint   `json:"id,omitempty" form:"id"`
	Role_name   string `json:"role_name,omitempty" form:"role_name"`
	Role_weight uint16 `json:"role_weight,omitempty" form:"role_weight"`
}

func NewRole() *Role {
	return &Role{}
}

type RoleListResult struct {
	Id          uint   `json:"id,omitempty" form:"id"`
	Role_name   string `json:"role_name,omitempty" form:"role_name"`
	Role_weight string `json:"role_weight,omitempty" form:"role_weight"`
}

type RoleParams struct {
	Role_name string `json:"role_name"`
	Page
}
