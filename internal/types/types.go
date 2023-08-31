package types

type Column struct {
	Name        string `json:"name"   gorm:"column:name;default:''"`
	Type        string `json:"type_name"   gorm:"column:type_name;default:''"`
	Is_identity bool   `json:"is_identity"   gorm:"column:is_identity;default:false"`
	Description string `json:"description"   gorm:"column:description;default:''"`
}

type Table struct {
	Name    string `json:"name"   gorm:"column:name;default:''"`
	Columns []Column
}
