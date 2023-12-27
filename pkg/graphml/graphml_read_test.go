package graphml

import (
	"github.com/ManyakRus/image_database/internal/config"
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/micro"
	"testing"
)

func TestReadFile(t *testing.T) {
	dir := micro.ProgramDir()
	Filename := dir + "test" + micro.SeparatorFile() + "test.graphml"
	doc, err := ReadFile(Filename)
	if err != nil {
		t.Error("TestReadFile() error: ", err)
	}
	if doc == nil {
		t.Error("TestReadFile() error: doc =nil")
	}
}

func TestFindMassEntity(t *testing.T) {
	dir := micro.ProgramDir()
	Filename := dir + "test" + micro.SeparatorFile() + "test.graphml"
	doc, err := ReadFile(Filename)
	if err != nil {
		t.Error("TestFindMassEntity() error: ", err)
	}
	if doc == nil {
		t.Error("TestFindMassEntity() error: doc =nil")
	}

	Otvet := FindMassElement(doc)
	if len(Otvet) == 0 {
		t.Error("TestFindMassEntity() error: len =0")
	}

}

func TestFindMapNodeStruct(t *testing.T) {
	dir := micro.ProgramDir()
	Filename := dir + "test" + micro.SeparatorFile() + "test.graphml"
	doc, err := ReadFile(Filename)
	if err != nil {
		t.Error("TestFindMapNodeStruct() error: ", err)
	}
	if doc == nil {
		t.Error("TestFindMapNodeStruct() error: doc =nil")
	}

	MassElement := FindMassElement(doc)
	if len(MassElement) == 0 {
		t.Error("TestFindMapNodeStruct() error: len =0")
	}

	Otvet := FindMapNodeStruct(MassElement)
	if len(Otvet) == 0 {
		t.Error("TestFindMapNodeStruct() error: ", err)
	}

}

func TestStartReadFile(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()
	StartReadFile()
}
