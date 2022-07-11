package masterDomainEntity

type Roles struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Roles) TableName() string {
	return "wsapp_user_roles"
}
