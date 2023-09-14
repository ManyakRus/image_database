package logic

import (
	"github.com/ManyakRus/image_database/internal/postgres"
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/ManyakRus/image_database/pkg/graphml"
	"github.com/ManyakRus/starter/log"
	"sort"
)

//var MassTable []types.Table

func StartFillAll(FileName string) bool {
	Otvet := false

	//заполним MapAll
	MapAll, err := postgres.FillMapTable()
	if err != nil {
		log.Error("FillMapTable() error: ", err)
		return Otvet
	}

	if len(MapAll) > 0 {
		Otvet = true
	}

	if Otvet == false {
		println("warning: Empty file not saved !")
		return Otvet
	}

	//создадим документ
	DocXML, ElementInfoGraph := graphml.CreateDocument()

	//заполним прямоугольники в документ
	err = FillEntities(ElementInfoGraph, &MapAll)
	if err != nil {
		log.Error("FillEntities() error: ", err)
		return Otvet
	}

	//заполним стрелки в документ
	err = FillEdges(ElementInfoGraph, &MapAll)
	if err != nil {
		log.Error("FillEdges() error: ", err)
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

// FillEntities - заполняет прямоугольники Entities в файл .xml
func FillEntities(ElementInfoGraph graphml.ElementInfoStruct, MapAll *map[string]*types.Table) error {
	var err error

	for _, table1 := range *MapAll {
		TextAttributes := ""
		MassColumns := MassFromMapColumns(table1.MapColumns)
		for _, column1 := range MassColumns {
			if TextAttributes != "" {
				TextAttributes = TextAttributes + "\n"
			}
			TextAttributes = TextAttributes + column1.Name + "  " + column1.Type
		}
		ElementInfo1 := graphml.CreateElement_Entity(ElementInfoGraph, table1.Name, TextAttributes)
		table1.ElementInfo = ElementInfo1
	}

	return err
}

// MassFromMapColumns - возвращает Slice из Map
func MassFromMapColumns(MapColumns map[string]types.Column) []types.Column {
	Otvet := make([]types.Column, 0)

	for _, v := range MapColumns {
		Otvet = append(Otvet, v)
	}

	sort.Slice(Otvet[:], func(i, j int) bool {
		return Otvet[i].OrderNumber < Otvet[j].OrderNumber
	})

	return Otvet
}

// FillEdges - заполняет стрелки в файл .xml
func FillEdges(ElementInfoGraph graphml.ElementInfoStruct, MapAll *map[string]*types.Table) error {
	var err error

	MapAll0 := *MapAll

	for _, table1 := range *MapAll {
		for _, column1 := range table1.MapColumns {
			if column1.TableKey == "" || column1.ColumnKey == "" {
				continue
			}
			//только если есть внешний ключ
			//тыблица из ключа
			TableKey, ok := MapAll0[column1.TableKey]
			if ok == false {
				log.Warn("Error. Not found table name: ", column1.TableKey)
				continue
			}

			//колонка из ключа
			ColumnKey, ok := TableKey.MapColumns[column1.ColumnKey]
			if ok == false {
				log.Warn("Error. Not found column name: ", column1.ColumnKey)
				continue
			}

			//
			decription := column1.Name + " - " + ColumnKey.Name
			graphml.CreateElement_Edge(ElementInfoGraph, table1.ElementInfo, TableKey.ElementInfo, "", decription, column1.OrderNumber+1, ColumnKey.OrderNumber+1)
		}
	}

	return err
}
