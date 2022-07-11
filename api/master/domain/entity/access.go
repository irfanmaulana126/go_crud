package masterDomainEntity

type Access struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OrderItem   int    `json:"order"`
}

func (Access) TableName() string {
	return "wsapp_accesses"
}
