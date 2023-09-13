package graphml

import (
	"fmt"
	"github.com/beevik/etree"
	_ "github.com/beevik/etree"
	"math"
	"strconv"
	"strings"
)

// FONT_SIZE_ENTITY - размер шрифта Entity
var FONT_SIZE_ENTITY = 12

// FONT_SIZE_BENDS - размер шрифта стрелки куриная лапка
var FONT_SIZE_BENDS = 8

// FONT_SIZE_SHAPE - размер шрифта прямоугольника
var FONT_SIZE_SHAPE = 12

// FONT_SIZE_SHAPE - размер шрифта групп
var FONT_SIZE_GROUP = 18

// FONT_SIZE_EDGE - размер шрифта стрелок
var FONT_SIZE_EDGE = 8

type ElementInfoStruct struct {
	Element     *etree.Element
	Name        string
	Attribute   string
	Description string
	Width       float64
	Height      float64
	Parent      *ElementInfoStruct
}

// CreateElement_Entity - создаёт элемент - Entity
func CreateElement_Entity(ElementInfoMain ElementInfoStruct, ElementName, ElementAttribute string) ElementInfoStruct {

	Width := findWidth_Entity(ElementName + "\n" + ElementAttribute)
	Height := findHeight_Entity(ElementName + ElementAttribute)
	sWidth := fmt.Sprintf("%.1f", float64(Width))
	sHeight := fmt.Sprintf("%.1f", float64(Height))

	sFontSize := strconv.Itoa(FONT_SIZE_ENTITY)

	//ищем graph
	var ElementGraph *etree.Element
	ElementGraph2 := ElementInfoMain.Element.SelectElement("graph")
	if ElementGraph2 != nil {
		ElementGraph = ElementGraph2
	} else {
		ElementGraph = ElementInfoMain.Element
	}

	//node
	ElementNode := ElementGraph.CreateElement("node")

	var ElementInfoNode ElementInfoStruct
	ElementInfoNode.Element = ElementNode
	ElementInfoNode.Name = ElementName
	ElementInfoNode.Parent = &ElementInfoMain
	ElementInfoNode.Attribute = ElementAttribute
	ElementInfoNode.Width = Width
	ElementInfoNode.Height = Height

	sId := FindId(ElementInfoMain, ElementInfoNode)
	ElementNode.CreateAttr("id", sId)
	//ElementNode.CreateAttr("id", "n"+strconv.Itoa(ElementNode.Index()))

	//data
	ElementData := ElementNode.CreateElement("data")
	ElementData.CreateAttr("key", "d4")

	//data
	ElementData2 := ElementNode.CreateElement("data")
	ElementData2.CreateAttr("key", "d5")

	//y:GenericNode
	ElementYGenericNode := ElementData2.CreateElement("y:GenericNode")
	ElementYGenericNode.CreateAttr("configuration", "com.yworks.entityRelationship.big_entity")

	//YGeometry
	ElementYGeometry := ElementYGenericNode.CreateElement("y:Geometry")
	ElementYGeometry.CreateAttr("height", sHeight)
	ElementYGeometry.CreateAttr("width", sWidth)
	ElementYGeometry.CreateAttr("x", "-270.0")
	ElementYGeometry.CreateAttr("y", "-65.0")

	//YFill
	ElementYFill := ElementYGenericNode.CreateElement("y:Fill")
	ElementYFill.CreateAttr("color", "#E8EEF7")
	ElementYFill.CreateAttr("color2", "#B7C9E3")
	ElementYFill.CreateAttr("transparent", "false")

	//BorderStyle
	ElementBorderStyle := ElementYGenericNode.CreateElement("y:BorderStyle")
	ElementBorderStyle.CreateAttr("color", "#000000")
	ElementBorderStyle.CreateAttr("type", "line")
	ElementBorderStyle.CreateAttr("width", "1.0")

	//NodeLabel
	ElementNodeLabel := ElementYGenericNode.CreateElement("y:NodeLabel")
	ElementNodeLabel.CreateAttr("alignment", "center")
	ElementNodeLabel.CreateAttr("autoSizePolicy", "content")
	ElementNodeLabel.CreateAttr("backgroundColor", "#B7C9E3")
	ElementNodeLabel.CreateAttr("configuration", "com.yworks.entityRelationship.label.name")
	ElementNodeLabel.CreateAttr("fontFamily", "Dialog")
	ElementNodeLabel.CreateAttr("fontSize", sFontSize)
	ElementNodeLabel.CreateAttr("fontStyle", "plain")
	ElementNodeLabel.CreateAttr("hasLineColor", "false")
	ElementNodeLabel.CreateAttr("height", sHeight)
	ElementNodeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementNodeLabel.CreateAttr("iconTextGap", "4")
	ElementNodeLabel.CreateAttr("modelName", "internal")
	ElementNodeLabel.CreateAttr("modelPosition", "t")
	ElementNodeLabel.CreateAttr("textColor", "#000000")
	ElementNodeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementNodeLabel.CreateAttr("visible", "true")
	ElementNodeLabel.CreateAttr("width", sWidth)
	ElementNodeLabel.CreateAttr("x", "16.0")
	ElementNodeLabel.CreateAttr("xml:space", "preserve")
	ElementNodeLabel.CreateAttr("y", "4.0")
	ElementNodeLabel.CreateText(ElementName)

	//NodeLabel
	ElementNodeLabel2 := ElementYGenericNode.CreateElement("y:NodeLabel")
	ElementNodeLabel2.CreateAttr("alignment", "left")
	ElementNodeLabel2.CreateAttr("autoSizePolicy", "content")
	ElementNodeLabel2.CreateAttr("configuration", "com.yworks.entityRelationship.label.attributes")
	ElementNodeLabel2.CreateAttr("fontFamily", "Dialog")
	ElementNodeLabel2.CreateAttr("fontSize", sFontSize)
	ElementNodeLabel2.CreateAttr("fontStyle", "plain")
	ElementNodeLabel2.CreateAttr("hasBackgroundColor", "false")
	ElementNodeLabel2.CreateAttr("hasLineColor", "false")
	ElementNodeLabel2.CreateAttr("height", sHeight)
	ElementNodeLabel2.CreateAttr("horizontalTextPosition", "center")
	ElementNodeLabel2.CreateAttr("iconTextGap", "4")
	ElementNodeLabel2.CreateAttr("modelName", "free")
	ElementNodeLabel2.CreateAttr("modelPosition", "anywhere")
	ElementNodeLabel2.CreateAttr("textColor", "#000000")
	ElementNodeLabel2.CreateAttr("verticalTextPosition", "top")
	ElementNodeLabel2.CreateAttr("visible", "true")
	ElementNodeLabel2.CreateAttr("width", sWidth)
	ElementNodeLabel2.CreateAttr("x", "2.0")
	ElementNodeLabel2.CreateAttr("xml:space", "preserve")
	ElementNodeLabel2.CreateAttr("y", "30.0")
	ElementNodeLabel2.CreateText(ElementAttribute)

	//y:LabelModel
	ElementYLabelModel := ElementNodeLabel2.CreateElement("y:LabelModel")

	//y:ErdAttributesNodeLabelModel
	ElementYLabelModel.CreateElement("y:ErdAttributesNodeLabelModel")

	//y:ModelParameter
	ElementYModelParameter := ElementNodeLabel2.CreateElement("y:ModelParameter")

	//y:ErdAttributesNodeLabelModelParameter
	ElementYModelParameter.CreateElement("y:ErdAttributesNodeLabelModelParameter")

	//y:StyleProperties
	ElementYStyleProperties := ElementYGenericNode.CreateElement("y:StyleProperties")

	//y:Property
	ElementYProperty := ElementYStyleProperties.CreateElement("y:Property")
	ElementYProperty.CreateAttr("class", "java.lang.Boolean")
	ElementYProperty.CreateAttr("name", "y.view.ShadowNodePainter.SHADOW_PAINTING")
	ElementYProperty.CreateAttr("value", "true")

	return ElementInfoNode
}

// CreateElement_Edge - создаёт элемент graphml - стрелка
func CreateElement_Edge(ElementInfoGraph, ElementInfoFrom, ElementInfoTo ElementInfoStruct, label, Description string) ElementInfoStruct {

	//
	sx := float32(-ElementInfoFrom.Width / 2)
	sy := float32(-ElementInfoFrom.Height/2) + 3 + float32(FONT_SIZE_ENTITY)/2
	tx := float32(-ElementInfoTo.Width / 2)
	ty := float32(-ElementInfoTo.Height/2) + 3 + float32(FONT_SIZE_ENTITY)/2

	TextSx := fmt.Sprintf("%.2f", sx)
	TextSy := fmt.Sprintf("%.2f", sy)
	TextTx := fmt.Sprintf("%.2f", tx)
	TextTy := fmt.Sprintf("%.2f", ty)

	//node
	ElementEdge := ElementInfoGraph.Element.CreateElement("edge")

	var ElementInfoEdge ElementInfoStruct
	ElementInfoEdge.Element = ElementEdge
	ElementInfoEdge.Parent = &ElementInfoGraph
	ElementInfoEdge.Name = label
	ElementInfoEdge.Description = Description

	//EdgeId := FindId(ElementInfoGraph, ElementEdge)
	//EdgeID := EdgeId
	EdgeID := "e" + strconv.Itoa(ElementEdge.Index())
	ElementEdge.CreateAttr("id", EdgeID)
	//Source := "n" + strconv.Itoa(IndexElementFrom) + "::" + "n" + strconv.Itoa(IndexElementTo)
	IdFrom := FindId(ElementInfoGraph, ElementInfoFrom)
	IdTo := FindId(ElementInfoGraph, ElementInfoTo)
	ElementEdge.CreateAttr("source", IdFrom)
	ElementEdge.CreateAttr("target", IdTo)

	//data
	ElementData := ElementEdge.CreateElement("data")
	ElementData.CreateAttr("key", "d8")
	ElementData.CreateAttr("xml:space", "preserve")
	//ElementInfoStruct.CreateText("<![CDATA[descr]]>")
	//ElementInfoStruct.CreateElement("![CDATA[descr]]")
	ElementData.CreateCData(Description)

	//data2
	ElementData2 := ElementEdge.CreateElement("data")
	ElementData2.CreateAttr("key", "d9")

	//y:PolyLineEdge
	ElementYPolyLineEdge := ElementData2.CreateElement("y:PolyLineEdge")

	//y:Path
	ElementYPath := ElementYPolyLineEdge.CreateElement("y:Path")
	ElementYPath.CreateAttr("sx", TextSx)
	ElementYPath.CreateAttr("sy", TextSy)
	ElementYPath.CreateAttr("tx", TextTx)
	ElementYPath.CreateAttr("ty", TextTy)

	//y:LineStyle
	ElementYLineStyle := ElementYPolyLineEdge.CreateElement("y:LineStyle")
	ElementYLineStyle.CreateAttr("color", "#000000")
	ElementYLineStyle.CreateAttr("type", "line")
	ElementYLineStyle.CreateAttr("width", "1.0")

	//y:Arrows
	ElementYArrows := ElementYPolyLineEdge.CreateElement("y:Arrows")
	ElementYArrows.CreateAttr("source", "crows_foot_many")
	ElementYArrows.CreateAttr("target", "none")

	//y:EdgeLabel
	ElementYEdgeLabel := ElementYPolyLineEdge.CreateElement("y:EdgeLabel")
	ElementYEdgeLabel.CreateAttr("alignment", "center")
	ElementYEdgeLabel.CreateAttr("configuration", "AutoFlippingLabel")
	ElementYEdgeLabel.CreateAttr("distance", "0.0")
	ElementYEdgeLabel.CreateAttr("fontFamily", "Dialog")
	ElementYEdgeLabel.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_EDGE))
	ElementYEdgeLabel.CreateAttr("fontStyle", "plain")
	ElementYEdgeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementYEdgeLabel.CreateAttr("hasLineColor", "false")
	ElementYEdgeLabel.CreateAttr("height", "17.96875")
	ElementYEdgeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementYEdgeLabel.CreateAttr("iconTextGap", "4")
	ElementYEdgeLabel.CreateAttr("modelName", "centered")
	ElementYEdgeLabel.CreateAttr("modelPosition", "head")
	ElementYEdgeLabel.CreateAttr("preferredPlacement", "anywhere")
	//ElementYEdgeLabel.CreateAttr("modelName", "two_pos")
	//ElementYEdgeLabel.CreateAttr("modelPosition", "head")
	//ElementYEdgeLabel.CreateAttr("preferredPlacement", "on_edge")
	ElementYEdgeLabel.CreateAttr("ratio", "0.5")
	ElementYEdgeLabel.CreateAttr("textColor", "#000000")
	ElementYEdgeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementYEdgeLabel.CreateAttr("visible", "true")
	ElementYEdgeLabel.CreateAttr("width", "41.8")
	ElementYEdgeLabel.CreateAttr("x", "71.5")
	ElementYEdgeLabel.CreateAttr("xml:space", "preserve")
	ElementYEdgeLabel.CreateAttr("y", "0.5")
	ElementYEdgeLabel.CreateAttr("bottomInset", "0")
	ElementYEdgeLabel.CreateAttr("leftInset", "0")
	ElementYEdgeLabel.CreateAttr("rightInset", "0")
	ElementYEdgeLabel.CreateAttr("topInset", "0")
	ElementYEdgeLabel.CreateText(label)

	//y:PreferredPlacementDescriptor
	ElementYPreferredPlacementDescriptor := ElementYEdgeLabel.CreateElement("y:PreferredPlacementDescriptor")
	ElementYPreferredPlacementDescriptor.CreateAttr("angle", "0.0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleOffsetOnRightSide", "0")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleReference", "absolute")
	ElementYPreferredPlacementDescriptor.CreateAttr("angleRotationOnRightSide", "co")
	ElementYPreferredPlacementDescriptor.CreateAttr("distance", "-1.0")
	//ElementYPreferredPlacementDescriptor.CreateAttr("frozen", "true")
	ElementYPreferredPlacementDescriptor.CreateAttr("placement", "anywhere")
	ElementYPreferredPlacementDescriptor.CreateAttr("side", "on_edge")
	ElementYPreferredPlacementDescriptor.CreateAttr("sideReference", "relative_to_edge_flow")

	//y:BendStyle
	ElementYBendStyle := ElementYPolyLineEdge.CreateElement("y:BendStyle")
	ElementYBendStyle.CreateAttr("smoothed", "false")

	return ElementInfoEdge
}

// CreateElement_Group - создаёт элемент xgml - группа
func CreateElement_Group(ElementInfoGraph ElementInfoStruct, GroupCaption string, Width, Height float64) ElementInfoStruct {

	//Width := FindWidth_Group(GroupCaption)
	//Height := FindHeight_Group(GroupCaption)
	sWidth := fmt.Sprintf("%.1f", float32(Width))
	sHeight := fmt.Sprintf("%.1f", float32(Height))
	sWidth = "0.0"
	sHeight = "0.0" //авторазмер

	//ищем graph
	var ElementGraph *etree.Element
	ElementGraph2 := ElementInfoGraph.Element.SelectElement("graph")
	if ElementGraph2 != nil {
		ElementGraph = ElementGraph2
	} else {
		ElementGraph = ElementInfoGraph.Element
	}

	//node
	ElementNode := ElementGraph.CreateElement("node")

	var ElementInfoGroup ElementInfoStruct
	ElementInfoGroup.Element = ElementNode
	ElementInfoGroup.Parent = &ElementInfoGraph
	ElementInfoGroup.Name = GroupCaption
	ElementInfoGroup.Description = ""

	//NodeId := "n" + strconv.Itoa(ElementNode.Index())
	NodeId := FindId(ElementInfoGraph, ElementInfoGroup)
	ElementNode.CreateAttr("id", NodeId)
	ElementNode.CreateAttr("yfiles.foldertype", "group")

	//data
	ElementData := ElementNode.CreateElement("data")
	ElementData.CreateAttr("key", "d5")

	//YProxyAutoBoundsNode
	ElementYProxyAutoBoundsNode := ElementData.CreateElement("y:ProxyAutoBoundsNode")

	//YRealizers
	ElementYRealizers := ElementYProxyAutoBoundsNode.CreateElement("y:Realizers")
	ElementYRealizers.CreateAttr("active", "0")

	//----------------------- visible ---------------------------------------------

	//YGroupNode
	ElementYGroupNode := ElementYRealizers.CreateElement("y:GroupNode")

	//YGeometry
	ElementYGeometry := ElementYGroupNode.CreateElement("y:Geometry")
	ElementYGeometry.CreateAttr("height", sHeight)
	ElementYGeometry.CreateAttr("width", sWidth)
	ElementYGeometry.CreateAttr("x", "0.0")
	ElementYGeometry.CreateAttr("y", "0.0")

	//YFill
	ElementYFill := ElementYGroupNode.CreateElement("y:Fill")
	ElementYFill.CreateAttr("color", "#E8EEF7")
	ElementYFill.CreateAttr("color2", "#B7C9E3")
	ElementYFill.CreateAttr("transparent", "false")

	//YBorderStyle
	ElementYBorderStyle := ElementYGroupNode.CreateElement("y:BorderStyle")
	ElementYBorderStyle.CreateAttr("color", "#F5F5F5")
	ElementYBorderStyle.CreateAttr("type", "dashed")
	ElementYBorderStyle.CreateAttr("width", "1.0")

	//YNodeLabel
	ElementYNodeLabel := ElementYGroupNode.CreateElement("y:NodeLabel")
	ElementYNodeLabel.CreateAttr("alignment", "right")
	ElementYNodeLabel.CreateAttr("autoSizePolicy", "content")
	//ElementYNodeLabel.CreateAttr("backgroundColor", "#EBEBEB")
	ElementYNodeLabel.CreateAttr("borderDistance", "0.0")
	ElementYNodeLabel.CreateAttr("fontFamily", "Dialog")
	ElementYNodeLabel.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_GROUP))
	ElementYNodeLabel.CreateAttr("fontStyle", "bold")
	ElementYNodeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementYNodeLabel.CreateAttr("hasLineColor", "false")
	ElementYNodeLabel.CreateAttr("height", sHeight)
	ElementYNodeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementYNodeLabel.CreateAttr("iconTextGap", "4")
	ElementYNodeLabel.CreateAttr("modelName", "internal")
	ElementYNodeLabel.CreateAttr("modelPosition", "t")
	ElementYNodeLabel.CreateAttr("textColor", "#000000")
	ElementYNodeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementYNodeLabel.CreateAttr("width", sWidth)
	ElementYNodeLabel.CreateAttr("x", "0")
	ElementYNodeLabel.CreateAttr("xml:space", "preserve")
	ElementYNodeLabel.CreateAttr("y", "0")
	ElementYNodeLabel.CreateText(GroupCaption)

	//YShape
	ElementYShape := ElementYGroupNode.CreateElement("y:Shape")
	ElementYShape.CreateAttr("type", "rectangle")

	//YState
	ElementYState := ElementYGroupNode.CreateElement("y:State")
	ElementYState.CreateAttr("closed", "false")
	ElementYState.CreateAttr("closedHeight", "80.0")
	ElementYState.CreateAttr("closedWidth", "100.0")
	ElementYState.CreateAttr("innerGraphDisplayEnabled", "false")

	//YInsets
	ElementYInsets := ElementYGroupNode.CreateElement("y:Insets")
	ElementYInsets.CreateAttr("bottom", "0")
	ElementYInsets.CreateAttr("bottomF", "0.0")
	ElementYInsets.CreateAttr("left", "0")
	ElementYInsets.CreateAttr("leftF", "0.0")
	ElementYInsets.CreateAttr("right", "0")
	ElementYInsets.CreateAttr("rightF", "0.0")
	ElementYInsets.CreateAttr("top", "0")
	ElementYInsets.CreateAttr("topF", "0.0")

	//YBorderInsets
	ElementYBorderInsets := ElementYGroupNode.CreateElement("y:BorderInsets")
	ElementYBorderInsets.CreateAttr("bottom", "54")
	ElementYBorderInsets.CreateAttr("bottomF", "54.0")
	ElementYBorderInsets.CreateAttr("left", "0")
	ElementYBorderInsets.CreateAttr("leftF", "0.0")
	ElementYBorderInsets.CreateAttr("right", "23")
	ElementYBorderInsets.CreateAttr("rightF", "23.35")
	ElementYBorderInsets.CreateAttr("top", "0")
	ElementYBorderInsets.CreateAttr("topF", "0.0")

	//----------------------- not visible ---------------------------------------------

	//YGroupNode
	ElementYGroupNode2 := ElementYRealizers.CreateElement("y:GroupNode")

	//YGeometry
	ElementYGeometry2 := ElementYGroupNode2.CreateElement("y:Geometry")
	ElementYGeometry2.CreateAttr("height", "40.0")
	ElementYGeometry2.CreateAttr("width", sWidth)
	ElementYGeometry2.CreateAttr("x", "0.0")
	ElementYGeometry2.CreateAttr("y", "0.0")

	//YFill
	ElementYFill2 := ElementYGroupNode2.CreateElement("y:Fill")
	ElementYFill2.CreateAttr("color", "#E8EEF7")
	ElementYFill2.CreateAttr("color2", "#B7C9E3")
	ElementYFill2.CreateAttr("transparent", "false")

	//YBorderStyle
	ElementYBorderStyle2 := ElementYGroupNode2.CreateElement("y:BorderStyle")
	ElementYBorderStyle2.CreateAttr("color", "#000000")
	ElementYBorderStyle2.CreateAttr("type", "dashed")
	ElementYBorderStyle2.CreateAttr("width", "1.0")

	//YNodeLabel
	ElementYNodeLabel2 := ElementYGroupNode2.CreateElement("y:NodeLabel")
	ElementYNodeLabel2.CreateAttr("alignment", "right")
	ElementYNodeLabel2.CreateAttr("autoSizePolicy", "content")
	//ElementYNodeLabel2.CreateAttr("backgroundColor", "#EBEBEB")
	ElementYNodeLabel2.CreateAttr("borderDistance", "0.0")
	ElementYNodeLabel2.CreateAttr("fontFamily", "Dialog")
	ElementYNodeLabel2.CreateAttr("fontSize", strconv.Itoa(FONT_SIZE_GROUP))
	ElementYNodeLabel2.CreateAttr("fontStyle", "bold")
	ElementYNodeLabel.CreateAttr("hasBackgroundColor", "false")
	ElementYNodeLabel2.CreateAttr("hasLineColor", "false")
	ElementYNodeLabel2.CreateAttr("hasText", "true") //только у 2
	ElementYNodeLabel2.CreateAttr("height", sHeight)
	ElementYNodeLabel2.CreateAttr("horizontalTextPosition", "center")
	ElementYNodeLabel2.CreateAttr("iconTextGap", "4")
	ElementYNodeLabel2.CreateAttr("modelName", "internal")
	ElementYNodeLabel2.CreateAttr("modelPosition", "t")
	ElementYNodeLabel2.CreateAttr("textColor", "#000000")
	ElementYNodeLabel2.CreateAttr("verticalTextPosition", "bottom")
	ElementYNodeLabel2.CreateAttr("width", sWidth)
	ElementYNodeLabel2.CreateAttr("x", "0")
	ElementYNodeLabel2.CreateAttr("xml:space", "preserve") //только у 2
	ElementYNodeLabel2.CreateAttr("y", "0")
	ElementYNodeLabel2.CreateText(GroupCaption) //только у 2

	//YShape
	ElementYShape2 := ElementYGroupNode2.CreateElement("y:Shape")
	ElementYShape2.CreateAttr("type", "roundrectangle")

	//YState
	ElementYState2 := ElementYGroupNode2.CreateElement("y:State")
	ElementYState2.CreateAttr("closed", "true")
	ElementYState2.CreateAttr("closedHeight", "80.0")
	ElementYState2.CreateAttr("closedWidth", "100.0")
	ElementYState2.CreateAttr("innerGraphDisplayEnabled", "false")

	//YInsets
	ElementYInsets2 := ElementYGroupNode2.CreateElement("y:Insets")
	ElementYInsets2.CreateAttr("bottom", "15")
	ElementYInsets2.CreateAttr("bottomF", "15.0")
	ElementYInsets2.CreateAttr("left", "15")
	ElementYInsets2.CreateAttr("leftF", "15.0")
	ElementYInsets2.CreateAttr("right", "15")
	ElementYInsets2.CreateAttr("rightF", "15.0")
	ElementYInsets2.CreateAttr("top", "15")
	ElementYInsets2.CreateAttr("topF", "15.0")

	//YBorderInsets
	ElementYBorderInsets2 := ElementYGroupNode2.CreateElement("y:BorderInsets")
	ElementYBorderInsets2.CreateAttr("bottom", "54")
	ElementYBorderInsets2.CreateAttr("bottomF", "54.0")
	ElementYBorderInsets2.CreateAttr("left", "0")
	ElementYBorderInsets2.CreateAttr("leftF", "0.0")
	ElementYBorderInsets2.CreateAttr("right", "23")
	ElementYBorderInsets2.CreateAttr("rightF", "23.35")
	ElementYBorderInsets2.CreateAttr("top", "0")
	ElementYBorderInsets2.CreateAttr("topF", "0.0")

	//----------------------- продолжение ---------------------------------------------
	//YBorderInsets
	ElementGraphGraph := ElementNode.CreateElement("graph")
	ElementGraphGraph.CreateAttr("edgedefault", "directed")
	ElementGraphGraph.CreateAttr("id", NodeId+":")

	return ElementInfoGroup
}

// CreateElement_SmallEntity - создаёт элемент - Entity
func CreateElement_SmallEntity(ElementInfoMain ElementInfoStruct, ElementName string, Width float64, AttributeIndex int) ElementInfoStruct {

	//Width := findWidth_SmallEntity(ElementName)
	Height := findHeight_SmallEntity(ElementName)
	sWidth := fmt.Sprintf("%.1f", float64(Width))
	sHeight := fmt.Sprintf("%.1f", float64(Height))
	sY := fmt.Sprintf("%.1f", float64(AttributeIndex)*Height)

	sFontSize := strconv.Itoa(FONT_SIZE_ENTITY)

	//ищем graph
	var ElementGraph *etree.Element
	ElementGraph2 := ElementInfoMain.Element.SelectElement("graph")
	if ElementGraph2 != nil {
		ElementGraph = ElementGraph2
	} else {
		ElementGraph = ElementInfoMain.Element
	}

	//node
	ElementNode := ElementGraph.CreateElement("node")

	var ElementInfoNode ElementInfoStruct
	ElementInfoNode.Element = ElementNode
	ElementInfoNode.Name = ElementName
	ElementInfoNode.Parent = &ElementInfoMain
	ElementInfoNode.Attribute = ""
	ElementInfoNode.Width = Width
	ElementInfoNode.Height = Height

	sId := FindId(ElementInfoMain, ElementInfoNode)
	ElementNode.CreateAttr("id", sId)
	//ElementNode.CreateAttr("id", "n"+strconv.Itoa(ElementNode.Index()))

	//data
	ElementData := ElementNode.CreateElement("data")
	ElementData.CreateAttr("key", "d4")

	//data
	ElementData2 := ElementNode.CreateElement("data")
	ElementData2.CreateAttr("key", "d5")

	//y:GenericNode
	ElementYGenericNode := ElementData2.CreateElement("y:GenericNode")
	ElementYGenericNode.CreateAttr("configuration", "com.yworks.entityRelationship.small_entity")

	//YGeometry
	ElementYGeometry := ElementYGenericNode.CreateElement("y:Geometry")
	ElementYGeometry.CreateAttr("height", sHeight)
	ElementYGeometry.CreateAttr("width", sWidth)
	ElementYGeometry.CreateAttr("x", "0")
	ElementYGeometry.CreateAttr("y", sY)

	//YFill
	ElementYFill := ElementYGenericNode.CreateElement("y:Fill")
	ElementYFill.CreateAttr("color", "#E8EEF7")
	ElementYFill.CreateAttr("color2", "#B7C9E3")
	ElementYFill.CreateAttr("transparent", "false")

	//BorderStyle
	ElementBorderStyle := ElementYGenericNode.CreateElement("y:BorderStyle")
	ElementBorderStyle.CreateAttr("hasColor", "false")
	//ElementBorderStyle.CreateAttr("color", "#000000")
	ElementBorderStyle.CreateAttr("type", "line")
	ElementBorderStyle.CreateAttr("width", "1.0")

	//NodeLabel
	ElementNodeLabel := ElementYGenericNode.CreateElement("y:NodeLabel")
	ElementNodeLabel.CreateAttr("alignment", "left")
	ElementNodeLabel.CreateAttr("autoSizePolicy", "content")
	ElementNodeLabel.CreateAttr("backgroundColor", "#B7C9E3")
	ElementNodeLabel.CreateAttr("borderDistance", "0.0")
	ElementNodeLabel.CreateAttr("configuration", "com.yworks.entityRelationship.label.name")
	ElementNodeLabel.CreateAttr("fontFamily", "Dialog")
	ElementNodeLabel.CreateAttr("fontSize", sFontSize)
	ElementNodeLabel.CreateAttr("fontStyle", "plain")
	ElementNodeLabel.CreateAttr("hasLineColor", "false")
	ElementNodeLabel.CreateAttr("height", sHeight)
	ElementNodeLabel.CreateAttr("horizontalTextPosition", "center")
	ElementNodeLabel.CreateAttr("iconTextGap", "4")
	ElementNodeLabel.CreateAttr("modelName", "internal")
	ElementNodeLabel.CreateAttr("modelPosition", "tl")
	ElementNodeLabel.CreateAttr("textColor", "#000000")
	ElementNodeLabel.CreateAttr("verticalTextPosition", "bottom")
	ElementNodeLabel.CreateAttr("visible", "true")
	ElementNodeLabel.CreateAttr("width", sWidth)
	ElementNodeLabel.CreateAttr("x", "16.0")
	ElementNodeLabel.CreateAttr("xml:space", "preserve")
	ElementNodeLabel.CreateAttr("y", "4.0")
	ElementNodeLabel.CreateText(ElementName)

	//y:LabelModel
	ElementYLabelModel := ElementNodeLabel.CreateElement("y:LabelModel")

	//y:SmartNodeLabelModel
	ElementYSmartNodeLabelModel := ElementYLabelModel.CreateElement("y:SmartNodeLabelModel")
	ElementYSmartNodeLabelModel.CreateAttr("distance", "0.0")

	////y:ErdAttributesNodeLabelModel
	//ElementYLabelModel.CreateElement("y:ErdAttributesNodeLabelModel")

	////y:ModelParameter
	ElementYModelParameter := ElementNodeLabel.CreateElement("y:ModelParameter")

	//y:SmartNodeLabelModelParameter
	ElementYSmartNodeLabelModelParameter := ElementYModelParameter.CreateElement("y:SmartNodeLabelModelParameter")
	ElementYSmartNodeLabelModelParameter.CreateAttr("labelRatioX", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("labelRatioY", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("nodeRatioX", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("nodeRatioY", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("offsetX", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("offsetY", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("upX", "0.0")
	ElementYSmartNodeLabelModelParameter.CreateAttr("upY", "-1.0")

	//y:StyleProperties
	ElementYStyleProperties := ElementYGenericNode.CreateElement("y:StyleProperties")

	//y:Property
	ElementYProperty := ElementYStyleProperties.CreateElement("y:Property")
	ElementYProperty.CreateAttr("class", "java.lang.Boolean")
	ElementYProperty.CreateAttr("name", "y.view.ShadowNodePainter.SHADOW_PAINTING")
	ElementYProperty.CreateAttr("value", "true")

	return ElementInfoNode
}

// findWidth_Entity - возвращает число - ширину элемента
func findWidth_Entity(ElementName string) float64 {
	Otvet := float64(FONT_SIZE_ENTITY) * 2

	LenMax := findMaxLenRow(ElementName)
	//var OtvetF float64
	Otvet = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE)*float64(0.48)
	//Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Entity - возвращает число - высоту элемента
func findHeight_Entity(ElementName string) float64 {

	var Otvet float64

	Otvet = float64(12 + FONT_SIZE_ENTITY*3)

	RowsTotal := countLines(ElementName)

	Otvet = float64(Otvet) + (float64(RowsTotal-1) * math.Round(float64(FONT_SIZE_ENTITY)*float64(1.16)))

	return Otvet

}

// findWidth_SmallEntity - возвращает число - ширину элемента
func findWidth_SmallEntity(ElementName string) int {
	Otvet := FONT_SIZE_ENTITY * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE)*float64(0.48)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_SmallEntity - возвращает число - высоту элемента
func findHeight_SmallEntity(ElementName string) float64 {

	var Otvet float64

	Otvet = float64(6 + FONT_SIZE_ENTITY)

	RowsTotal := countLines(ElementName)

	Otvet = float64(Otvet) + (float64(RowsTotal-1) * float64(FONT_SIZE_ENTITY))

	return Otvet

}

// findWidth_Bends - возвращает число - ширину элемента
func findWidth_Bends(ElementName string) int {
	Otvet := FONT_SIZE_BENDS * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE/2)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Bends - возвращает число - высоту элемента
func findHeight_Bends(ElementName string) int {

	Otvet := 10 + FONT_SIZE_BENDS*3

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (RowsTotal-1)*FONT_SIZE_SHAPE*2

	return Otvet

}

// findWidth_Shape - возвращает число - ширину элемента
func findWidth_Shape(ElementName string) int {
	Otvet := FONT_SIZE_SHAPE * 2

	LenMax := findMaxLenRow(ElementName)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*float64(FONT_SIZE_SHAPE/2)
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Shape - возвращает число - высоту элемента
func findHeight_Shape(ElementName string) float64 {

	Otvet := 10 + float64(FONT_SIZE_SHAPE)*3

	RowsTotal := countLines(ElementName)

	Otvet = Otvet + (float64(RowsTotal)-1)*float64(FONT_SIZE_SHAPE)*2

	return Otvet

}

// FindWidth_Group - возвращает число - ширину элемента
func FindWidth_Group(ElementName string) float64 {
	var Otvet float64 = 10

	LenMax := findMaxLenRow(ElementName)
	//var OtvetF float64
	Otvet = float64(Otvet) + float64(LenMax)*10
	//Otvet = int(math.Round(OtvetF))

	return Otvet
}

// FindHeight_Group - возвращает число - высоту элемента
func FindHeight_Group(ElementName string) float64 {

	var Otvet float64 = 0

	//RowsTotal := countLines(ElementName)
	HeightSmallEntity := findHeight_SmallEntity(ElementName)

	//Otvet = Otvet + float64(RowsTotal-1)*18
	Otvet = Otvet + HeightSmallEntity

	return Otvet

}

// findWidth_Edge - возвращает число - ширину элемента
func findWidth_Edge(Label string) int {
	Otvet := 10

	LenMax := findMaxLenRow(Label)
	var OtvetF float64
	OtvetF = float64(Otvet) + float64(LenMax)*10
	Otvet = int(math.Round(OtvetF))

	return Otvet
}

// findHeight_Edge - возвращает число - высоту элемента
func findHeight_Edge(Label string) int {

	Otvet := 30

	RowsTotal := countLines(Label)

	Otvet = Otvet + (RowsTotal-1)*18

	return Otvet

}

// countLines - возвращает количество переводов строки
func countLines(s string) int {
	Otvet := 1

	Otvet2 := strings.Count(s, "\n")
	Otvet = Otvet + Otvet2

	return Otvet
}

// findMaxLenRow - возвращает количество символов в строке максимум
func findMaxLenRow(ElementName string) int {
	Otvet := 0

	Mass := strings.Split(ElementName, "\n")

	for _, Mass1 := range Mass {
		MassRune := []rune(Mass1)
		len1 := len(MassRune)
		if len1 > Otvet {
			Otvet = len1
		}
	}

	return Otvet
}

// CreateDocument - создаёт новый документ .xgml
func CreateDocument() (*etree.Document, ElementInfoStruct) {

	DocXML := etree.NewDocument()
	DocXML.CreateProcInst("xml", `version="1.0" encoding="UTF-8" standalone="no"`)

	ElementGraphMl := DocXML.CreateElement("graphml")

	var ElementInfoGraphML ElementInfoStruct
	ElementInfoGraphML.Element = ElementGraphMl
	ElementInfoGraphML.Parent = nil

	ElementGraphMl.CreateAttr("xmlns", "http://graphml.graphdrawing.org/xmlns")
	ElementGraphMl.CreateAttr("xmlns:java", "http://www.yworks.com/xml/yfiles-common/1.0/java")
	ElementGraphMl.CreateAttr("xmlns:sys", "http://www.yworks.com/xml/yfiles-common/markup/primitives/2.0")
	ElementGraphMl.CreateAttr("xmlns:x", "http://www.yworks.com/xml/yfiles-common/markup/2.0")
	ElementGraphMl.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	ElementGraphMl.CreateAttr("xmlns:y", "http://www.yworks.com/xml/graphml")
	ElementGraphMl.CreateAttr("xmlns:y", "http://www.yworks.com/xml/graphml")
	ElementGraphMl.CreateAttr("xmlns:yed", "http://www.yworks.com/xml/yed/3")
	ElementGraphMl.CreateAttr("xsi:schemaLocation", "http://graphml.graphdrawing.org/xmlns http://www.yworks.com/xml/schema/graphml/1.1/ygraphml.xsd")

	ElementD0 := ElementGraphMl.CreateElement("key")
	ElementD0.CreateAttr("for", "port")
	ElementD0.CreateAttr("id", "d0")
	ElementD0.CreateAttr("yfiles.type", "portgraphics")

	ElementD1 := ElementGraphMl.CreateElement("key")
	ElementD1.CreateAttr("for", "port")
	ElementD1.CreateAttr("id", "d1")
	ElementD1.CreateAttr("yfiles.type", "portgeometry")

	ElementD2 := ElementGraphMl.CreateElement("key")
	ElementD2.CreateAttr("for", "port")
	ElementD2.CreateAttr("id", "d2")
	ElementD2.CreateAttr("yfiles.type", "portuserdata")

	ElementD3 := ElementGraphMl.CreateElement("key")
	ElementD3.CreateAttr("attr.name", "url")
	ElementD3.CreateAttr("attr.type", "string")
	ElementD3.CreateAttr("for", "node")
	ElementD3.CreateAttr("id", "d3")

	ElementD4 := ElementGraphMl.CreateElement("key")
	ElementD4.CreateAttr("attr.name", "description")
	ElementD4.CreateAttr("attr.type", "string")
	ElementD4.CreateAttr("for", "node")
	ElementD4.CreateAttr("id", "d4")

	ElementD5 := ElementGraphMl.CreateElement("key")
	ElementD5.CreateAttr("for", "node")
	ElementD5.CreateAttr("id", "d5")
	ElementD5.CreateAttr("yfiles.type", "nodegraphics")

	ElementD6 := ElementGraphMl.CreateElement("key")
	ElementD6.CreateAttr("for", "graphml")
	ElementD6.CreateAttr("id", "d6")
	ElementD6.CreateAttr("yfiles.type", "resources")

	ElementD7 := ElementGraphMl.CreateElement("key")
	ElementD7.CreateAttr("attr.name", "url")
	ElementD7.CreateAttr("attr.type", "string")
	ElementD7.CreateAttr("for", "edge")
	ElementD7.CreateAttr("id", "d7")

	ElementD8 := ElementGraphMl.CreateElement("key")
	ElementD8.CreateAttr("attr.name", "description")
	ElementD8.CreateAttr("attr.type", "string")
	ElementD8.CreateAttr("for", "edge")
	ElementD8.CreateAttr("id", "d8")

	ElementD9 := ElementGraphMl.CreateElement("key")
	ElementD9.CreateAttr("for", "edge")
	ElementD9.CreateAttr("id", "d9")
	ElementD9.CreateAttr("yfiles.type", "edgegraphics")

	ElementGraph := ElementGraphMl.CreateElement("graph")
	ElementGraph.CreateAttr("edgedefault", "directed")
	ElementGraph.CreateAttr("id", "G")

	return DocXML, ElementInfoGraphML
}

// FindId - находит ИД в формате "n1::n1::n1"
func FindId(ElementInfoMain, ElementInfo ElementInfoStruct) string {
	Otvet := ""
	//if Element == nil {
	//	return Otvet
	//}

	//if Element == ElementGraph0 {
	//	return Otvet
	//}

	if ElementInfo.Element.Tag == "node" {
		Otvet = "n" + strconv.Itoa(ElementInfo.Element.Index())
		//return Otvet
	}

	ParentSID := ""
	if ElementInfo.Parent != nil {
		ParentSID = FindId(ElementInfoMain, *ElementInfo.Parent)
	}
	if ParentSID != "" {
		if Otvet == "" {
			Otvet = ParentSID
		} else {
			Otvet = ParentSID + "::" + Otvet
		}
	}

	return Otvet
}
