package domain

type Object struct {
	ID         string `json:"id" gorm:"type:uuid;primary_key"`
	Name       string `json:"name"`
	RegisterID string `json:"register_id"`
}

func (Object) TableName() string {
	return "object"
}
