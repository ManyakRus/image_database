package logic

import (
	"bytes"
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/packages_folder"
	"github.com/ManyakRus/image_database/internal/parse_go"
	"github.com/ManyakRus/image_database/pkg/graphml"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/beevik/etree"
	"golang.org/x/tools/go/packages"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

// MapNameURL - связь URL и имени внешнего сервиса
var MapServiceURL = make(map[string]string)

// MapPackagesElements - связь Пакета golang / Элемент файла .graphml
var MapPackagesElements = make(map[*packages.Package]*etree.Element, 0)

// MapServiceNameElements - связь ИД Пакета golang / Элемент файла .graphml
var MapServiceNameElements = make(map[string]*etree.Element, 0)

func StartFillAll(FileName string) bool {
	Otvet := false

	//RepositoryName := FindRepositoryName()
	RepositoryName := config.Settings.SERVICE_NAME

	DocXML, ElementGraph := graphml.CreateDocument()

	//рисуем элемент
	ElementShapeMain := graphml.CreateElement_Shape(ElementGraph, RepositoryName)

	//
	MassPackages := packages_folder.FindMassPackageFromFolder(config.Settings.DIRECTORY_SOURCE)
	FillPackage(MassPackages, ElementGraph)

	//заполним связи
	log.Info("Start fill links")
	FillLinks(ElementGraph, ElementShapeMain)

	if len(MapServiceNameElements) > 0 {
		Otvet = true
	}

	if Otvet == false {
		println("warning: Empty file not saved !")
		return Otvet
	}

	log.Info("Start save file")
	DocXML.Indent(2)
	err := DocXML.WriteToFile(FileName)
	if err != nil {
		log.Error("WriteToFile() FileName: ", FileName, " error: ", err)
	}

	return Otvet
}

// FillLinks - заполняет связи (стрелки) между пакетами
func FillLinks(ElementGraph, ElementShapeMain *etree.Element) {
	for ServiceName, ElementTo := range MapServiceNameElements {
		descr := config.Settings.SERVICE_NAME + " -> " + ServiceName
		graphml.CreateElement_Edge(ElementGraph, ElementShapeMain, ElementTo, "", descr)
	}
}

// FillLinks_goroutine - заполняет связи (стрелки) между пакетами для горутин go, синим цветом
func FillLinks_goroutine(ElementGraph *etree.Element) {
	for PackageFrom, ElementFrom := range MapPackagesElements {
		for _, Filename1 := range PackageFrom.GoFiles {

			AstFile, err := parse_go.ParseFile(Filename1)
			if err != nil {
				log.Warn("ParseFile() ", Filename1, " error: ", err)
				continue
			}

			MassGoImport := parse_go.FindGo(AstFile)
			for _, GoImport1 := range MassGoImport {
				//Go_package_name := GoImport1.Go_package_name
				Go_package_import := GoImport1.Go_package_import
				Go_func_name := GoImport1.Go_func_name

				ElementImport, ok := MapServiceNameElements[Go_package_import]
				if ok == false {
					//посторонние импорты
					//continue
					ElementImport = ElementFrom
				}
				label := Go_func_name
				descr := PackageFrom.Name + " -> " + GoImport1.Go_package_name
				graphml.CreateElement_Edge_blue(ElementGraph, ElementFrom, ElementImport, label, descr)
			}

			//ElementImport, ok := MapServiceNameElements[PackageImport.ID]
			//if ok == false {
			//	//посторонние импорты
			//	//log.Panic("MapPackagesElements[PackageImport] error: ok =false")
			//	continue
			//}
			//graphml.CreateElement_Edge(ElementGraph, ElementFrom.Index(), ElementImport.Index())
		}
	}
}

func FillPackage(MassPackages []*packages.Package, ElementGraph *etree.Element) {

	for _, v := range MassPackages {
		PackageName := v.PkgPath

		////test
		//pos1 := strings.Index(PackageName, "camunda")
		//if pos1 > 0 {
		//	print("test")
		//}

		ServiceName := FindNeedShow(PackageName)
		if ServiceName != "" {
			//проверка уже нарисован
			_, ok := MapServiceNameElements[ServiceName]
			if ok == true {
				//уже нарисован
				continue
			}

			//рисуем элемент
			ElementShape := graphml.CreateElement_Shape(ElementGraph, ServiceName)
			MapPackagesElements[v] = ElementShape
			MapServiceNameElements[ServiceName] = ElementShape

			//
			continue
		}

		//рекурсия всех импортов пакета
		if len(v.Imports) > 0 {
			// Convert map to slice of MassPackagesImports.
			MassPackagesImports := []*packages.Package{}
			for _, value := range v.Imports {
				pos1 := strings.Index(value.PkgPath, ".")
				//pos2 := strings.Index(value.PkgPath, "\\")
				if pos1 > 0 {
					MassPackagesImports = append(MassPackagesImports, value)
				}
			}

			FillPackage(MassPackagesImports, ElementGraph)
		}
	}

}

func FindNeedShow(PackageURL string) string {
	Otvet := ""

	lenPackage := len(PackageURL)

	for URL, ServiceName := range MapServiceURL {
		lenConn := len(URL)
		if lenConn > lenPackage {
			continue
		}
		if URL != PackageURL[:lenConn] {
			continue
		}
		Otvet = ServiceName
		break
	}

	return Otvet
}

//func FillFolder(ElementGraph, ElementGroup *etree.Element, Folder *folders.Folder) {
//
//	FolderName := Folder.Name
//
//	PackageFolder1 := packages_folder.FindPackageFromFolder(Folder)
//	PackageName := PackageFolder1.Name
//	if PackageName == "" && len(Folder.Folders) == 0 {
//		return
//	}
//
//	GroupName := FolderName
//	lines_count, func_count := FindLinesCount_package(PackageFolder1.Package)
//	if lines_count > 0 || func_count > 0 {
//		GroupName = GroupName + " ("
//		if func_count > 0 {
//			GroupName = GroupName + strconv.Itoa(func_count) + " func"
//			GroupName = GroupName + ", "
//		}
//		if lines_count > 0 {
//			GroupName = GroupName + strconv.Itoa(lines_count) + " lines"
//		}
//		GroupName = GroupName + ")"
//	}
//
//	//добавим группа (каталог)
//	var ElementGroup2 *etree.Element
//	if ElementGroup != nil {
//		ElementGroup2 = graphml.CreateElement_Group(ElementGroup, GroupName)
//	} else {
//		ElementGroup2 = graphml.CreateElement_Group(ElementGraph, GroupName)
//	}
//	if PackageName != "" {
//		//добавим пакет(package)
//		ElementShape := graphml.CreateElement_Shape(ElementGroup2, PackageName)
//		MapPackagesElements[PackageFolder1.Package] = ElementShape
//		MapServiceNameElements[PackageFolder1.Package.ID] = ElementShape
//		//MapPackagesElements[&PackageFolder1] = ElementShape
//	}
//
//	//сортировка
//	MassKeys := make([]string, 0, len(Folder.Folders))
//	for k := range Folder.Folders {
//		MassKeys = append(MassKeys, k)
//	}
//	sort.Strings(MassKeys)
//
//	//обход всех папок
//	for _, key1 := range MassKeys {
//		Folder1, ok := Folder.Folders[key1]
//		if ok == false {
//			log.Panic("Folder.Folders[key1] ok =false")
//		}
//		FillFolder(ElementGraph, ElementGroup2, Folder1)
//	}
//
//}

//// FindFileNameShort - возвращает имя файла(каталога) без пути
//func FindFileNameShort(path string) string {
//	Otvet := ""
//	if path == "" {
//		return Otvet
//	}
//	Otvet = filepath.Base(path)
//
//	return Otvet
//}

func FindLinesCount_package(Package1 *packages.Package) (int, int) {
	LinesCount := 0
	FuncCount := 0

	if Package1 == nil {
		return 0, 0
	}

	if Package1.GoFiles == nil {
		return 0, 0
	}

	for _, s := range Package1.GoFiles {
		count, func_count := FindLinesCount(s)
		LinesCount = LinesCount + count
		FuncCount = FuncCount + func_count
	}

	return LinesCount, FuncCount
}

func FindLinesCount(FileName string) (int, int) {
	LinesCount := 0
	FuncCount := 0

	bytes1, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(bytes1)
	LinesCount, err = LinesCount_reader(reader)
	if err != nil {
		log.Fatal(err)
	}

	FuncCount = FindFuncCount(&bytes1)

	return LinesCount, FuncCount
}

func LinesCount_reader(r io.Reader) (int, error) {
	Otvet := 0
	var err error

	buf := make([]byte, 8192)

	for {
		c, err := r.Read(buf)
		if err != nil {
			if err == io.EOF && c == 0 {
				break
			} else {
				return Otvet, err
			}
		}

		for _, b := range buf[:c] {
			if b == '\n' {
				Otvet++
			}
		}
	}

	if err == io.EOF {
		err = nil
	}

	return Otvet, err
}

// FindFuncCount - находит количество функций(func) в файле
func FindFuncCount(bytes *[]byte) int {
	Otvet := 0

	s := string(*bytes)
	sFind := "(\n|\t| )func( |\t)"

	Otvet = CountMatches(s, regexp.MustCompile(sFind))

	return Otvet
}

// CountMatches - находит количество совпадений в regexp
func CountMatches(s string, re *regexp.Regexp) int {
	total := 0
	for start := 0; start < len(s); {
		remaining := s[start:] // slicing the string is cheap
		loc := re.FindStringIndex(remaining)
		if loc == nil {
			break
		}
		// loc[0] is the start index of the match,
		// loc[1] is the end index (exclusive)
		start += loc[1]
		total++
	}
	return total
}

// FindRepositoryName - находит имя папки (репозитория)
func FindRepositoryName() string {
	Otvet := ""

	dir := config.Settings.DIRECTORY_SOURCE
	if dir == "" {
		return Otvet
	}
	if dir == "./" {
		dir = micro.ProgramDir()
	}
	dir = DeleteFileSeperator(dir)
	Otvet = path.Base(dir)

	return Otvet
}

// DeleteFileSeperator - убирает в конце / или \
func DeleteFileSeperator(dir string) string {
	Otvet := dir

	len1 := len(Otvet)
	if len1 == 0 {
		return Otvet
	}

	LastWord := Otvet[len1-1 : len1]
	if LastWord == micro.SeparatorFile() {
		Otvet = Otvet[0 : len1-1]
	}

	return Otvet
}
