package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
	"strings"
	"time"
)

type TableColumn struct {
	TableName         string `json:"table_name"   gorm:"column:table_name;default:''"`
	ColumnName        string `json:"column_name"   gorm:"column:column_name;default:''"`
	ColumnType        string `json:"type_name"   gorm:"column:type_name;default:''"`
	ColumnIsIdentity  string `json:"is_identity"   gorm:"column:is_identity;default:''"`
	ColumnDescription string `json:"description"   gorm:"column:description;default:''"`
}

// FillMassTable - возвращает массив MassTable данными из БД
func FillMassTable() ([]types.Table, error) {
	var err error
	MassTable := make([]types.Table, 0)

	TextSQL := `
-- Все таблицы и колонки в схеме public

SELECT 
	c.table_name, 
	c.column_name,
	c.udt_name as type_name,
	c.is_identity as is_identity,
	pgd.description
	
	
FROM 
	pg_catalog.pg_statio_all_tables as st
	
inner join 
	pg_catalog.pg_description pgd 
on 
	pgd.objoid = st.relid

right join 
	information_schema.columns c 
on 
	pgd.objsubid   = c.ordinal_position
	and c.table_schema = st.schemaname
	and c.table_name   = st.relname

where 1=1
	and c.table_schema='public'

order by 
	table_name, 
	is_identity desc,
	column_name
`

	SCHEMA := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")
	if SCHEMA != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "public", SCHEMA)
	}

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//запрос
	tx := db.Raw(TextSQL)
	err = tx.Error
	if err != nil {
		sError := fmt.Sprint("db.Raw() error: ", err)
		log.Panicln(sError)
		return MassTable, err
	}

	//ответ в структуру
	MassTableColumn := make([]TableColumn, 0)
	tx = tx.Scan(&MassTableColumn)
	err = tx.Error
	if err != nil {
		sError := fmt.Sprint("Get_error()  error: ", err)
		log.Panicln(sError)
		return MassTable, err
	}

	//проверка 0 строк
	if tx.RowsAffected == 0 {
		sError := fmt.Sprint("db.Raw() RowsAffected =0 ")
		log.Warn(sError)
		err = errors.New(sError)
		log.Panicln(sError)
		return MassTable, err
	}

	//заполним MassTable
	TableName0 := ""
	Table1 := CreateTable()
	for _, v := range MassTableColumn {
		if v.TableName != TableName0 {
			if TableName0 != "" {
				MassTable = append(MassTable, Table1)
			}
			Table1 = CreateTable()
			Table1.Name = v.TableName
		}

		Column1 := types.Column{}
		Column1.Name = v.ColumnName
		Column1.Type = v.ColumnType
		if v.ColumnIsIdentity == "YES" {
			Column1.Is_identity = true
		}
		Column1.Description = v.ColumnDescription

		Table1.Columns = append(Table1.Columns, Column1)

		TableName0 = v.TableName
	}
	if Table1.Name != "" {
		MassTable = append(MassTable, Table1)
	}

	return MassTable, err
}

func CreateTable() types.Table {
	Otvet := types.Table{}
	Otvet.Columns = make([]types.Column, 0)

	return Otvet
}
