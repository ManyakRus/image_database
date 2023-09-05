package logic

import (
	"github.com/ManyakRus/image_database/internal/postgres"
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/ManyakRus/image_database/pkg/graphml"
	"github.com/ManyakRus/starter/log"
)

//var MassTable []types.Table

func StartFillAll(FileName string) bool {
	Otvet := false

	TableAll, err := postgres.FillMassTable()
	if err != nil {
		log.Error("FillMassTable() error: ", err)
		return Otvet
	}

	if len(TableAll) > 0 {
		Otvet = true
	}

	if Otvet == false {
		println("warning: Empty file not saved !")
		return Otvet
	}

	DocXML, ElementInfoGraph := graphml.CreateDocument()

	err = FillEntities(ElementInfoGraph, TableAll)
	if err != nil {
		log.Error("FillEntities() error: ", err)
		return Otvet
	}

	log.Info("Start save file")
	DocXML.Indent(2)
	err = DocXML.WriteToFile(FileName)
	if err != nil {
		log.Error("WriteToFile() FileName: ", FileName, " error: ", err)
	}

	return Otvet
}

func FillEntities(ElementInfoGraph graphml.ElementInfoStruct, TableAll []types.Table) error {
	var err error

	for _, table1 := range TableAll {
		TextAttributes := ""
		for _, column1 := range table1.Columns {
			if TextAttributes != "" {
				TextAttributes = TextAttributes + "\n"
			}
			TextAttributes = TextAttributes + column1.Name + "  " + column1.Type
		}
		graphml.CreateElement_Entity(ElementInfoGraph, table1.Name, TextAttributes)
	}

	return err
}
