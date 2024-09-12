package types

import (
	"github.com/beevik/etree"
)

type Column struct {
	Name         string `json:"name"   gorm:"column:name;default:''"`
	Type         string `json:"type_name"   gorm:"column:type_name;default:''"`
	IsPrimaryKey bool   `json:"is_primary_key"   gorm:"column:is_primary_key;default:false"`
	Description  string `json:"description"   gorm:"column:description;default:''"`
	OrderNumber  int
	TableKey     string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnKey    string `json:"column_key"   gorm:"column:column_key;default:''"`
}

type Table struct {
	Name    string `json:"name"   gorm:"column:name;default:''"`
	Comment string
	//Element     *etree.Element
	ElementInfo ElementInfoStruct
	MapColumns  map[string]Column
	//Columns []Column
	OrderNumber int
}

type NodeStruct struct {
	Element *etree.Element
	Name    string
	X       float64
	Y       float64
}

type ElementInfoStruct struct {
	Element     *etree.Element
	Name        string
	Attribute   string
	Description string
	Width       float64
	Height      float64
	Parent      *ElementInfoStruct
}

var MapNodeStructOld = make(map[string]NodeStruct, 0)
