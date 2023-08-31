package logic

import (
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/beevik/etree"
	"golang.org/x/tools/go/packages"
)

// MapNameURL - связь URL и имени внешнего сервиса
var MapServiceURL = make(map[string]string)

// MapPackagesElements - связь Пакета golang / Элемент файла .graphml
var MapPackagesElements = make(map[*packages.Package]*etree.Element, 0)

// MapServiceNameElements - связь ИД Пакета golang / Элемент файла .graphml
var MapServiceNameElements = make(map[string]*etree.Element, 0)

var MassTable []types.Table

func StartFillAll(FileName string) bool {
	Otvet := false

	return Otvet
}
