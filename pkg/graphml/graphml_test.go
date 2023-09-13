package graphml

import (
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestCreateNewGraphml(t *testing.T) {
	dir := micro.ProgramDir()
	DocXML, ElementGraph := CreateDocument()

	Group1 := CreateElement_Group(ElementGraph, "Group1", 100, 40)
	Entity1 := CreateElement_SmallEntity(Group1, "Entity1", 100, 0)
	Entity2 := CreateElement_SmallEntity(Group1, "Entity222", 100, 1)
	CreateElement_Edge(ElementGraph, Entity1, Entity2, "edge1", "descr")
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
