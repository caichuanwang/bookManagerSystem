package modal

type Role struct {
	Id          uint   `json:"id,omitempty" form:"id"`
	Role_name   string `json:"role_name,omitempty" form:"role_name"`
	Role_weight uint16 `json:"role_weight,omitempty" form:"role_weight"`
}

func NewRoleWithAllKey() *Role {
	return &Role{Role_name: "", Role_weight: 1}
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

//func ToFeModal(r Role) *feModal.Role {
//	var feR feModal.Role
//	feR.Id = r.Id
//	feR.Role_name = r.Role_name
//	if r.Role_weight == 0 {
//		feR.Role_weight = "0"
//	} else {
//		feR.Role_weight = strconv.Itoa(int(r.Role_weight))
//	}
//	return &feR
//}
