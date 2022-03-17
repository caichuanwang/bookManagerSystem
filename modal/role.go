package modal

import (
	"bookManagementSystem/api/feModal"
	"database/sql"
	"strconv"
)

type Role struct {
	Id          uint16        `json:"id,omitempty" form:"id"`
	Role_name   string        `json:"role_name,omitempty" form:"role_name"`
	Role_weight sql.NullInt16 `json:"role_weight,omitempty" form:"role_weight"`
}

func NewRoleWithAllKey() *Role {
	return &Role{Role_name: "", Role_weight: sql.NullInt16{}}
}

func NewRole() *Role {
	return &Role{}
}

func ToFeModal(r Role) *feModal.Role {
	var feR feModal.Role
	feR.Id = r.Id
	feR.Role_name = r.Role_name
	if r.Role_weight.Int16 == 0 {
		feR.Role_weight = "0"
	} else {
		feR.Role_weight = strconv.Itoa(int(r.Role_weight.Int16))
	}
	return &feR
}
