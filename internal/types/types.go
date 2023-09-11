package types

import (
	"github.com/ManyakRus/image_database/pkg/graphml"
)

type Column struct {
	Name        string `json:"name"   gorm:"column:name;default:''"`
	Type        string `json:"type_name"   gorm:"column:type_name;default:''"`
	Is_identity bool   `json:"is_identity"   gorm:"column:is_identity;default:false"`
	Description string `json:"description"   gorm:"column:description;default:''"`
	OrderNumber int
	TableKey    string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnKey   string `json:"column_key"   gorm:"column:column_key;default:''"`
}

type Table struct {
	Name string `json:"name"   gorm:"column:name;default:''"`
	//Element     *etree.Element
	ElementInfo graphml.ElementInfoStruct
	MapColumns  map[string]Column
	//Columns []Column
	OrderNumber int
}

//var MapTable = make(map[string]Table, 0)
