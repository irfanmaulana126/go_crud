package masterDomainEntity

type Permissions struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionsReq struct {
	RoleCode string `json:"role_code"`
}

func (Permissions) TableName() string {
	return "wsapp_accesses"
}
