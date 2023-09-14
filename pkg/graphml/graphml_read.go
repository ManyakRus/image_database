package graphml

import (
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/beevik/etree"
	"strconv"
)

func ReadFile(Filename string) (*etree.Document, error) {
	var err error
	doc := etree.NewDocument()
	err = doc.ReadFromFile(Filename)
	if err != nil {
		log.Panic(err)
	}

	return doc, err
}

func FindMassElement(doc *etree.Document) []*etree.Element {
	MassElement := make([]*etree.Element, 0)

	ElementGraphMl := doc.SelectElement("graphml")
	ElementGraph := ElementGraphMl.SelectElement("graph")

	MassElement = ElementGraph.SelectElements("node")

	return MassElement
}

func FindMapNodeStruct(MassElement []*etree.Element) map[string]types.NodeStruct {
	MapNodeStruct := make(map[string]types.NodeStruct, 0)
	var err error

	for _, ElementNode1 := range MassElement {
		sx := ""
		sy := ""
		Name := ""
		MassData := ElementNode1.SelectElements("data")
		if len(MassData) == 0 {
			continue
		}
		var ElementData1 *etree.Element
		ElementData1 = MassData[len(MassData)-1]
		//for _, ElementData1 = range MassData {
		//key := ElementData1.SelectAttrValue("key", "")
		//if key != "d5" {
		//	continue
		//}

		ElementGenericNode := ElementData1.SelectElement("y:GenericNode")
		if ElementGenericNode == nil {
			continue
		}

		ElementGeometry := ElementGenericNode.SelectElement("y:Geometry")
		if ElementGeometry == nil {
			continue
		}

		sx = ElementGeometry.SelectAttrValue("x", "0")
		sy = ElementGeometry.SelectAttrValue("y", "0")

		ElementNodeLabel := ElementGenericNode.SelectElement("y:NodeLabel")
		if ElementNodeLabel == nil {
			continue
		}

		Name = ElementNodeLabel.Text()
		if Name == "" {
			log.Warn("Name = ''")
			continue
		}

		var x float64
		if sx != "" {
			x, err = strconv.ParseFloat(sx, 32)
			if err != nil {
				log.Warn("Name: ", Name+" ParseFloat(", sx, ") error: ", err)
			}
		}

		var y float64
		if sy != "" {
			y, err = strconv.ParseFloat(sy, 32)
			if err != nil {
				log.Warn("Name: ", Name+" ParseFloat(", sy, ") error: ", err)
			}
		}

		NodeStruct1 := types.NodeStruct{}
		NodeStruct1.Element = ElementNode1
		NodeStruct1.Name = Name
		NodeStruct1.X = x
		NodeStruct1.Y = y

		MapNodeStruct[Name] = NodeStruct1
		//}
	}

	return MapNodeStruct
}

// StartReadFile - читает старый файл в
func StartReadFile() {
	//dir := micro.ProgramDir()
	//Filename := dir + "test" + micro.SeparatorFile() + "test.graphml"
	Filename := config.Settings.FILENAME_GRAPHML

	ok, err := micro.FileExists(Filename)
	if ok == false {
		return
	}

	doc, err := ReadFile(Filename)
	if err != nil {
		log.Error("ReadFile() error: ", err)
		return
	}
	if doc == nil {
		log.Error("ReadFile() error: doc =nil")
		return
	}

	MassElement := FindMassElement(doc)
	if len(MassElement) == 0 {
		log.Warn("FindMassElement() error: len =0")
		return
	}

	types.MapNodeStructOld = FindMapNodeStruct(MassElement)
	if len(types.MapNodeStructOld) == 0 {
		log.Warn("FindMapNodeStruct() error: len =0")
		return
	}

}
