package graphml

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestCreateNewGraphml(t *testing.T) {
	dir := micro.ProgramDir()
	DocXML, ElementGraph := CreateDocument()

	Entity1 := CreateElement_Entity(ElementGraph, "Entity1", "Field1\nField2\nField3\n1234567890")
	Entity2 := CreateElement_Entity(ElementGraph, "Entity2", "Field1\nField2\nField3\n1234567890")
	CreateElement_Edge(ElementGraph, Entity1, Entity2, "edge1", "descr", 1, 4)
	//Shape2 := CreateElement_Shape(ElementGraph, "Shape2")
	//Group1 := CreateElement_Group(ElementGraph, "Group1")
	//Shape1 := CreateElement_Shape(Group1, "Shape1")
	//CreateElement_Edge(ElementGraph, Shape1, Shape2, "edge1", "descr")
	//CreateElement_Edge_blue(ElementGraph, Shape2, Shape1, "edge2", "descr2")
	//
	//if Shape1 == nil || Shape2 == nil {
	//
	//}

	FileName := dir + "test" + micro.SeparatorFile() + "test.graphml"
	//DocXML.IndentTabs()
	DocXML.Indent(2)
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		t.Error("TestCreateNewXGML() error: ", err)
	}
}
